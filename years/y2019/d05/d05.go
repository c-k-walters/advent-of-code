package y2019d05

import (
	"AoC/years/y2019/intcode"
	"bufio"
)

func RunDay05Pt1(scanner *bufio.Scanner) int {
    code := intcode.GetProgram(scanner)
    code.FeedInput(1)
    code.Run(225, 1)
    return code.GetLastOutput()
}

func RunDay05Pt2(scanner *bufio.Scanner) int {
    code := intcode.GetProgram(scanner)
    code.FeedInput(5)
    code.Run(225, 1)
    return code.GetLastOutput()
}
