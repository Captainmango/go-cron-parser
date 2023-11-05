package shared

import "fmt"

func NewCronFromArgs(input []string) (Cron, error) {
	inputLength := len(input)

	if inputLength != 6 {
		return nil, fmt.Errorf("could not create Cron from input %s", input)
	}

	cronArgs := input[1:]
	cron := Cron{}

	for i, str := range cronArgs {
		cron[CronIntervalOrder[i]] = str
	}

	return cron, nil
}