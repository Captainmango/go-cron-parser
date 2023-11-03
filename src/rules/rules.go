package rules

import "github.com/captainmango/go-cron-parser/src"

func List(start, end int, interval src.CronInterval) []int {
	output := []int{}

	if start > end {
		return output
	}

	for i := start; i <= end; i++ {
		output = append(output, i)
	}

	return output
}

func Divisor(num int, interval src.CronInterval) []int {
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

func All(interval src.CronInterval) []int {
	output := []int{}

	start := getStart(interval)
	limit := getLimit(interval)

	for i:= start; i <= limit; i++ {
		output = append(output, i)
	}

	return output
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