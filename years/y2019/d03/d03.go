package y2019d03

import (
	"AoC/common"
	"bufio"
	"math"
	"regexp"
	"strconv"
)

// Day 03 Part 1:
// Find the nearest intersection to the origin of a give wire by manhattan distance.
func RunDay03Pt1(scanner *bufio.Scanner) int {
    scanner.Scan()
    wire1Str := scanner.Text()
    scanner.Scan()
    wire2Str := scanner.Text()

    wire1 := common.ToSet(scanWire(wire1Str))
    wire2 := common.ToSet(scanWire(wire2Str))

    intersections := wire1.Intersection(wire2)
    
    minDist := math.MaxInt
    for xsect := range intersections {
        if dist := xsect.ManhattanDist(); dist < minDist {
            minDist = dist
        }
    }

    return minDist
}

// Day 03 Part 2:
// Find the fewest combined steps the wires must take to reach an intersection.
func RunDay03Pt2(scanner *bufio.Scanner) int {
    scanner.Scan()
    wire1Str := scanner.Text()
    scanner.Scan()
    wire2Str := scanner.Text()

    wire1Steps := scanWire(wire1Str)
    wire1 := common.ToSet(wire1Steps)
    wire2Steps := scanWire(wire2Str)
    wire2 := common.ToSet(wire2Steps)

    intersections := wire1.Intersection(wire2)
    
    minSteps := math.MaxInt
    for xsect := range intersections {
        stepCount := 0
        for ind, step := range wire1Steps {
            if common.IsEqual(step, xsect) {
                stepCount += ind + 1
                break
            }
        }
        for ind, step := range wire2Steps {
            if common.IsEqual(step, xsect) {
                stepCount += ind + 1
                break
            }
        }
        if stepCount < minSteps {
            minSteps = stepCount
        }
    }
    return minSteps
}

func scanWire(line string) []common.Vec2 { 
    steps := []common.Vec2{}
    pathExp := regexp.MustCompile("(U|D|R|L)[0-9]+")
    path := pathExp.FindAllString(line, -1)

    curr := common.Vec2{}
    for _, edge := range path {
        dir := edge[0] 
        dirVec := common.Vec2{}
        switch dir {
        case 'D':
            dirVec.X, dirVec.Y = 0, 1
        case 'U':
            dirVec.X, dirVec.Y = 0, -1
        case 'R':
            dirVec.X, dirVec.Y = 1, 0
        case 'L':
            dirVec.X, dirVec.Y = -1, 0
        }

        length, _ := strconv.Atoi(edge[1:])
        if length == 0 { continue }
        for range length {
            curr = common.Add(curr, dirVec)
            steps = append(steps, curr)
        }
    }

    return steps
}
