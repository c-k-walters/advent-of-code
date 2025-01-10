package common

import (
	"regexp"
	"strconv"
)

// search through string and pull out all int values
func ToInts(str string) []int {
    numExp := regexp.MustCompile("\\d+")
    strs := numExp.FindAllString(str, -1)
    nums := make([]int, len(strs))

    for ind, s := range strs {
        nums[ind], _ = strconv.Atoi(s)
    }

    return nums
}
