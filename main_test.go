package main

import (
	"testing"
)

func TestFun(t *testing.T) {
	// 1-3 pm profit $20
	// 6-9 pm profit $60

	startTime := []int{1, 2, 3, 4, 6}
	endTime := []int{3, 5, 10, 6, 9}
	profit := []int{20, 20, 100, 70, 60}

	out := jobScheduling(startTime, endTime, profit)
	want := 150
	if out != want {
		t.Errorf("got %d, want %d", out, want)
	}
}
