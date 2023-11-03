package src

func List(start, end int, interval CronInterval) []int {
	output := []int{}

	if start > end {
		return output
	}

	for i := start; i <= end; i++ {
		output = append(output, i)
	}

	return output
}

func Divisor(num int, interval CronInterval) []int {
	output := []int{}

	start := getStart(interval)
	limit := getLimit(interval)

	for i := start; i < limit; i++ {
		if i % num == 0 {
			output = append(output, i)
		}
	}

	return output
}

func getStart(interval CronInterval) int {
	switch interval {
	case MINUTE:
		return TIME_MIN
	case HOUR:
		return TIME_MIN
	default:
		return CAL_MIN
	}
}

func getLimit(interval CronInterval) int {
	return INTERVAL_MAX_VALUES[interval]
}