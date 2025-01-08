package y2022d02

import (
	"bufio"
	"fmt"
)

func RunDay02Pt1(scanner *bufio.Scanner) int {
    score := 0

    for scanner.Scan() {
        line := scanner.Text()

        score += int(line[2] - 'X') + 1

        roundResult := int(line[2] - 'X') - int(line[0] - 'A')
        if roundResult == -2 { roundResult = 1 }
        if roundResult == 2 { roundResult = -1 }
        score += (roundResult * 3) + 3
    }

    return score
}

func RunDay02Pt2(scanner *bufio.Scanner) int {
    score := 0

    for scanner.Scan() {
        line := scanner.Text()

        // X lose, Y draw, Z win - for 0, 3, 6 points we use offset to calc
        roundResult := int(line[2] - 'X') - 1
        score += roundResult * 3 + 3

        // points for shape made based on opponent shape
        roundShape := int(line[0] - 'A') + 1
        roundShape += roundResult
        if roundShape <= 0 { roundShape += 3 }
        if roundShape == 4 { roundShape -= 3 }
        score += roundShape
    }
    return score
}
