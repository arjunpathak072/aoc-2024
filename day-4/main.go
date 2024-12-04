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

var directionsV1 = []index{
	index{r: 0, c: 1},
	index{r: 0, c: -1},
	index{r: 1, c: 0},
	index{r: -1, c: 0},
	index{r: -1, c: -1},
	index{r: 1, c: 1},
	index{r: 1, c: -1},
	index{r: -1, c: 1},
}

var directionsV2 = []index{
	index{r: -1, c: -1},
	index{r: 1, c: 1},
	index{r: 1, c: -1},
	index{r: -1, c: 1},
}

func main() {
	puzzleInput := parseInput("day-4.input")
	fmt.Println("Part One: ", partOne(puzzleInput))
	fmt.Println("Part Two: ", partTwo(puzzleInput))
}

func parseInput(fileName string) [][]rune {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var input [][]rune
	for scanner.Scan() {
		input = append(input, []rune(scanner.Text()))
	}
	return input
}

func partOne(grid [][]rune) (wordCount int) {
	const word = "XMAS"
	maxRow, maxCol := len(grid), len(grid[0])

	for r := range grid {
		for c := range grid[r] {
			if grid[r][c] == rune(word[0]) {
				for _, dir := range directionsV1 {
					findXMAS(grid, index{r: r, c: c}, dir, 0, word, &wordCount, maxRow, maxCol)
				}
			}
		}
	}

	return wordCount
}

func partTwo(grid [][]rune) (wordCount int) {
	const word = "MAS"
	maxRow, maxCol := len(grid), len(grid[0])

	validAIndices := make(map[index]bool)

	for r := range grid {
		for c := range grid[r] {
			if grid[r][c] == rune(word[0]) {
				for _, dir := range directionsV2 {
					if findMAS(grid, index{r: r, c: c}, dir, 0, word, maxRow, maxCol) {
						aIdx := index{r: r + dir.r, c: c + dir.c}
						if validAIndices[aIdx] {
							wordCount++
						}
						validAIndices[aIdx] = true
					}
				}
			}
		}
	}

	return wordCount
}

func findMAS(grid [][]rune, idx index, dir index, at int, word string, maxRow, maxCol int) bool {
	if at+1 == len(word) {
		return true
	}

	neigbourIndex := index{r: idx.r + dir.r, c: idx.c + dir.c}
	if isValidIndex(neigbourIndex, maxRow, maxCol) {
		neighbourRune := grid[neigbourIndex.r][neigbourIndex.c]
		if neighbourRune == rune(word[at+1]) {
			return findMAS(grid, neigbourIndex, dir, at+1, word, maxRow, maxCol)
		}
	}
	return false
}

func findXMAS(grid [][]rune, idx index, dir index, at int, word string, wordCount *int, maxRow, maxCol int) {
	if at+1 == len(word) {
		(*wordCount)++
		return
	}

	neigbourIndex := index{r: idx.r + dir.r, c: idx.c + dir.c}
	if isValidIndex(neigbourIndex, maxRow, maxCol) {
		neighbourRune := grid[neigbourIndex.r][neigbourIndex.c]
		if neighbourRune == rune(word[at+1]) {
			findXMAS(grid, neigbourIndex, dir, at+1, word, wordCount, maxRow, maxCol)
		}
	}
}

func isValidIndex(idx index, maxRow, maxCol int) bool {
	return idx.r >= 0 && idx.c >= 0 && idx.r < maxRow && idx.c < maxCol
}
