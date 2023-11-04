package shared

type CronInterval int

type Cron map[CronInterval]string

type ParsedCron map[CronInterval][]int