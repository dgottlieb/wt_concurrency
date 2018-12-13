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
	Sequence  []Operation
}

func (instance *Instance) NextOp(op Operation) {
	instance.Sequence = append(instance.Sequence, op)
}

func (instance *Instance) Compile(filename string) {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fmt.Fprintln(file, "#include \"include/wt_raii.h\"")
	fmt.Fprintln(file)

	fmt.Fprintln(file, "int main() {")
	fmt.Fprintln(file, "\tsystem(\"rm -rf ./WT_HOME/journal/*\");")
	fmt.Fprintln(file, "\tsystem(\"rm ./WT_HOME/*\");")
	fmt.Fprintln(file)

	fmt.Fprintf(file, "\tconst std::string tableUri = \"table:%s\";\n", instance.TableName)
	fmt.Fprintln(file)

	fmt.Fprintln(file, "\tWtConn conn(\"./WT_HOME\");")
	fmt.Fprintln(file, "\tWtSession admin = conn.getSession();")
	fmt.Fprintln(file, "\tadmin.createTable(tableUri);")
	fmt.Fprintln(file, "\tadmin.alterTableLogging(tableUri, false);")
	fmt.Fprintln(file)

	for _, actor := range instance.Actors {
		fmt.Fprintf(file, "\tWtSession %s = conn.getSession();\n", actor.SessionName())
	}
	fmt.Fprintln(file)

	for idx, _ := range instance.Sequence {
		for _, line := range instance.Sequence[idx].Do() {
			fmt.Fprintf(file, "\t%s\n", line)
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

func ActorsFromLine(line string) []Actor {
	ret := make([]Actor, 0)
	actorId := 0
	for _, item := range strings.Split(line, "|") {
		if strings.TrimSpace(item) == "" {
			continue
		}

		ret = append(ret, Actor{actorId, item})
		actorId++
	}

	return ret
}

type Operation interface {
	Do() []string
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
}

func ParseBeginTxn(actor *Actor, item string) BeginTxn {
	options := KeyValues(item)
	ret := BeginTxn{Actor: *actor}
	if value, exists := options["isolation"]; exists {
		ret.Isolation = value
	}

	if value, exists := options["readAt"]; exists {
		var err error
		ret.ReadTimestamp, err = strconv.ParseUint(value, 10, 64)
		if err != nil {
			panic(err)
		}
	}

	return ret
}

type CommitTxn struct {
	Actor
	CommitTimestamp uint64
}

func ParseCommitTxn(actor *Actor, item string) CommitTxn {
	options := KeyValues(item)
	ret := CommitTxn{Actor: *actor}
	if value, exists := options["commitAt"]; exists {
		var err error
		ret.CommitTimestamp, err = strconv.ParseUint(value, 10, 64)
		if err != nil {
			panic(err)
		}
	}

	return ret
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

func (beginTxn BeginTxn) Do() []string {
	if beginTxn.ReadTimestamp > 0 {
		return []string{
			fmt.Sprintf(beginTxn.SessionName()+".beginAtTimestamp(%d);", beginTxn.ReadTimestamp),
		}
	}

	return []string{beginTxn.SessionName() + ".begin();"}
}

func (commitTxn CommitTxn) Do() []string {
	if commitTxn.CommitTimestamp > 0 {
		return []string{
			fmt.Sprintf(commitTxn.SessionName()+".commit(%d);", commitTxn.CommitTimestamp),
		}
	}

	return []string{commitTxn.SessionName() + ".commit();"}
}

func (read Read) CursorName() string {
	return fmt.Sprintf("%s_cursor", read.SessionName())
}

func (read Read) Do() []string {
	return []string{
		"{",
		fmt.Sprintf("\tWtCursor %s = %s.openCursor(tableUri);", read.CursorName(), read.SessionName()),
		fmt.Sprintf("\tstd::cout << \"Val: \" << %s.searchExact(%d) << std::endl;", read.CursorName(), read.Key),
		"}",
	}

}

func (write Write) CursorName() string {
	return fmt.Sprintf("%s_cursor", write.SessionName())
}

func (write Write) Do() []string {
	return []string{
		"{",
		fmt.Sprintf("\tWtCursor %s = %s.openCursor(tableUri);", write.CursorName(), write.SessionName()),
		fmt.Sprintf("\t%s.insert(%d, %d);", write.CursorName(), write.Key, write.Value),
		"}",
	}
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

func ParseOp(actors []Actor, line string) Operation {
	items := ParseAndNormalize(line)

	realOps := 0
	var op Operation
	for idx, item := range items {
		item = strings.TrimSpace(item)
		if len(item) == 0 {
			continue
		}

		fmt.Printf("%d -> `%s`\n", idx, item)

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
		default:
			panic(fmt.Sprintf("Unknown command. Line: %s", item))
		}
	}

	if realOps != 1 {
		panic(fmt.Sprintf("Expected exactly one op. Num: %d Line: %s", realOps, line))
	}

	return op
}

func ParseProgram(reader io.Reader) (*Instance, error) {
	scanner := bufio.NewScanner(reader)
	if scanner.Scan() == false {
		return nil, errors.New("No text in file.")
	}

	firstLine := scanner.Text()
	if scanner.Scan() == false {
		return nil, errors.New("No instructions in file.")
	}

	secondLine := scanner.Text()
	for _, chr := range secondLine {
		if chr != '|' && chr != '-' && chr != '+' {
			return nil, fmt.Errorf("Unexpected second line. Line:\n\t", secondLine)
		}
	}

	instance := Instance{
		TableName: "tableUri",
		Actors:    ActorsFromLine(firstLine),
	}

	for scanner.Scan() {
		instance.NextOp(ParseOp(instance.Actors, scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return &instance, nil
}
