package y2022d06

import (
	"AoC/common"
	"bufio"
)

func RunDay06Pt1(scanner *bufio.Scanner) int {
    scanner.Scan()
    line := scanner.Text()

    marker := len(line)
    for i := 3; i < len(line); i++ {
        markerSet := common.ToSet([]rune(line[i-3:i+1]))
        if len(markerSet) != 4 { continue }

        marker = i+1
        break
    }
    return marker
}

func RunDay06Pt2(scanner *bufio.Scanner) int {
    scanner.Scan()
    line := scanner.Text()

    marker := len(line)
    for i := 13; i < len(line); i++ {
        markerSet := common.ToSet([]rune(line[i-13:i+1]))
        if len(markerSet) != 14 { continue }

        marker = i+1
        break
    }
    return marker
}
