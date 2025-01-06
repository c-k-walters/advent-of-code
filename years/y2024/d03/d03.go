package y2024d03

import (
	"bufio"
	"regexp"
	"strconv"
)

func RunDay03Pt1(scanner *bufio.Scanner) int {
    sumOfProds := 0

    input := ""
    for scanner.Scan() {
        line := scanner.Text()
        input += line
    }

    mulExp := regexp.MustCompile("mul\\(\\d{1,3},\\d{1,3}\\)")
    mulExpressions := mulExp.FindAllString(input, -1)
    numExp := regexp.MustCompile("\\d+")

    for _, expression := range mulExpressions {
        numStrs := numExp.FindAllString(expression, 2)
        left, _ := strconv.Atoi(numStrs[0])
        right, _ := strconv.Atoi(numStrs[1])

        sumOfProds += left * right
    }
    
    return sumOfProds
}

func RunDay03Pt2(scanner *bufio.Scanner) int {
    sumOfProds := 0

    input := ""
    for scanner.Scan() {
        line := scanner.Text()
        input += line
    }

    condMul := regexp.MustCompile("don't\\(\\)|do\\(\\)|mul\\(\\d{1,3},\\d{1,3}\\)")
    mulExpressions := condMul.FindAllString(input, -1)
    numExp := regexp.MustCompile("\\d+")

    enabled := true
    for _, expression := range mulExpressions {
        switch {
        case expression == "don't()":
            enabled = false
            continue
        case expression == "do()":
            enabled = true
            continue
        }
        if enabled {
            numStrs := numExp.FindAllString(expression, 2)
            left, _ := strconv.Atoi(numStrs[0])
            right, _ := strconv.Atoi(numStrs[1])
            sumOfProds += left * right
        }
    }
    
    return sumOfProds
}
