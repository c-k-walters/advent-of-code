package y2019d02

import (
	"bufio"
	"os"
	"testing"
)

// Test day 02 part 1 on actual input
func TestDay02Part1_actual(t *testing.T) {
    fileName := "inputs/input.txt"
    file, err := os.Open(fileName)
    if err != nil { t.Fatalf("Failed to open file: %s", fileName) }
    defer file.Close()

    expected := 4945026
    actual := RunDay02Pt1(bufio.NewScanner(file))
    if expected != actual {
        t.Fatalf("Expected %d and found %d", expected, actual)
    }
}

// Test day 02 part 2 on actual input
func TestDay02Part2_actual(t *testing.T) {
    fileName := "inputs/input.txt"
    file, err := os.Open(fileName)
    if err != nil { t.Fatalf("Failed to open file: %s", fileName) }
    defer file.Close()

    expected := 5296
    actual := RunDay02Pt2(bufio.NewScanner(file))
    if expected != actual {
        t.Fatalf("Expected %d and found %d", expected, actual)
    }
}
