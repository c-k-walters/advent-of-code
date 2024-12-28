package y2019d06

import (
	"bufio"
	"os"
	"testing"
)

func TestDay06Pt1(t *testing.T) {
    fileName := "inputs/input.txt"
    file, err := os.Open(fileName)
    if err != nil { t.Fatalf("Failed to open file: %s", fileName) }
    defer file.Close()

    expected := 300598
    actual := RunDay06Pt1(bufio.NewScanner(file))
    if expected != actual {
        t.Fatalf("Expected %d and found %d", expected, actual)
    }
}

func TestDay06Pt2(t *testing.T) {
    fileName := "inputs/input.txt"
   // fileName := "inputs/example.txt"
    file, err := os.Open(fileName)
    if err != nil { t.Fatalf("Failed to open file: %s", fileName) }
    defer file.Close()

    expected := 0
    actual := RunDay06Pt2(bufio.NewScanner(file))
    if expected != actual {
        t.Fatalf("Expected %d and found %d", expected, actual)
    }
}
