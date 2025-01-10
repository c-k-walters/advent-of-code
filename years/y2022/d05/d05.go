package y2022d05

import (
	"AoC/common"
	"bufio"
)

func RunDay05Pt1(scanner *bufio.Scanner) string {
    stacks := make([][]rune, 10)
    for ind := range stacks { stacks[ind] = make([]rune, 100) }

    for i := 0; i < 8; i++ {
        scanner.Scan()
        line := scanner.Text()

        for j := 1; j < 10; j++ {
            stacks[j][7-i] = rune(line[1 + (j-1)*4])
        }
    }
    scanner.Scan(); scanner.Scan()

    for scanner.Scan() {
        line := scanner.Text()
        nums := common.ToInts(line)
        size, from, to := nums[0], nums[1], nums[2]

        grabbed := make([]rune, size)
        for ind, crate := range stacks[from] {
            if crate != ' ' { continue }

            for j := range size {
                
            }
        }

        // 32 is ' ' in each stack

    }

    return ""
}
