package y2019d06

import (
	"AoC/common"
	"bufio"
	"fmt"
	"regexp"
)

func RunDay06Pt1(scanner *bufio.Scanner) int {
    orbits := common.NewGraph(common.StringHash)
    orbExp := regexp.MustCompile("([A-Z]|[0-9])+")
    for scanner.Scan() {
        line := scanner.Text()
        objects := orbExp.FindAllString(line, 2)
        
        orbits.AddVertex(objects[0])
        orbits.AddVertex(objects[1])
        orbits.AddEdge(objects[0], objects[1], 1)
    }

    totalOrbits := 0

    common.DFS(*orbits, "COM", func(val string, depth int) bool {
        totalOrbits += depth
        return false
    })

    return totalOrbits
}

func RunDay06Pt2(scanner *bufio.Scanner) int {
    orbits := common.NewGraph(common.StringHash)
    orbExp := regexp.MustCompile("([A-Z]|[0-9])+")
    for scanner.Scan() {
        line := scanner.Text()
        objects := orbExp.FindAllString(line, 2)
        
        orbits.AddVertex(objects[0])
        orbits.AddVertex(objects[1])
        orbits.AddEdge(objects[0], objects[1], 1)
    }

    lastCommon := 0
    for i := range SanPath {
        if SanPath[i] != YouPath[i] { 
            break
        }
        lastCommon = i
    }

    fmt.Println(SanPath)
    fmt.Println(YouPath)

    // Subtract 4 from answer to account for common parent and child itself
    // in both lists.
    return len(SanPath[lastCommon:]) + len(YouPath[lastCommon+1:]) - 3
}
