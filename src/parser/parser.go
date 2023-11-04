package parser

import (
	"github.com/captainmango/go-cron-parser/src/shared"
)

type CronParserInterface interface {
	Parse(cron shared.Cron) shared.ParsedCron
}

type CronParserFunc func(cron shared.Cron) (shared.ParsedCron, error)

func (cf CronParserFunc) Parse(cron shared.Cron) (shared.ParsedCron, error) {
	return cf(cron)
}

type RuleFunc func(input []int, interval shared.CronInterval) ([]int, error)

type CronParser struct {}

func (c CronParser) Parse(cron shared.Cron) (shared.ParsedCron, error) {
	output := shared.ParsedCron{}

	for k, v := range cron {
		values, err := process(v, k)

		if err != nil {
			return output, err
		}

		output[k] = values
	}

	return output, nil
}

func NewParser() CronParser {
	return CronParser{}
}