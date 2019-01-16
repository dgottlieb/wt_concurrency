package wt_concurrency

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
)

//  	std::cout << "Line: Timestamp :read 35" << std::endl;
//  	std::cout << "Idx: 1" << std::endl;
//  	std::cout << "HasOutput: false" << std::endl;
//  	// canError
//  	{
//  		int errorCode =
//  			session_1.readAtTimestamp(35);
//  		if (errorCode == 0) { } else { std::cout << "Error: " << errorCode << " Str: " << wiredtiger_strerror(errorCode) << std::endl; }
//  	}

var lineRe = regexp.MustCompile("Line: (.+)")
var actorIdxRe = regexp.MustCompile("Idx: (\\d+)")
var hasOutputRe = regexp.MustCompile("HasOutput: (.+)")
var valRe = regexp.MustCompile("Val: (-?\\d+)")
var errorRe = regexp.MustCompile("Error: (\\d+) Str: (.+)")
var numActorsRe = regexp.MustCompile("Actors: (\\d+)")
var wtErrStrRe = regexp.MustCompile("\\[1\\d{9}:\\d{6}.*?(WT_.*)")

type Table struct {
	Actors     []string
	Recombined []*Recombined
	Footnotes  []string
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (table *Table) Output() {
	maxColSize := make([]int, len(table.Actors))
	for idx, actorName := range table.Actors {
		maxColSize[idx] = max(maxColSize[idx], len(actorName))
	}

	for _, recombined := range table.Recombined {
		var act string
		switch {
		// Basic, no error, no output
		case recombined.errCode == 0 && recombined.val == -2:
			act = recombined.line
		case recombined.errCode == 0 && recombined.val == -1:
			act = fmt.Sprintf("%v (WT_NOTFOUND)", recombined.line)
		case recombined.errCode == 0 && recombined.val >= 0:
			act = fmt.Sprintf("%v (%v)", recombined.line, recombined.val)
		case recombined.errCode != 0 && recombined.wtErrStr == "":
			act = fmt.Sprintf("%v (%v %v)", recombined.line, recombined.errCode, recombined.errStr)
		case recombined.errCode != 0 && recombined.wtErrStr != "":
			act = fmt.Sprintf("%v (%v %v) [%d]", recombined.line, recombined.errCode, recombined.errStr, len(table.Footnotes)+1)
			table.Footnotes = append(table.Footnotes, recombined.wtErrStr)
		default:
			panic(fmt.Sprintf("Unknown case. Rec: %+v", *recombined))
		}

		if recombined.actorIdx > -1 {
			maxColSize[recombined.actorIdx] = max(maxColSize[recombined.actorIdx], len(act))
		} else {
			act = ""
		}
		recombined.unpaddedAct = act
	}

	padArray(table.Actors, maxColSize)
	fmt.Printf("|%s|\n", strings.Join(table.Actors, "|"))
	underRow := make([]string, len(table.Actors))
	for idx, _ := range table.Actors {
		// All columns are padding on the left and right by one space.
		const padding = 2
		underRow[idx] = strings.Repeat("-", maxColSize[idx]+2)
	}
	fmt.Printf("|%s|\n", strings.Join(underRow, "+"))

	for _, recombined := range table.Recombined {
		acts := make([]string, len(table.Actors))
		if recombined.actorIdx == -1 {
			padArray(acts, maxColSize)
			fmt.Printf("|%s|\n", strings.Join(acts, "|"))
			continue
		}

		acts[recombined.actorIdx] = recombined.unpaddedAct
		padArray(acts, maxColSize)
		fmt.Printf("|%s|\n", strings.Join(acts, "|"))
	}

	if len(table.Footnotes) == 0 {
		return
	}

	fmt.Println("\nFootnotes:")
	for idx, footnote := range table.Footnotes {
		fmt.Printf("  [%d] %v\n", idx+1, footnote)
	}
}

func pad(act string, colSize int) string {
	spacesToAdd := colSize - len(act)
	return fmt.Sprintf(" %v%v", act, strings.Repeat(" ", spacesToAdd))
}

func padArray(acts []string, maxColSize []int) {
	for idx, act := range acts {
		acts[idx] = pad(act, maxColSize[idx]+1)
	}
}

type Recombined struct {
	line     string
	actorIdx int
	val      int
	errCode  int
	errStr   string
	wtErrStr string

	unpaddedAct string
}

func NewRecombined(line string) *Recombined {
	return &Recombined{line: line, actorIdx: -1, val: -2}
}

func ParseOutput(reader io.Reader) error {
	scanner := bufio.NewScanner(reader)
	var table Table
	var recombined *Recombined
	var err error
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if matches := lineRe.FindStringSubmatch(line); len(matches) > 0 {
			if recombined != nil {
				table.Recombined = append(table.Recombined, recombined)
			}
			recombined = NewRecombined(matches[1])
			continue

		} else if matches := actorIdxRe.FindStringSubmatch(line); len(matches) > 0 {
			recombined.actorIdx, err = strconv.Atoi(matches[1])
			if err != nil {
				panic(err)
			}
			continue
		} else if matches := hasOutputRe.FindStringSubmatch(line); len(matches) > 0 {
			continue
		} else if matches := valRe.FindStringSubmatch(line); len(matches) > 0 {
			recombined.val, err = strconv.Atoi(matches[1])
			if err != nil {
				panic(err)
			}
			continue
		} else if matches := errorRe.FindStringSubmatch(line); len(matches) > 0 {
			recombined.errCode, err = strconv.Atoi(matches[1])
			if err != nil {
				panic(err)
			}
			recombined.errStr = matches[2]
		} else if matches := numActorsRe.FindStringSubmatch(line); len(matches) > 0 {
			var numActors int
			numActors, err = strconv.Atoi(matches[1])
			if err != nil {
				panic(err)
			}

			table.Actors = make([]string, numActors)
			for idx := 0; idx < numActors; idx++ {
				if scanner.Scan() == false {
					panic("cannot parse actors")
				}
				table.Actors[idx] = strings.TrimSpace(scanner.Text())
			}
		} else if matches := wtErrStrRe.FindStringSubmatch(line); len(matches) > 0 {
			recombined.wtErrStr = matches[1]
		}
	}

	table.Recombined = append(table.Recombined, recombined)
	table.Output()
	return nil
}
