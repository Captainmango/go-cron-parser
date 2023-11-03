package src

type CronInterval int

const (
	MINUTE CronInterval = iota
	HOUR
	DAY_OF_MONTH
	MONTH
	DAY_OF_WEEK
)

const (
	TIME_MIN int = 0
	CAL_MIN int = 1
)

var INTERVAL_MAX_VALUES = map[CronInterval]int {
	MINUTE : 59,
	HOUR : 23,
	DAY_OF_MONTH : 31,
	MONTH : 12,
	DAY_OF_WEEK : 7, 
}