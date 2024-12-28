package y2019d04

import (
	"testing"
)

// Test day 04 part 1 on actual input
func TestDay04Part1_actual(t *testing.T) {
    expected := 1767
    actual := RunDay04Pt1(145852, 616942)
    if expected != actual {
        t.Fatalf("Expected %d and found %d", expected, actual)
    }
}

// Test day 04 part 1 on actual input
func TestDay04Part2_actual(t *testing.T) {
    expected := 1192
    actual := RunDay04Pt2(145852, 616942)
    if expected != actual {
        t.Fatalf("Expected %d and found %d", expected, actual)
    }
}
