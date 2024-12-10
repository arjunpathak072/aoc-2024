package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type index struct {
	r, c int
}

type bounds struct {
	row, col int
}

var directions = []index{
	index{-1, 0},
	index{0, -1},
	index{1, 0},
	index{0, 1},
	index{-1, 1},
	index{1, -1},
	index{-1, -1},
	index{1, 1},
}

func main() {
	puzzleInput, b := parseInput("day-8.input")
	fmt.Println("Part One: ", partOne(puzzleInput, b))
	fmt.Println("Part Two: ", partTwo(puzzleInput, b))
}

func parseInput(fileName string) (freqToAntennas map[rune][]index, b bounds) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	freqToAntennas = make(map[rune][]index)

	for r := 0; scanner.Scan(); r++ {
		for c, freq := range scanner.Text() {
			if freq != '.' {
				freqToAntennas[freq] = append(freqToAntennas[freq], index{r: r, c: c})
			}
		}
		b.col = len(scanner.Text())
		b.row = r
	}
	b.row += 1
	return
}

func partOne(freqToAntennas map[rune][]index, b bounds) int {
	validLocations := make(map[index]bool)
	for _, antennas := range freqToAntennas {
		for i := 0; i < len(antennas); i++ {
			for j := i + 1; j < len(antennas); j++ {
				an1, an2 := getAntinodes(antennas[i], antennas[j])
				if isWithinBounds(an1, b) {
					validLocations[an1] = true
				}
				if isWithinBounds(an2, b) {
					validLocations[an2] = true
				}
			}
		}
	}
	return len(validLocations)
}

func partTwo(freqToAntennas map[rune][]index, b bounds) int {
	validLocations := make(map[index]bool)

	for _, antennas := range freqToAntennas {
		for i := 0; i < len(antennas); i++ {
			for j := i + 1; j < len(antennas); j++ {
				dR := antennas[i].r - antennas[j].r
				dC := antennas[i].c - antennas[j].c

				nextIdx := index{r: antennas[i].r - dR, c: antennas[i].c - dC}
				for isWithinBounds(nextIdx, b) {
					validLocations[nextIdx] = true
					nextIdx.r -= dR
					nextIdx.c -= dC
				}

				nextIdx = index{r: antennas[j].r + dR, c: antennas[j].c + dC}
				for isWithinBounds(nextIdx, b) {
					validLocations[nextIdx] = true
					nextIdx.r += dR
					nextIdx.c += dC
				}
			}
		}
	}
	return len(validLocations)
}

func getAntinodes(i, j index) (first index, second index) {
	return index{2*i.r - j.r, 2*i.c - j.c}, index{2*j.r - i.r, 2*j.c - i.c}
}

func isWithinBounds(i index, b bounds) bool {
	return i.r >= 0 && i.c >= 0 && i.r < b.col && i.c < b.col
}

func abs(num int) int {
	if num < 0 {
		return -1 * num
	}
	return num
}
