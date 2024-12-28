package y2019d03

import (
	"bufio"
	"os"
	"testing"
)

// Test day 02 part 1 on actual input
func TestDay03Part1_actual(t *testing.T) {
    fileName := "inputs/input.txt"
    file, err := os.Open(fileName)
    if err != nil { t.Fatalf("Failed to open file: %s", fileName) }
    defer file.Close()

    expected := 651
    actual := RunDay03Pt1(bufio.NewScanner(file))
    if expected != actual {
        t.Fatalf("Expected %d and found %d", expected, actual)
    }
}


// Test day 02 part 1 on actual input
func TestDay03Part2_actual(t *testing.T) {
    fileName := "inputs/input.txt"
    file, err := os.Open(fileName)
    if err != nil { t.Fatalf("Failed to open file: %s", fileName) }
    defer file.Close()

    expected := 7534
    actual := RunDay03Pt2(bufio.NewScanner(file))
    if expected != actual {
        t.Fatalf("Expected %d and found %d", expected, actual)
    }
}
