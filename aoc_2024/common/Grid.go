package common

import (
	"bufio"
	"fmt"
)

type Grid struct {
    W, H int
    Tiles []rune
}

func NewGrid(scanner *bufio.Scanner) Grid {
    grid := Grid{}
    i := 0
    for scanner.Scan() {
        line := scanner.Text()
        grid.W = len(line) 
        grid.Tiles = append(grid.Tiles, []rune(line)...)
        grid.H++
        i++
    }
    return grid
}

func (this Grid) At(complex complex128) rune {
    x, y := int(real(complex)), int(imag(complex))
    if x < 0 || x >= this.W || y < 0 || y >= this.H {
        return -1
    }
    return this.Tiles[int(imag(complex)) * this.W + int(real(complex))]
}

func (this *Grid) Set(runic rune, complex complex128) {
    x, y := int(real(complex)), int(imag(complex))
    if x < 0 || x >= this.W || y < 0 || y >= this.H {
        return
    }
    this.Tiles[y * this.W + x] = runic
}

func (this Grid) Print() {
    for i, tile := range this.Tiles {
        if i % this.W == 0 { fmt.Print("\n") }
        fmt.Print(string(tile))
    }
    fmt.Print("\n")
}
