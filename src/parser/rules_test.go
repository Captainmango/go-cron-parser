package parser

import (
	"reflect"
	"testing"

	"github.com/captainmango/go-cron-parser/src/shared"
)

func TestRangeRule(t *testing.T) {
	t.Run("it can process a range of numbers", func(t *testing.T) {
		want := []int{1, 2, 3, 4, 5}
		got, _ := rangeRule([]int{1, 5}, shared.MINUTE)

		compareSlices(t, want, got)
	})

	t.Run("it can return errors", func(t *testing.T) {
		_, err := rangeRule([]int{60, 90}, shared.MINUTE)
		checkError(t, err)
	})
}

func TestDivisorRule(t *testing.T) {
	t.Run("it can process a divisor rule", func(t *testing.T) {
		want := []int{0, 15, 30, 45}
		got, _ := divisorRule([]int{15}, shared.MINUTE)

		compareSlices(t, want, got)
	})

	t.Run("it can return errors", func(t *testing.T) {
		_, err := divisorRule([]int{90}, shared.MINUTE)
		checkError(t, err)
	})
}

func TestAllRule(t *testing.T) {
	t.Run("prints all numbers for interval", func(t *testing.T) {
		want := []int{1, 2, 3, 4, 5, 6, 7}
		got, _ := wildcardRule([]int{}, shared.DAY_OF_WEEK)

		compareSlices(t, want, got)
	})
}

func TestListRule(t *testing.T) {
	t.Run("it prints individual numbers", func(t *testing.T) {
		want := []int {1,4,7}
		got, _ := listRule([]int{1,4,7}, shared.DAY_OF_WEEK)

		compareSlices(t, want, got)
	})

	t.Run("it can return errors", func(t *testing.T) {
		_, err := listRule([]int{90}, shared.MINUTE)
		checkError(t, err)
	})
}

func TestSingleRule(t *testing.T) {
	t.Run("it can print a single number", func(t *testing.T) {
		want := []int{2}
		got, _ := singleRule([]int{2}, shared.DAY_OF_WEEK)

		compareSlices(t, want, got)
	})

	t.Run("it can return errors", func(t *testing.T) {
		_, err := singleRule([]int{90}, shared.MINUTE)
		checkError(t, err)
	})
}

func compareSlices(t testing.TB, s1, s2 []int) {
	t.Helper()

	if !reflect.DeepEqual(s1, s2) {
		t.Errorf("slices are not equal. Wanted %v \n Got %v", s1, s2)
	}
}

func checkError(t testing.TB, err error) {
	t.Helper()

	if err == nil {
		t.Error("encountered error when none was expected.")
	}
}
