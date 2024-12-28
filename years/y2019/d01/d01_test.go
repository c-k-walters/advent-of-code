package y2019d01

import (
	"bufio"
	"os"
	"testing"
)

// Test day 01 part 1 on actual input
func TestDay01Part1_actual(t *testing.T) {
    fileName := "inputs/input.txt"
    file, err := os.Open(fileName)
    if err != nil { t.Fatalf("Failed to open file: %s", fileName) }
    defer file.Close()

    expected := 3380731
    actual := RunDay01Pt1(bufio.NewScanner(file))
    if expected != actual {
        t.Fatalf("Expected %d and found %d", expected, actual)
    }
}

// Test day 01 part 2 on actual input
func TestDay01Part2_actual(t *testing.T) {
    fileName := "inputs/input.txt"
    file, err := os.Open(fileName)
    if err != nil { t.Fatalf("Failed to open file: %s", fileName) }
    defer file.Close()

    expected := 5068210
    actual := RunDay01Pt2(bufio.NewScanner(file))
    if expected != actual {
        t.Fatalf("Expected %d and found %d", expected, actual)
    }
}
