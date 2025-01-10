package y2022d04

import (
	"AoC/common"
	"bufio"
)


func RunDay04Pt1(scanner *bufio.Scanner) int {
    overlappingPairs := 0

    for scanner.Scan() {
        pair := scanner.Text()
        elf1, elf2 := make([]int, 2), make([]int, 2)
        nums := common.ToInts(pair)
        elf1[0] = nums[0]
        elf1[1] = nums[1]
        elf2[0] = nums[2]
        elf2[1] = nums[3]

        if elf1[0] <= elf2[0] && elf1[1] >= elf2[1] ||
        elf2[0] <= elf1[0] && elf2[1] >= elf1[1] {
            overlappingPairs++
        }
    }
    return overlappingPairs
}

func RunDay04Pt2(scanner *bufio.Scanner) int {
    overlappingPairs := 0

    for scanner.Scan() {
        pair := scanner.Text()
        elf1, elf2 := make([]int, 2), make([]int, 2)
        nums := common.ToInts(pair)
        elf1[0] = nums[0]
        elf1[1] = nums[1]
        elf2[0] = nums[2]
        elf2[1] = nums[3]

        elf1Area := common.Set[int]{}
        for i := elf1[0]; i <= elf1[1]; i++ { elf1Area.Add(i) }

        for j := elf2[0]; j <= elf2[1]; j++ {
            if elf1Area.Contains(j) { overlappingPairs++; break }
        }
    }
    return overlappingPairs
}
