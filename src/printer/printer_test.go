package printer_test

import (
	"bytes"
	"testing"

	"github.com/captainmango/go-cron-parser/src/printer"
	"github.com/captainmango/go-cron-parser/src/shared"
)

func TestPrinter(t *testing.T) {
	t.Run("it prints to std out", func(t *testing.T) {
		input := shared.ParsedCron{
			shared.DAY_OF_WEEK: []int{1,3,5},
		}

		stdout := &bytes.Buffer{}

		printer := printer.NewPrinter()
		printer.Print(input, stdout)

		expectedOut := "day_of_week | [1 3 5] \n"
		compareOutput(t, expectedOut, stdout.String())
	})
}

func compareOutput(t testing.TB, want, got string) {
	t.Helper()

	if want != got {
		t.Errorf("output does not match expected. \n Got %s \n Want %s", got, want)
	}
}