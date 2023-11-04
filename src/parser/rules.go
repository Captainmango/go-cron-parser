package parser

import "github.com/captainmango/go-cron-parser/src"

func rangeRule(start, end int, interval src.CronInterval) []int {
	output := []int{}

	if start > end {
		return output
	}

	for i := start; i <= end; i++ {
		output = append(output, i)
	}

	return output
}

func divisorRule(num int, interval src.CronInterval) []int {
	output := []int{}

	start := getStart(interval)
	limit := getLimit(interval)

	if num > limit {
		return output
	}

	for i := start; i <= limit; i++ {
		if i%num == 0 {
			output = append(output, i)
		}
	}

	return output
}

func wildcardRule(interval src.CronInterval) []int {
	output := []int{}

	start := getStart(interval)
	limit := getLimit(interval)

	for i := start; i <= limit; i++ {
		output = append(output, i)
	}

	return output
}

func listRule(nums []int, interval src.CronInterval) []int {
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

func singleRule(num int, interval src.CronInterval) []int {
	limit := getLimit(interval)
	start := getStart(interval)

	if num > limit || num < start {
		return []int{}
	}
	
	return []int{num}
}

func getStart(interval src.CronInterval) int {
	switch interval {
	case src.MINUTE:
		return src.TIME_MIN
	case src.HOUR:
		return src.TIME_MIN
	default:
		return src.CAL_MIN
	}
}

func getLimit(interval src.CronInterval) int {
	return src.INTERVAL_MAX_VALUES[interval]
}
