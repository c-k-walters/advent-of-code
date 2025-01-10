package y2022d05

import (
	"bufio"
	"os"
	"testing"
)

func TestDay05Part1(t *testing.T) {
    fileName := "inputs/input.txt"
    file, err := os.Open(fileName)
    if err != nil { t.Fatalf("Failed to open file: %s", fileName) }
    defer file.Close()

    expected := ""
    actual := RunDay05Pt1(bufio.NewScanner(file))
    if expected != actual {
        t.Fatalf("Expected %s and found %s", expected, actual)
    }
}


