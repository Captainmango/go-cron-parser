package printer

import (
	"fmt"
	"io"

	"github.com/captainmango/go-cron-parser/src/shared"
)

type CronPrinterInterface interface {
	Print(parsedCron shared.ParsedCron, w io.Writer) 
}

type CronPrinter struct {}

func (c CronPrinter) Print(parsedCron shared.ParsedCron, w io.Writer) {
	for interval, values := range parsedCron {
		lineContent := fmt.Sprintf("%15s | %v \n", getIntervalName(interval), values)
		w.Write([]byte(lineContent))
	}
}

func getIntervalName(interval shared.CronInterval) string {
	return shared.INTERVAL_NAMES[interval]
}

func NewPrinter() CronPrinter {
	return CronPrinter{}
}