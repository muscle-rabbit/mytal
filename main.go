package main

import (
	"fmt"
	"log"
	"mytail/sub"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Name = "mytail"
	app.Usage = "This app display lines of text from the tail."
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "n",
			Value: "0",
			Usage: "Displays n lines of text from the tail. ",
		},
	}
	app.Version = "0.0.1"

	app.Action = func(context *cli.Context) error {
		lines := sub.FromFile(context.Args().Get(0))
		paragraphBegin := 0
		numberOfLines := len(lines)
		numberOfLinesFromTail := context.Int("n")
		defaultLines := 10
		if numberOfLines > defaultLines {
			paragraphBegin = numberOfLines - defaultLines
		}
		if numberOfLinesFromTail < numberOfLines && numberOfLinesFromTail > 0 {
			paragraphBegin = len(lines) - context.Int("n")
		}

		for i := paragraphBegin; i < len(lines); i++ {
			fmt.Println(lines[i])
		}
		return nil
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
