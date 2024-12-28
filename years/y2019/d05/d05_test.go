package y2019d05

import (
	"bufio"
	"os"
	"testing"
)

// Test day 05 part 1 on actual input
func TestDay05Part1_actual(t *testing.T) {
    fileName := "inputs/input.txt"
    file, err := os.Open(fileName)
    if err != nil { t.Fatalf("Failed to open file: %s", fileName) }
    defer file.Close()

    expected := 4887191
    actual := RunDay05Pt1(bufio.NewScanner(file))
    if expected != actual {
        t.Fatalf("Expected %d and found %d", expected, actual)
    }
}

// Test day 05 part 2 on actual input
func TestDay05Part2_actual(t *testing.T) {
    fileName := "inputs/input.txt"
    file, err := os.Open(fileName)
    if err != nil { t.Fatalf("Failed to open file: %s", fileName) }
    defer file.Close()

    expected := 0
    actual := RunDay05Pt2(bufio.NewScanner(file))
    if expected != actual {
        t.Fatalf("Expected %d and found %d", expected, actual)
    }
}
