package src_test

import (
	"reflect"
	"testing"

	"github.com/captainmango/go-cron-parser/src"
)

func TestListRule(t *testing.T) {
	t.Run("it can process a list of numbers", func(t *testing.T) {
		want := []int{1,2,3,4,5}
		got := src.List(1, 5, src.MINUTE)

		compareSlices(t, want, got)
	})

	t.Run("it returns empty slice if numbers don't follow in series", func(t *testing.T) {
		want := []int{}
		got := src.List(5, 1, src.MINUTE)

		compareSlices(t, want, got)
	})
}

func TestDivisorRule(t *testing.T) {
	t.Run("it can process a divisor rule", func(t *testing.T) {
		want := []int{0,15,30,45}
		got := src.Divisor(15, src.MINUTE)

		compareSlices(t, want, got)
	})
}

func compareSlices(t testing.TB, s1, s2 []int) {
	t.Helper()

	if !reflect.DeepEqual(s1, s2) {
		t.Errorf("arrays are not equal. Wanted %v \n Got %v", s1, s2)
	}
}