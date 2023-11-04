package main

import (
	"fmt"
	"os"

	"github.com/captainmango/go-cron-parser/src/parser"
	"github.com/captainmango/go-cron-parser/src/printer"
	"github.com/captainmango/go-cron-parser/src/shared"
)

func main() {
	parser := parser.NewParser()
	printer := printer.NewPrinter()

	input := os.Args
	cron, err := shared.NewCronFromArgs(input)

	if err != nil {
		fmt.Printf("Encountered error creating cron: %s", err)
		return
	}

	parsedCron, err := parser.Parse(cron)

	if err != nil {
		fmt.Printf("Encountered error parsing cron: %s", err)
		return
	}

	printer.Print(parsedCron, os.Stdout)
}