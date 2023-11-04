package parser_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/captainmango/go-cron-parser/src/parser"
	"github.com/captainmango/go-cron-parser/src/shared"
)

func TestParser(t *testing.T) {
	parser := parser.NewParser()

	t.Run("can parse a cron struct", func(t *testing.T) {
		cron := shared.Cron{
			shared.DAY_OF_WEEK: "1",
		}

		want := shared.ParsedCron{
			shared.DAY_OF_WEEK: []int{1}, 
		}
		got, _ := parser.Parse(cron)

		compareCrons(t, want, got)
	})

	t.Run("it can return errors", func(t *testing.T) {
		cron := shared.Cron{
			shared.DAY_OF_WEEK: "99",
		}

		_, err := parser.Parse(cron)

		checkError(t, err)
	})

	testCases := []struct{
		desc string
		input shared.Cron
		output shared.ParsedCron
	} {
		{
			desc: "wildcard rule",
			input: shared.Cron{shared.DAY_OF_WEEK: "*"},
			output: shared.ParsedCron{shared.DAY_OF_WEEK: []int{1,2,3,4,5,6,7}},
		},
		{
			desc: "divisor rule",
			input: shared.Cron{shared.DAY_OF_WEEK: "*/2"},
			output: shared.ParsedCron{shared.DAY_OF_WEEK: []int{2,4,6}},
		},
		{
			desc: "range rule",
			input: shared.Cron{shared.DAY_OF_WEEK: "1-3"},
			output: shared.ParsedCron{shared.DAY_OF_WEEK: []int{1,2,3}},
		},
		{
			desc: "list rule",
			input: shared.Cron{shared.DAY_OF_WEEK: "1,4"},
			output: shared.ParsedCron{shared.DAY_OF_WEEK: []int{1,4}},
		},
		{
			desc: "single rule",
			input: shared.Cron{shared.DAY_OF_WEEK: "1"},
			output: shared.ParsedCron{shared.DAY_OF_WEEK: []int{1}},
		},
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("it correctly parse a %s", testCase.desc), func(t *testing.T) {
			got, _ := parser.Parse(testCase.input)
			want := testCase.output

			compareCrons(t, got, want)
		})
	}
	
}

func compareCrons(t testing.TB, want, got shared.ParsedCron) {
	t.Helper()

	if !reflect.DeepEqual(want, got) {
		t.Errorf("cron did not parse correctly. got %v, want %v", got, want)
	}
}

func checkError(t testing.TB, err error) {
	t.Helper()

	if err == nil {
		t.Error("encountered error when none was expected.")
	}
}