package y2022d01

import (
	"bufio"
	"sort"
	"strconv"
)

func RunDay01Pt1(scanner *bufio.Scanner) int {
    currentStore := 0
    mostStore := 0

    for scanner.Scan() {
        line := scanner.Text()
        if line == "" {
            if currentStore > mostStore { mostStore = currentStore }
            currentStore = 0
            continue
        }

        cals, _ := strconv.Atoi(line)
        currentStore += cals
    }
    return mostStore
}

func RunDay01Pt2(scanner *bufio.Scanner) int {
    currentStore := 0
    mostStores := []int{0, 0, 0, 0}

    for scanner.Scan() {
        line := scanner.Text()
        if line == "" {
            if currentStore > mostStores[1] {
                mostStores[0] = currentStore
                sort.Ints(mostStores)
            }
            currentStore = 0
            continue
        }

        cals, _ := strconv.Atoi(line)
        currentStore += cals
    }

    sum := 0
    for _, val := range mostStores[1:] { sum += val }
    return sum
}
