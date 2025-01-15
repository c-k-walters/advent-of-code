package y2016d02

import (
	"AoC/common"
	"bufio"
)

func RunDay02Pt1(scanner *bufio.Scanner) int {
    code := 0

    keypad := map[common.Vec2]int {
        {X: 0, Y: 0}: 1,
        {X: 1, Y: 0}: 2,
        {X: 2, Y: 0}: 3,
        {X: 0, Y: 1}: 4,
        {X: 1, Y: 1}: 5,
        {X: 2, Y: 1}: 6,
        {X: 0, Y: 2}: 7,
        {X: 1, Y: 2}: 8,
        {X: 2, Y: 2}: 9,
    }

    pos := common.Vec2{X:1, Y:1}
    for scanner.Scan() {
        line := scanner.Text()
        for _, runic := range line {
            switch runic {
                case 'D':
                    pos.Y++
                case 'U':
                    pos.Y--
                case 'L':
                    pos.X--
                case 'R':
                    pos.X++
            }
            if pos.X < 0 { pos.X = 0 }
            if pos.X > 2 { pos.X = 2 }
            if pos.Y < 0 { pos.Y = 0 }
            if pos.Y > 2 { pos.Y = 2 }
        }
        code *= 10
        code += keypad[pos]
    }

    return code
}

func RunDay02Pt2(scanner *bufio.Scanner) string {
    code := ""
    return code
}
