package wt_concurrency

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Program []Instance

type Instance struct {
	TableName string
	Actors    []Actor
	Sequence  []*WrappedOp
}

func (instance *Instance) NextOp(op *WrappedOp) {
	instance.Sequence = append(instance.Sequence, op)
}

func contains(canError []int, stmtIdx int) bool {
	for _, idx := range canError {
		if idx == stmtIdx {
			return true
		}
	}

	return false
}

func (instance *Instance) Compile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fmt.Fprintln(file, string(MustAsset("include/wt_raii.h")))
	fmt.Fprintln(file)

	fmt.Fprintln(file, "int main() {")
	fmt.Fprintln(file, "\tsystem(\"rm ./WT_HOME/*\");")
	fmt.Fprintln(file, "\tsystem(\"mkdir -p ./WT_HOME/journal/\");")
	fmt.Fprintln(file)

	fmt.Fprintf(file, "\tconst std::string tableUri = \"table:%s\";\n", instance.TableName)
	fmt.Fprintln(file)

	fmt.Fprintln(file, "\tWtConn conn(\"./WT_HOME\");")
	fmt.Fprintln(file, "\tWtSession admin = conn.getSession();")
	fmt.Fprintln(file, "\tadmin.createTable(tableUri);")
	fmt.Fprintln(file, "\tadmin.alterTableLogging(tableUri, false);")
	fmt.Fprintln(file)

	fmt.Fprintf(file, "\tstd::cout << \"Actors: %v\" << std::endl;\n", len(instance.Actors))
	for _, actor := range instance.Actors {
		fmt.Fprintf(file, "\tWtSession %s = conn.getSession();\n", actor.SessionName())
		fmt.Fprintf(file, "\tstd::cout << \"%v\" << std::endl;\n", actor.Name)
		fmt.Fprintf(file, "\tWtCursor %s = %s.openCursor(tableUri);", actor.CursorName(), actor.SessionName())
	}
	fmt.Fprintln(file)

	for idx, _ := range instance.Sequence {
		row := instance.Sequence[idx]
		fmt.Fprintf(file, "\tstd::cout << \"Line: %v\" << std::endl;\n", row.Raw)
		fmt.Fprintf(file, "\tstd::cout << \"Idx: %v\" << std::endl;\n", row.ActorIdx)
		fmt.Fprintf(file, "\tstd::cout << \"HasOutput: %v\" << std::endl;\n", row.HasOutput)
		canError := row.CanError()
		for stmtIdx, line := range row.Do() {
			switch {
			case contains(canError, stmtIdx) && row.HasOutput:
				fmt.Fprintf(file, "\t// canError, hasOutput\n")
				fmt.Fprintf(file, "\t{\n\t\tint ret =\n\t\t")
				fmt.Fprintf(file, "\t%s\n", line)
				fmt.Fprintf(file, "\t\tif (ret >= -1) { std::cout << \"Val: \" << ret << std::endl; } else { std::cout << \"Error: \" << ret << \" Str: \" << wiredtiger_strerror(ret) << std::endl; }\n")
				fmt.Fprintf(file, "\t}\n")
			case contains(canError, stmtIdx):
				fmt.Fprintf(file, "\t// canError\n")
				fmt.Fprintf(file, "\t{\n\t\tint errorCode =\n\t\t")
				fmt.Fprintf(file, "\t%s\n", line)
				fmt.Fprintf(file, "\t\tif (errorCode == 0) { } else { std::cout << \"Error: \" << errorCode << \" Str: \" << wiredtiger_strerror(errorCode) << std::endl; }\n")
				fmt.Fprintf(file, "\t}\n")
			default:
				fmt.Fprintf(file, "\t%s\n", line)
			}
		}
	}
	fmt.Fprintln(file, "}")
}

type Actor struct {
	ColumnId int
	Name     string
}

func (actor Actor) SessionName() string {
	return fmt.Sprintf("session_%d", actor.ColumnId)
}

func (actor Actor) CursorName() string {
	return fmt.Sprintf("session_%d_cursor", actor.ColumnId)
}

func ActorsFromLine(line string) []Actor {
	ret := make([]Actor, 0)
	actorId := 0
	for _, item := range strings.Split(line, "|") {
		trimmed := strings.TrimSpace(item)
		if trimmed == "" {
			continue
		}

		ret = append(ret, Actor{actorId, trimmed})
		actorId++
	}

	return ret
}

type Operation interface {
	Do() []string
	CanError() []int
}

type WrappedOp struct {
	Operation
	Raw       string
	ActorIdx  int
	HasOutput bool
}

var kvRe = regexp.MustCompile(":(\\w+) (\\w+)")

func KeyValues(line string) map[string]string {
	ret := make(map[string]string)
	found := kvRe.FindAllStringSubmatch(line, -1)
	for _, item := range found {
		ret[item[1]] = item[2]
	}

	return ret
}

type BeginTxn struct {
	Actor
	ReadTimestamp uint64
	Isolation     string
	IgnorePrepare bool
}

func ParseBeginTxn(actor *Actor, item string) BeginTxn {
	options := KeyValues(item)
	ret := BeginTxn{Actor: *actor}
	if value, exists := options["isolation"]; exists {
		ret.Isolation = value
	}

	var err error
	if value, exists := options["readAt"]; exists {
		ret.ReadTimestamp, err = strconv.ParseUint(value, 10, 64)
		if err != nil {
			panic(err)
		}
	}

	if value, exists := options["ignorePrepare"]; exists {
		ret.IgnorePrepare, err = strconv.ParseBool(value)
		if err != nil {
			panic(err)
		}
	}

	return ret
}

func (beginTxn BeginTxn) Do() []string {
	if beginTxn.ReadTimestamp > 0 && beginTxn.IgnorePrepare {
		panic("Does not support beginning with a read timestamp and ignore prepared because I can't be bothered to implement it atm.")
	}

	if beginTxn.ReadTimestamp > 0 {
		return []string{
			fmt.Sprintf(beginTxn.SessionName()+".beginAtTimestamp(%d);", beginTxn.ReadTimestamp),
		}
	}

	if beginTxn.IgnorePrepare {
		return []string{fmt.Sprintf(beginTxn.SessionName()+".begin(%v);", beginTxn.IgnorePrepare)}
	}

	return []string{beginTxn.SessionName() + ".begin();"}
}

func (beginTxn BeginTxn) CanError() []int {
	return []int{0}
}

type CommitTxn struct {
	Actor
	CommitTimestamp uint64
}

func ParseCommitTxn(actor *Actor, item string) CommitTxn {
	options := KeyValues(item)
	ret := CommitTxn{Actor: *actor}
	if value, exists := options["commit"]; exists {
		var err error
		ret.CommitTimestamp, err = strconv.ParseUint(value, 10, 64)
		if err != nil {
			panic(err)
		}
	}

	return ret
}

func (commitTxn CommitTxn) Do() []string {
	if commitTxn.CommitTimestamp > 0 {
		return []string{
			fmt.Sprintf("%s.commit(%d);", commitTxn.SessionName(), commitTxn.CommitTimestamp),
		}
	}

	return []string{commitTxn.SessionName() + ".commit();"}
}

func (commitTxn CommitTxn) CanError() []int {
	return []int{0}
}

type RollbackTxn struct {
	Actor
}

func (rollback RollbackTxn) Do() []string {
	return []string{fmt.Sprintf("%s.rollback();", rollback.SessionName())}
}

func (rollback RollbackTxn) CanError() []int {
	return []int{0}
}

type Write struct {
	Actor
	Key   int
	Value int
}

var writeRe = regexp.MustCompile("Write (\\w) (\\d+)")

func ParseWrite(actor *Actor, item string) Write {
	// options := KeyValues(item)
	ret := Write{Actor: *actor}
	matches := writeRe.FindStringSubmatch(item)
	ret.Key = int(matches[1][0] - byte('A'))
	var err error
	ret.Value, err = strconv.Atoi(matches[2])
	if err != nil {
		panic(err)
	}

	return ret
}

func (write Write) Do() []string {
	return []string{
		fmt.Sprintf("\t%s.insert(%d, %d);", write.CursorName(), write.Key, write.Value),
	}
}

func (write Write) CanError() []int {
	return []int{0}
}

type Read struct {
	Actor
	Key int
}

var readRe = regexp.MustCompile("Read (\\w)")

func ParseRead(actor *Actor, item string) Read {
	// options := KeyValues(item)
	ret := Read{Actor: *actor}
	keyChr := readRe.FindStringSubmatch(item)[1][0]
	ret.Key = int(keyChr - byte('A'))
	return ret
}

func (read Read) CursorName() string {
	return fmt.Sprintf("%s_cursor", read.SessionName())
}

func (read Read) Do() []string {
	return []string{
		fmt.Sprintf("\t%s.searchExact(%d);", read.CursorName(), read.Key),
	}
}

func (read Read) CanError() []int {
	return []int{0}
}

type TimestampTxn struct {
	Actor
	Read   int
	Commit int
}

func ParseTimestamp(actor *Actor, item string) TimestampTxn {
	options := KeyValues(item)
	ret := TimestampTxn{Actor: *actor}
	var err error
	if val, exists := options["commit"]; exists {
		ret.Commit, err = strconv.Atoi(val)
		if err != nil {
			panic(err)
		}
	}
	if val, exists := options["read"]; exists {
		ret.Read, err = strconv.Atoi(val)
		if err != nil {
			panic(err)
		}
	}

	return ret
}

func (timestamp TimestampTxn) Do() []string {
	ret := make([]string, 0)
	if timestamp.Read > 0 {
		ret = append(ret,
			fmt.Sprintf("%s.readAtTimestamp(%d);", timestamp.SessionName(), timestamp.Read))
	}
	if timestamp.Commit > 0 {
		ret = append(ret, fmt.Sprintf("%s.setTimestamp(%d);", timestamp.SessionName(), timestamp.Commit))
	}
	return ret
}

func (timestamp TimestampTxn) CanError() []int {
	ret := make([]int, 0)
	for idx, _ := range timestamp.Do() {
		ret = append(ret, idx)
	}
	return ret
}

type NoOp struct{}

func (noop NoOp) Do() []string {
	return []string{}
}

func (noop NoOp) CanError() []int {
	return []int{}
}

var TRUE bool = true
var FALSE bool = false

type Alter struct {
	Actor
	TableName string
	Logging   *bool
}

func ParseAlter(instance *Instance, actor *Actor, item string) Alter {
	options := KeyValues(item)
	ret := Alter{Actor: *actor}
	if val, exists := options["logging"]; exists {
		switch val {
		case "on":
			ret.Logging = &TRUE
		case "off":
			ret.Logging = &FALSE
		default:
			panic(fmt.Sprintf("Unknown logging setting. Must be <on|off>. Val: %v", val))
		}
	}
	if ret.TableName == "" {
		ret.TableName = instance.TableName
	}

	return ret
}

func (alter Alter) Do() []string {
	ret := make([]string, 0)
	if alter.Logging != nil {
		ret = append(ret, fmt.Sprintf("%v.alterTableLogging(%v, %v);", alter.SessionName(), alter.TableName, *alter.Logging))
	}

	return ret
}

func (alter Alter) CanError() []int {
	return []int{0}
}

type GlobalTimestamp struct {
	Stable int
	Oldest int
	Commit int
	Force  bool
}

func ParseGlobalTimestamp(instance *Instance, item string) GlobalTimestamp {
	options := KeyValues(item)
	ret := GlobalTimestamp{}

	var err error
	if val, exists := options["stable"]; exists {
		ret.Stable, err = strconv.Atoi(val)
		if err != nil {
			panic(err)
		}
	}
	if val, exists := options["oldest"]; exists {
		ret.Oldest, err = strconv.Atoi(val)
		if err != nil {
			panic(err)
		}
	}
	if val, exists := options["commit"]; exists {
		ret.Commit, err = strconv.Atoi(val)
		if err != nil {
			panic(err)
		}
	}
	if val, exists := options["force"]; exists {
		ret.Force, err = strconv.ParseBool(val)
		if err != nil {
			panic(err)
		}
	}

	return ret
}

func (global GlobalTimestamp) Do() []string {
	ret := make([]string, 0)
	if global.Stable > 0 {
		ret = append(ret, fmt.Sprintf("conn.setStableTimestamp(%d);", global.Stable))
	}
	if global.Oldest > 0 {
		ret = append(ret, fmt.Sprintf("conn.setOldestTimestamp(%d);", global.Oldest))
	}
	if global.Commit > 0 {
		panic("Resetting the all_committed time is not yet supported.")
	}
	if global.Force {
		panic("Working with force=true is not yet supported.")
	}

	return ret
}

func (global GlobalTimestamp) CanError() []int {
	ret := make([]int, 0)
	for idx, _ := range global.Do() {
		ret = append(ret, idx)
	}

	return ret
}

type Checkpoint struct {
	Actor
	Stable bool
}

func ParseCheckpoint(actor *Actor, item string) Checkpoint {
	options := KeyValues(item)
	ret := Checkpoint{*actor, true} // Default stable to true.
	var err error
	if val, exists := options["stable"]; exists {
		ret.Stable, err = strconv.ParseBool(val)
		if err != nil {
			panic(err)
		}
	}

	return ret
}

func (checkpoint Checkpoint) Do() []string {
	if checkpoint.Stable {
		return []string{"conn.stableCheckpoint();"}
	} else {
		return []string{"conn.unstableCheckpoint();"}
	}
}

func (checkpoint Checkpoint) CanError() []int {
	return []int{0}
}

type PrepareTxn struct {
	Actor
	PrepareTimestamp int
}

var prepareRe = regexp.MustCompile("Prepare (\\d+)")

func ParsePrepare(actor *Actor, item string) PrepareTxn {
	ret := PrepareTxn{Actor: *actor}
	matches := prepareRe.FindStringSubmatch(item)

	var err error
	ret.PrepareTimestamp, err = strconv.Atoi(matches[1])
	if err != nil {
		panic(err)
	}
	return ret
}

func (prepare PrepareTxn) Do() []string {
	ret := make([]string, 0)
	ret = append(ret, fmt.Sprintf("%s.prepare(%d);", prepare.SessionName(), prepare.PrepareTimestamp))
	return ret
}

func (prepare PrepareTxn) CanError() []int {
	return []int{0}
}

func ParseAndNormalize(line string) []string {
	ret := make([]string, 0)
	items := strings.Split(line, "|")
	for _, item := range items {
		if len(item) == 0 {
			continue
		}

		ret = append(ret, strings.TrimSpace(item))
	}

	return ret
}

func ParseOp(instance *Instance, actors []Actor, line string) *WrappedOp {
	items := ParseAndNormalize(line)

	realOps := 0
	var op Operation
	var actorIdx int
	var hasOutput bool = false
	var raw string
	for idx, item := range items {
		item = strings.TrimSpace(item)
		if len(item) == 0 {
			continue
		}

		actorIdx = idx
		raw = item
		realOps++

		switch {
		case strings.HasPrefix(item, "Begin"):
			op = ParseBeginTxn(&actors[idx], item)
		case strings.HasPrefix(item, "Commit"):
			op = ParseCommitTxn(&actors[idx], item)
		case strings.HasPrefix(item, "Write"):
			op = ParseWrite(&actors[idx], item)
		case strings.HasPrefix(item, "Read"):
			op = ParseRead(&actors[idx], item)
			hasOutput = true
		case strings.HasPrefix(item, "Timestamp"):
			op = ParseTimestamp(&actors[idx], item)
		case item == "Rollback":
			op = RollbackTxn{actors[idx]}
		case strings.HasPrefix(item, "Alter"):
			op = ParseAlter(instance, &actors[idx], item)
		case strings.HasPrefix(item, "GlobalTimestamp"):
			op = ParseGlobalTimestamp(instance, item)
		case strings.HasPrefix(item, "Checkpoint"):
			op = ParseCheckpoint(&actors[idx], item)
		case strings.HasPrefix(item, "Prepare"):
			op = ParsePrepare(&actors[idx], item)
		default:
			panic(fmt.Sprintf("Unknown command. Line: %s", item))
		}
	}

	if realOps == 0 {
		return &WrappedOp{NoOp{}, line, -1, false}
	}

	if realOps != 1 {
		panic(fmt.Sprintf("Expected exactly one op. Num: %d Line: %s", realOps, line))
	}

	return &WrappedOp{op, raw, actorIdx, hasOutput}
}

func ParseProgram(reader io.Reader) (*Instance, error) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line)[0] != '#' {
			break
		}
	}

	firstLine := scanner.Text()
	if scanner.Scan() == false {
		return nil, errors.New("No instructions in file.")
	}

	secondLine := scanner.Text()
	for _, chr := range secondLine {
		if chr != '|' && chr != '-' && chr != '+' {
			return nil, fmt.Errorf("Unexpected second line. Line: %v", secondLine)
		}
	}

	instance := Instance{
		TableName: "tableUri",
		Actors:    ActorsFromLine(firstLine),
	}

	for scanner.Scan() {
		instance.NextOp(ParseOp(&instance, instance.Actors, scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return &instance, nil
}
