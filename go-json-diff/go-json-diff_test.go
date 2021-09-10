package gojsondiff

import (
	"os"
	"testing"
)

func TestContains(t *testing.T) {
	entries := []struct {
		slice []interface{}
		value interface{}
		want  bool
	}{
		{[]interface{}{"a", "b", "c"}, "a", true},
		{[]interface{}{"a", "b", "c"}, "d", false},
		{[]interface{}{1, "a", "3"}, "1", false},
		{[]interface{}{1, "a", "3"}, 1, true},
	}

	for _, entry := range entries {
		got := contains(entry.slice, entry.value)

		if got != entry.want {
			t.Errorf("Expected '%v', but got %v", entry.want, got)
		}
	}
}

func TestGetFileSize(t *testing.T) {
	entries := []struct {
		path string
		want int64
	}{
		{"a.json", 136},
		{"b.json", 57},
	}

	for _, entry := range entries {
		file, err := os.Open("../testdata/" + entry.path)
		if err != nil {
			t.Fatalf("Error loading file: %s", err)
		}
		defer file.Close()

		got := getFileSize(file)

		if got != entry.want {
			t.Errorf("Expected file to have length '%v', but got %v", entry.want, got)
		}
	}
}

func TestCompare(t *testing.T) {
	entries := []struct {
		pathA string
		pathB string
		want  bool
	}{
		{"a.json", "b.json", false},
		{"b.json", "c.json", true},
		{"a.json", "d.json", true},
	}

	for _, entry := range entries {
		fileA, _ := os.Open("../testdata/" + entry.pathA)
		fileB, _ := os.Open("../testdata/" + entry.pathB)

		defer fileA.Close()
		defer fileB.Close()

		comparator := NewComparator(fileA, fileB)
		got := comparator.Compare()

		if got != entry.want {
			t.Errorf("Should have returned '%v', but returned '%v'", entry.want, got)
		}
	}
}
