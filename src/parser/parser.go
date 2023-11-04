package parser

import "github.com/captainmango/go-cron-parser/src/shared"

type CronParserInterface interface {
	Parse(cron shared.Cron) shared.ParsedCron
}

type CronParserFunc func(cron shared.Cron) shared.ParsedCron

type RuleFunc func(input []int, interval shared.CronInterval) []int

type CronParser struct {}

func (c CronParser) Parse(cron shared.Cron) shared.ParsedCron {
	output := shared.ParsedCron{}

	for k, v := range cron {
		values := process(v, k)
		output[k] = values
	}

	return output
}