package y2024d06

import (
	"AoC/common"
	"bufio"
)

func RunDay06Pt1(scanner *bufio.Scanner) int {
    lab := common.NewGrid(scanner)
    // mapping (real imag) -> (x, y)
    // direction uses i, 1, -1, -i
    gPos, gDir := findGuard(lab)

    tilesTraversed := common.Set[complex128]{}

    for {
        if imag(gPos) < 0 || int(imag(gPos)) >= lab.H ||
        real(gPos) < 0 || int(real(gPos)) >= lab.W {
            break
        }
        tilesTraversed.Add(gPos)

        nexPos := gPos + gDir
        nextRune := lab.At(nexPos) 

        switch nextRune {
        case '#':
            gDir = gDir * 1i
        default:
            gPos = nexPos
        }
    }

    return len(tilesTraversed)
}

func findGuard(grid common.Grid) (pos, dir complex128) {
    for ind, runic := range grid.Tiles {
        if runic == '<' || runic == '>' || runic == '^' || runic == 'v' {
            // pos = x + y(i)
            pos = complex(float64(ind % grid.W), float64(ind / grid.H))
            // ^  >  v  <
            // -i 1  i  -1 
            switch runic {
            case '>':
                dir = 1
            case '^':
                dir = -1i
            case '<':
                dir = -1
            case 'v':
                dir = 1i
            }
        }
    }
    return
}
