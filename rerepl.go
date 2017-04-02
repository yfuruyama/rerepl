package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/peterh/liner"
)

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

		patternAndTarget := strings.SplitN(line, " ", 2)
		if len(patternAndTarget) != 2 {
			fmt.Printf("invalid input: %s\n", line)
			continue
		}

		pattern := patternAndTarget[0]
		target := patternAndTarget[1]

		re, err := regexp.Compile(pattern)
		matched := re.MatchString(target)
		groups := re.FindStringSubmatch(target)

		fmt.Printf("matched: %t\n", matched)
		fmt.Printf("captures: \n")
		if groups != nil {
			for idx, value := range groups[1:] {
				fmt.Printf("  %d: %s\n", idx+1, value)
			}
		}
	}

	// write history
	if f, err := os.Create(historyPath); err == nil {
		liner.WriteHistory(f)
		f.Close()
	}
}
