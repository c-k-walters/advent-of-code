package y2019d07

import (
	"AoC/years/y2019/intcode"
	"bufio"
	"regexp"
	"strconv"
)

func RunDay07Pt1(scanner *bufio.Scanner) int {
    scanner.Scan()
    line := scanner.Text()
    numExp := regexp.MustCompile("\\d+")
    numStrs := numExp.FindAllString(line, -1)
    mem := make([]int, len(numStrs))
    for ind, str := range numStrs { mem[ind], _ = strconv.Atoi(str) }

    // Find largest output for a given combination of input phase settings
    options := permutations([]int{0, 1, 2, 3, 4})
    largestOutput := 0

    for _, option := range options {
        amps := make([]intcode.Intcode, 5)

        for ind := range amps {
            amps[ind] = intcode.Intcode{Memory: mem}
            amps[ind].FeedInput(option[ind])
        }

        input := 0
        for _, ic := range amps {
            ic.FeedInput(input)
            ic.RunWithoutInput()
            input = ic.GetLastOutput()
        }

        if input > largestOutput { largestOutput = input }
    }
    return largestOutput
}

func RunDay07Pt2(scanner *bufio.Scanner) int {
    scanner.Scan()
    line := scanner.Text()
    numExp := regexp.MustCompile("\\d+")
    numStrs := numExp.FindAllString(line, -1)
    mem := make([]int, len(numStrs))
    for ind, str := range numStrs { mem[ind], _ = strconv.Atoi(str) }

    // Find largest output for a given combination of input phase settings
    options := permutations([]int{5, 6, 7, 8, 9})
    largestOutput := 0

    for _, option := range options {
        amps := make([]intcode.Intcode, 5)

        for ind := range amps {
            amps[ind] = intcode.Intcode{Memory: mem}
            amps[ind].FeedInput(option[ind])
        }

        input := 0
        finished := false
        for !finished {
            for _, ic := range amps {
                ic.FeedInput(input)
                ic.RunWithoutInput()
                next, err := ic.PopOutput()
                if err != nil { finished = true }
                input = next
            }
            amps[0].FeedInput(input)
        }
        if input > largestOutput { largestOutput = input }
    }
    return largestOutput
}

func permutations(nums []int) [][]int {
    var result [][]int
    permuteHelper(nums, []int{}, &result)
    return result
}

func permuteHelper(nums, current []int, result *[][]int) {
    if len(nums) == 0 {
        permutation := make([]int, len(current))
        copy(permutation, current)
        *result = append(*result, permutation)
        return
    }

    for ind, num := range nums {
        subset := append([]int{}, nums[:ind]...)
        if ind < len(nums) { subset = append(subset, nums[ind+1:]...) }

        permuteHelper(subset, append(current, num), result)
    }
}
