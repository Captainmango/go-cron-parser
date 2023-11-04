package shared

import "fmt"

var cronIntervalOrder = []CronInterval {
	MINUTE,
	HOUR,
	DAY_OF_MONTH,
	MONTH,
	DAY_OF_WEEK,
}

func NewCronFromArgs(input []string) (Cron, error) {
	inputLength := len(input)

	if inputLength != 6 {
		return nil, fmt.Errorf("could not create Cron from input %s", input)
	}

	cronArgs := input[1:]
	cron := Cron{}

	for i, str := range cronArgs {
		cron[cronIntervalOrder[i]] = str
	}

	return cron, nil
}