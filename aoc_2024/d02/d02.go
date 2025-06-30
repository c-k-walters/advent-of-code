package y2024d02

import (
	"bufio"
	"regexp"
	"strconv"
)

func RunDay02Pt1(scanner *bufio.Scanner) int {
    safeReports := 0
    numExp := regexp.MustCompile("[0-9]+")

    for scanner.Scan() {
        line := scanner.Text()
        numStrs := numExp.FindAllString(line, -1)
        nums := make([]int, len(numStrs))
        for ind, str := range numStrs { nums[ind], _ = strconv.Atoi(str) }

        if ok, _ := isValidReport(nums); ok { safeReports++ }
    }
    return safeReports
}

func RunDay02Pt2(scanner *bufio.Scanner) int {
    safeReports := 0
    numExp := regexp.MustCompile("[0-9]+")

    for scanner.Scan() {
        line := scanner.Text()
        numStrs := numExp.FindAllString(line, -1)
        nums := make([]int, len(numStrs))
        for ind, str := range numStrs { nums[ind], _ = strconv.Atoi(str) }

        ok, failedInd := isValidReport(nums)

        if ok {
            safeReports++
            continue
        }

        // retry report with damper
        for i := failedInd - 2; i <= failedInd + 1; i++ {
            if i < 0 || i >= len(nums) { continue }

            newReport := append([]int{}, nums[:i]...)
            if i+1 < len(nums) { newReport = append(newReport, nums[i+1:]...) }
            if ok, _ := isValidReport(newReport); ok {
                safeReports++
                break
            }
        }
    }
    return safeReports
}

func isValidReport(nums []int) (valid bool, failed int) {
    increasing := false
    decreasing := false 
    last := -1

    for ind, val := range nums {
        if last == -1 { last = val; continue }
        if val == last { failed = ind; return }
        if !increasing && !decreasing { 
            increasing = val > last
            decreasing = val < last
        }

        // setup of this loop is done, evaluate for rules
        diff := val - last
        if (increasing && diff < 0) || (decreasing && diff > 0) {
            failed = ind
            return
        }

        if diff < 0 { diff *= -1 }
        if diff > 3 { failed = ind; return }
        last = val
    }
    valid = true
    return
}
