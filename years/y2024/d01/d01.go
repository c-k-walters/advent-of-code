package y2024d01

import (
	"bufio"
	"regexp"
	"sort"
	"strconv"
)

func RunDay01Pt1(scanner *bufio.Scanner) int {
    left, right := []int{}, []int{}
    listExp := regexp.MustCompile("[0-9]+")
    
    for scanner.Scan() {
        line := scanner.Text()
        vals := listExp.FindAllString(line, 2)
        if len(vals) != 2 { return -1 }

        l, _ := strconv.Atoi(vals[0])
        r, _ := strconv.Atoi(vals[1])

        left = append(left, l)
        right = append(right, r)
    }

    sort.Ints(left)
    sort.Ints(right)

    dist := 0
    for i := 0; i < len(left); i++ {
        l, r := left[i], right[i]
        if r > l { dist += r-l } else if l > r { dist += l-r }
    }
    return dist
}

func RunDay01Pt2(scanner *bufio.Scanner) int {
    left, right := []int{}, map[int]int{}
    listExp := regexp.MustCompile("[0-9]+")
    
    for scanner.Scan() {
        line := scanner.Text()
        vals := listExp.FindAllString(line, 2)
        if len(vals) != 2 { return -1 }

        l, _ := strconv.Atoi(vals[0])
        r, _ := strconv.Atoi(vals[1])

        left = append(left, l)
        right[r] += 1
    }

    similarityScore := 0
    for _, num := range left {
        similarityScore += num * right[num] 
    }

    return similarityScore
}
