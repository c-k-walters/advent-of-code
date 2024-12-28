package y2019d02

import (
	"AoC/years/y2019/intcode"
	"bufio"
)

// Day 02 Part 1:
func RunDay02Pt1(scanner *bufio.Scanner) int {
    intcode := intcode.GetProgram(scanner)
    return intcode.Run(12, 2)
}

// Day 02 Part 2:
func RunDay02Pt2(scanner *bufio.Scanner) int {
    comp := intcode.GetProgram(scanner)

    for i := range 100 {
        for j := range 100 {
            copy := intcode.Intcode{}
            copy.Memory = append(copy.Memory, comp.Memory...)
            output := copy.Run(i, j)
            if output == 19690720 { return 100*i + j }
        }
    }
    return -1
}
