package main

import (
	"testing"
)

func TestRun(t *testing.T) {
	got := run(100, 5, "50 10 20 30 40")
	want := 90

	RunHelper(got, want, t)
}

func TestRun1(t *testing.T) {
	got := run(10, 2, "15 20")
	want := 0

	RunHelper(got, want, t)
}

func TestRun2(t *testing.T) {
	got := run(175, 7, "50 1 60 2 70 3 40")
	want := 170

	RunHelper(got, want, t)
}

func TestRun3(t *testing.T) {
	got := run(101, 7, "50 1 60 2 70 3 40")
	want := 100

	RunHelper(got, want, t)
}

func RunHelper(got, want int, t *testing.T) {
	t.Helper()

	if got != want {
		t.Errorf("Want %d, got %d", want, got)
	} else {
		t.Log("PASS - answer: ", got)
	}
}

func BenchmarkRun(b *testing.B) {
	for n := 0; n < b.N; n++ {
		run(100, 5, "50 10 20 30 40")
	}
}
