package main

import (
	"os"
	"testing"
)

func TestParseDiff(t *testing.T) {

	bytes, err := os.ReadFile("testdata/diff.txt")

	if err != nil {
		t.Fatal(err)
	}

	diffs, err := parseDiff(string(bytes))

	if err != nil {
		t.Fatal(err)
	}

	expectedDiffsLen := 45

	if len(diffs) != expectedDiffsLen {
		t.Fatalf("expected len of diffs is %d, but actual value is %d", expectedDiffsLen, len(diffs))
	}

}