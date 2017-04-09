package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/peterh/liner"
)

type Result struct {
	matched  bool
	captures []string
}

func main() {
	liner := liner.NewLiner()
	defer liner.Close()

	liner.SetCtrlCAborts(true)

	// read history
	historyPath := filepath.Join(os.TempDir(), ".rerepl_history")
	if f, err := os.Open(historyPath); err == nil {
		liner.ReadHistory(f)
		f.Close()
	}

	for {
		line, err := liner.Prompt("> ")
		if err != nil {
			break
		}
		if line == "" {
			continue
		}
		liner.AppendHistory(line)

		result, err := EvalLine(line)
		if err != nil {
			fmt.Printf(err.Error())
			continue
		}

		fmt.Printf("matched: %t\n", result.matched)
		fmt.Printf("captures: \n")
		for idx, value := range result.captures {
			fmt.Printf("  %d: %s\n", idx+1, value)
		}
	}

	// write history
	if f, err := os.Create(historyPath); err == nil {
		liner.WriteHistory(f)
		f.Close()
	}
}

func EvalLine(line string) (*Result, error) {
	patternAndTarget := strings.SplitN(line, " ", 2)
	if len(patternAndTarget) != 2 {
		return nil, errors.New(fmt.Sprintf("invalid input: %s\n", line))
	}

	pattern := patternAndTarget[0]
	target := patternAndTarget[1]

	re, err := regexp.Compile(pattern)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("invalid regexp: %s\n", pattern))
	}

	matched := re.MatchString(target)
	groups := re.FindStringSubmatch(target)

	captures := make([]string, 0)

	if groups != nil {
		for _, value := range groups[1:] {
			captures = append(captures, value)
		}
	}

	return &Result{
		matched:  matched,
		captures: captures,
	}, nil
}
