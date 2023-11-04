package parser

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/captainmango/go-cron-parser/src/shared"
)

var listPattern = `[0-9]+,[0-9]+`
var divisorPattern = `^\*\/[0-9]+$`
var wildcardPattern = `^\*$`
var rangePattern = `^[0-9]+-[0-9]+$`
var singlePattern = `^[0-9]+$`

func process(input string, interval shared.CronInterval) []int {
	rule, inputNums, err := getRuleAndInput(input)

	if err != nil {
		panic(err)
	}

	output := rule(inputNums, interval)

	return output
}

func rangeRule(nums []int, interval shared.CronInterval) []int {
	output := []int{}
	start := nums[0]
	end := nums[1]

	if start > end {
		return output
	}

	for i := start; i <= end; i++ {
		output = append(output, i)
	}

	return output
}

func divisorRule(num []int, interval shared.CronInterval) []int {
	output := []int{}

	start := getStart(interval)
	limit := getLimit(interval)

	if num[0] > limit {
		return output
	}

	for i := start; i <= limit; i++ {
		if i%num[0] == 0 {
			output = append(output, i)
		}
	}

	return output
}

func wildcardRule(_nums []int, interval shared.CronInterval) []int {
	output := []int{}

	start := getStart(interval)
	limit := getLimit(interval)

	for i := start; i <= limit; i++ {
		output = append(output, i)
	}

	return output
}

func listRule(nums []int, interval shared.CronInterval) []int {
	output := []int{}
	limit := getLimit(interval)
	start := getStart(interval)

	for _, num := range nums {
		if num <= limit && num >= start {
			output = append(output, num)
		}
	}

	return output
}

func singleRule(num []int, interval shared.CronInterval) []int {
	limit := getLimit(interval)
	start := getStart(interval)

	val := num[0]

	if val > limit || val < start {
		return []int{}
	}

	return []int{val}
}

func getRuleAndInput(input string) (RuleFunc, []int, error) {
	patterns := map[string]RuleFunc{
		listPattern: listRule,
		divisorPattern: divisorRule,
		rangePattern: rangeRule,
		singlePattern: singleRule,
		wildcardPattern: wildcardRule,
	}

	for pattern, rule := range patterns {
		result, _ := regexp.MatchString(pattern, input)

		if result {
			output, err := getInput(input)

			if err != nil {
				return nil, nil, err
			}

			return rule, output, nil
		}
	}

	return nil, nil, fmt.Errorf("unable to find rule for input %s", input)
}

func getInput(input string) ([]int, error) {
	re := regexp.MustCompile("[0-9]+")
	res := re.FindAllString(input, -1)

	output := []int{}

	for _, str := range res {
		num, err := strconv.Atoi(str)

		if err != nil {
			return nil, fmt.Errorf("failed to convert value %s to integer", str)
		}

		output = append(output, num)
	}

	return output, nil
}

func getStart(interval shared.CronInterval) int {
	switch interval {
	case shared.MINUTE:
		return shared.TIME_MIN
	case shared.HOUR:
		return shared.TIME_MIN
	default:
		return shared.CAL_MIN
	}
}

func getLimit(interval shared.CronInterval) int {
	return shared.INTERVAL_MAX_VALUES[interval]
}
