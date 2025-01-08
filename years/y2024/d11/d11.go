package y2024d11

import (
	"bufio"
	"regexp"
    "strconv"
)

func RunDay11Pt1(scanner *bufio.Scanner) int {
    scanner.Scan()
    line := scanner.Text()

    nums := regexp.MustCompile("\\d+")
    stonesStrs := nums.FindAllString(line, -1)
    stones := make([]int, 0)
    for _, val := range stonesStrs {
        stone, _ := strconv.Atoi(val)
        stones = append(stones, stone)
    }

    return blinkAt(stones, 25)
}

func RunDay11Pt2(scanner *bufio.Scanner) int {
    scanner.Scan()
    line := scanner.Text()

    nums := regexp.MustCompile("\\d+")
    stonesStrs := nums.FindAllString(line, -1)
    stones := make([]int, 0)
    for _, val := range stonesStrs {
        stone, _ := strconv.Atoi(val)
        stones = append(stones, stone)
    }

    return blinkAt(stones, 75)
}

// Helper Functions: Stone Recursion

// Struct to hold a given state of stone expansion
type memo struct {
    stone int
    blinksLeft int
}

func blinkAt(stones []int, numBlinks int) int {
    results := map[memo]int{}

    totalStones := 0
    for _, stone := range stones {
        blinkAtStone(stone, numBlinks, &results)
        totalStones += results[memo{stone, numBlinks}]
    }

    return totalStones
}

// Recursive function to find the number of stones resulting after a number of 
// blinks for the current stone occur.
// Use memoization to reduce search space.
func blinkAtStone(stone int, blinksLeft int, results *map[memo]int) int {
    // Base case: we don't blink anymore, left with just the current stone
    if blinksLeft == 0 { return 1 }

    // check if we have seen this state before
    if result, ok := (*results)[memo{stone, blinksLeft}]; ok {
        return result
    }

    // else determine the number of stones resulting from the current state
    // by recursion
    var count int
    switch {
    case stone == 0:
        count = blinkAtStone(1, blinksLeft - 1, results)
    case hasEvenNumDigits(stone):
        strVal := strconv.Itoa(stone)
        left, right := strVal[:len(strVal)/2], strVal[len(strVal)/2:]
        leftStone, _ := strconv.Atoi(left)
        rightStone, _ := strconv.Atoi(right)
        count = blinkAtStone(leftStone, blinksLeft - 1, results) +
            blinkAtStone(rightStone, blinksLeft - 1, results)
    default:
        count = blinkAtStone(stone*2024, blinksLeft - 1, results)     
    }

    (*results)[memo{stone, blinksLeft}] = count
    return count
}

func hasEvenNumDigits(num int) bool {
    return len(strconv.Itoa(num)) % 2 == 0
}
