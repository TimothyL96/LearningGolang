package main

import (
	"strings"
	"testing"
)

func TestBuildTree(t *testing.T) {
	preArr := []int{1, 2, 4, 5, 14, 18, 6, 7, 8, 9, 13, 10, 11, 12, 14}
	inArr := []int{5, 4, 18, 14, 2, 8, 7, 6, 9, 1, 13, 10, 12, 11, 14}

	got := strings.TrimSpace(getPostOrder(buildTree(preArr, inArr)))
	want := "5 18 14 4 8 7 9 6 2 12 14 11 10 13 1"

	BuildTreeHelper(got, want, t)
}

func TestBuildTree1(t *testing.T) {
	var preArr []int
	var inArr []int

	got := strings.TrimSpace(getPostOrder(buildTree(preArr, inArr)))
	want := ""

	BuildTreeHelper(got, want, t)
}

func BuildTreeHelper(got, want string, t *testing.T) {
	t.Helper()

	if got != want {
		t.Errorf("\nWant:\t %s\nGot:\t %s", want, got)
	} else {
		t.Log("PASSED - Result: ", got)
	}
}
