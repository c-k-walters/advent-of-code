package y2019d01

import (
	"bufio"
	"strconv"
)

// Day 01 Part 1:
// Find the total fuel required by the mass found in a given scanner.
func RunDay01Pt1(scanner *bufio.Scanner) int {
    totalFuel := 0

    for {
        mass, ok := getMass(scanner)
        if !ok { break }

        totalFuel += calcFuel(mass)
    }

    return totalFuel
}

// Day 01 Part 2:
// Find the total fuel required by the mass found in a given scanner.
// Make sure to account for the mass of the new fuel add.
func RunDay01Pt2(scanner *bufio.Scanner) int {
    totalFuel := 0

    for {
        mass, ok := getMass(scanner)
        if !ok { break }

        totalFuel += calcFuelForFuel(mass)
    }

    return totalFuel
}

// Fuel for a given mass, accounting for additional mass of the fuel itself.
func calcFuelForFuel(mass int) int {
    fuel := calcFuel(mass)
    if fuel < 0 { return 0 }
    fuel += calcFuelForFuel(fuel)
    return fuel
}

// Each line of the input is a mass value.
func getMass(scanner *bufio.Scanner) (int, bool) {
    ok := scanner.Scan()
    if !ok { return -1, false }
    line := scanner.Text()
    mass, _ := strconv.Atoi(line)
    return mass, true
}

// Fuel for a given mass is equal to floor(mass / 3) - 2.
func calcFuel(mass int) int {
    return mass / 3 - 2
}
