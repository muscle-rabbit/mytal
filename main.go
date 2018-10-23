package main

import (
	"flag"
	"fmt"
	"mytail/sub"
)

func main() {
	var ip *int = flag.Int("n", 0, "help message for flagname")
	flag.Parse()
	var lines = sub.FromFile(flag.Arg(0))
	var paragraphBegin = 0
	var numberOfLines = len(lines)
	var numberOfLinesFromTail = *ip
	var defaultLines = 10
	if numberOfLines > defaultLines {
		paragraphBegin = numberOfLines - defaultLines
	}
	if numberOfLinesFromTail < numberOfLines && numberOfLinesFromTail > 0 {
		paragraphBegin = numberOfLines - numberOfLinesFromTail
	}

	for i := paragraphBegin; i < len(lines); i++ {
		fmt.Println(lines[i])
	}
}
