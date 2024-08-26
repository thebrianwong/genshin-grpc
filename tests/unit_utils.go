package tests

import "testing"

func Compare(t *testing.T, expected any, actual any) {
	if expected != actual {
		t.Error("Expected:", expected, "| Got:", actual)
	}
}
