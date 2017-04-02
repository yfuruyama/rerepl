package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/peterh/liner"
)

func main() {
	liner := liner.NewLiner()
	defer liner.Close()

	liner.SetCtrlCAborts(true)
	for {
		line, err := liner.Prompt("> ")
		if err != nil {
			break
		}
		patternAndTarget := strings.Split(line, " ")
		pattern := patternAndTarget[0]
		target := patternAndTarget[1]

		re, err := regexp.Compile(pattern)
		matched := re.MatchString(target)
		groups := re.FindStringSubmatch(target)

		fmt.Printf("matched: %t\n", matched)
		fmt.Printf("captures: \n")
		if groups != nil {
			for idx, value := range groups[1:] {
				fmt.Printf("  %d: %s\n", idx, value)
			}
		}
	}
}
