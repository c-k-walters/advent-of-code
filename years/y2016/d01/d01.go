package y2016d01

import (
	"AoC/common"
	"bufio"
	"math"
	"regexp"
	"strconv"
)


func RunDay01Pt1(scanner *bufio.Scanner) int {
    scanner.Scan()
    line := scanner.Text()
    moveExp := regexp.MustCompile("(L|R)\\d+")

    moves := moveExp.FindAllString(line, -1)

    // ^  >  v  <
    // -i 1  i  -1 
    var pos, dir complex64 = complex(0, 0), -1i
    
    for _, move := range moves {
        distStr := move[1:]
        dist, _ := strconv.Atoi(distStr)
        switch move[0] {
        case 'L':
            dir *= -1i
        case 'R':
            dir *= 1i
        }
        pos += dir * complex(float32(dist), 0)
    }
    return manhattanDist(pos)
}

func RunDay01Pt2(scanner *bufio.Scanner) int {
    scanner.Scan()
    line := scanner.Text()
    moveExp := regexp.MustCompile("(L|R)\\d+")

    moves := moveExp.FindAllString(line, -1)
    seen := common.Set[complex64]{}

    // ^  >  v  <
    // -i 1  i  -1 
    var pos, dir complex64 = complex(0, 0), -1i
    
    for _, move := range moves {
        distStr := move[1:]
        dist, _ := strconv.Atoi(distStr)
        switch move[0] {
        case 'L':
            dir *= -1i
        case 'R':
            dir *= 1i
        }
        last := pos
        for range dist {
            last += dir * complex(1, 0)
            if seen.Contains(last) { return manhattanDist(last) }
            seen.Add(last)
        }
        pos = last
    }
    return -1
}

func manhattanDist(vec complex64) int {
    dist := math.Abs(float64(real(vec))) + math.Abs(float64(imag(vec))) 
    return int(dist)
}
