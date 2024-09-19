package main

import (
	"reflect"
	"testing"
)

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2, 3, 4, 5}, []int{5, 3, 2, 0}, []int{2, 3, 5, 6})
	want := []int{15, 10, 16}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q but want %q", got, want)
	}
}
