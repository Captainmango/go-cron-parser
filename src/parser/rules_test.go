package parser

import (
	"reflect"
	"testing"

	"github.com/captainmango/go-cron-parser/src"
)

func TestRangeRule(t *testing.T) {
	t.Run("it can process a range of numbers", func(t *testing.T) {
		want := []int{1, 2, 3, 4, 5}
		got := rangeRule(1, 5, src.MINUTE)

		compareSlices(t, want, got)
	})

	t.Run("it returns empty slice if numbers don't follow in series", func(t *testing.T) {
		want := []int{}
		got := rangeRule(5, 1, src.MINUTE)

		compareSlices(t, want, got)
	})
}

func TestDivisorRule(t *testing.T) {
	t.Run("it can process a divisor rule", func(t *testing.T) {
		want := []int{0, 15, 30, 45}
		got := divisorRule(15, src.MINUTE)

		compareSlices(t, want, got)
	})

	t.Run("it doesn't break if divisor doesn't fit in interval", func(t *testing.T) {
		want := []int{}
		got := divisorRule(88, src.HOUR)

		compareSlices(t, want, got)
	})
}

func TestAllRule(t *testing.T) {
	t.Run("prints all numbers for interval", func(t *testing.T) {
		want := []int{1, 2, 3, 4, 5, 6, 7}
		got := wildcardRule(src.DAY_OF_WEEK)

		compareSlices(t, want, got)
	})
}

func TestListRule(t *testing.T) {
	t.Run("it prints individual numbers", func(t *testing.T) {
		want := []int {1,4,7}
		got := listRule([]int{1,4,7}, src.DAY_OF_WEEK)

		compareSlices(t, want, got)
	})
}

func TestSingleRule(t *testing.T) {
	t.Run("it can print a single number", func(t *testing.T) {
		want := []int{2}
		got := singleRule(2, src.DAY_OF_WEEK)

		compareSlices(t, want, got)
	})
}

func compareSlices(t testing.TB, s1, s2 []int) {
	t.Helper()

	if !reflect.DeepEqual(s1, s2) {
		t.Errorf("slices are not equal. Wanted %v \n Got %v", s1, s2)
	}
}
