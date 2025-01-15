package y2022d03

import (
	"AoC/common"
	"bufio"
)

func RunDay03Pt1(scanner *bufio.Scanner) int {
    prioritySum := 0

    for scanner.Scan() {
        ruckSack := scanner.Text()
        left, right := ruckSack[:len(ruckSack)/2], ruckSack[len(ruckSack)/2:]
        seen := common.ToSet([]rune(left))
        for _, item := range right {
            if seen.Contains(item) {
                if item >= 'a' && item <= 'z' {
                    prioritySum += int(item - 'a') + 1
                } else {
                    prioritySum += int(item - 'A') + 27
                }
                break
            }
        }
    }
    return prioritySum
}

func RunDay03Pt2(scanner *bufio.Scanner) int {
    prioritySum := 0
    for scanner.Scan() {
        first := common.ToSet([]rune(scanner.Text()))
        scanner.Scan()
        second := common.ToSet([]rune(scanner.Text()))
        scanner.Scan()
        third := common.ToSet([]rune(scanner.Text()))
        
        badge := first.Intersection(second).Intersection(third)
        for item := range badge {
            if item >= 'a' && item <= 'z' {
                prioritySum += int(item - 'a') + 1
            } else {
                prioritySum += int(item - 'A') + 27
            }
            break
        }
    }
    return prioritySum
}
