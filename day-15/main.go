package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
)

const (
	wall     = '#'
	box      = 'O'
	space    = '.'
	robot    = '@'
	boxOpen  = '['
	boxClose = ']'
	left     = '<'
	right    = '>'
	up       = '^'
	down     = 'v'
)

type index struct {
	r, c int
}

var moveToDirection = map[rune]index{
	'<': index{r: 0, c: -1},
	'v': index{r: 1, c: 0},
	'^': index{r: -1, c: 0},
	'>': index{r: 0, c: 1},
}

var dirToMove = map[index]rune{
	index{r: 0, c: -1}: '<',
	index{r: 1, c: 0}: 'v',
	index{r: -1, c: 0}: '^',
	index{r: 0, c: 1}: '>',
}

func main() {
	grid, moves, sIdx := parseInput("input", false)
	fmt.Println("Part One:", partOne(grid, moves, sIdx))
	grid, moves, sIdx = parseInput("example", true)
	fmt.Println("Part Two:", partTwo(grid, moves, sIdx))
}

func parseInput(ext string, partTwo bool) (grid [][]rune, moves []rune, sIdx index) {
	gridFile, err := os.Open("day-15-grid." + ext)
	if err != nil {
		log.Fatal(err)
	}
	defer gridFile.Close()

	scanner := bufio.NewScanner(gridFile)

	if !partTwo {
		for scanner.Scan() {
			grid = append(grid, []rune(scanner.Text()))
			if c := slices.Index(grid[len(grid)-1], robot); c != -1 {
				r := len(grid) - 1
				sIdx = index{r: r, c: c}
				grid[r][c] = space
			}
		}
	} else {
		for scanner.Scan() {
			currRow := make([]rune, 0, len(scanner.Text())*2)
			for i, char := range scanner.Text() {
				switch char {
				case wall:
					currRow = append(currRow, wall)
					currRow = append(currRow, wall)
				case box:
					currRow = append(currRow, boxOpen)
					currRow = append(currRow, boxClose)
				case robot:
					currRow = append(currRow, space)
					currRow = append(currRow, space)
					sIdx = index{r: len(grid), c: i * 2}
				case space:
					currRow = append(currRow, space)
					currRow = append(currRow, space)
				}
			}
			grid = append(grid, currRow)
		}
	}

	movesFile, err := os.Open("day-15-moves." + ext)
	if err != nil {
		log.Fatal(err)
	}
	defer movesFile.Close()

	scanner = bufio.NewScanner(movesFile)

	for scanner.Scan() {
		moves = append(moves, []rune(scanner.Text())...)
	}
	return
}

func partOne(grid [][]rune, moves []rune, sIdx index) (sum int) {
	currIdx := sIdx
	for _, move := range moves {
		dir := moveToDirection[move]
		nextIdx := index{r: currIdx.r + dir.r, c: currIdx.c + dir.c}
		if canMoveTo(grid, nextIdx, dir) {
			grid[nextIdx.r][nextIdx.c] = space
			currIdx = nextIdx
		}
	}
	for r := range grid {
		for c := range grid[r] {
			if grid[r][c] == box {
				sum += (100 * r) + c
			}
		}
	}
	return
}

func canMoveTo(grid [][]rune, idx index, dir index) bool {
	if grid[idx.r][idx.c] == space {
		return true
	}
	if grid[idx.r][idx.c] == wall {
		return false
	}

	idx.r += dir.r
	idx.c += dir.c

	if isWithinBounds(idx, len(grid), len(grid[0])) && canMoveTo(grid, idx, dir) {
		grid[idx.r][idx.c] = box
		return true
	} else {
		return false
	}
}

func partTwo(grid [][]rune, moves []rune, sIdx index) (sum int) {
	currIdx := sIdx
	for _, move := range moves {
		dir := moveToDirection[move]
		nextIdx := index{r: currIdx.r + dir.r, c: currIdx.c + dir.c}
		if canMoveToV2(grid, nextIdx, dir) {
			grid[nextIdx.r][nextIdx.c] = robot
			grid[currIdx.r][currIdx.c] = space
			currIdx = nextIdx
		}
	}
	for r := range grid {
		for c := range grid[r] {
			if grid[r][c] == boxOpen {
				sum += (100 * r) + c
			}
		}
	}
	return
}

func canMoveToV2(grid [][]rune, idx index, dir index) bool {
	if grid[idx.r][idx.c] == space {
		return true
	}
	if grid[idx.r][idx.c] == wall {
		return false
	}

	if dirToMove[dir] == '<' {
		nextToOpenIdx := index{r: idx.r + dir.r, c: idx.c-1 + dir.c}
		if canMoveToV2(grid, nextToOpenIdx, dir) {
			grid[nextToOpenIdx.r][nextToOpenIdx.c] = boxOpen
			grid[idx.r][idx.c-1] = boxClose
			return true
		}
	} else if dirToMove[dir] == '>' {
		nextToCloseIdx := index{r: idx.r + dir.r, c: idx.c+1 + dir.c}
		if canMoveToV2(grid, nextToCloseIdx, dir) {
			grid[nextToCloseIdx.r][nextToCloseIdx.c] = boxClose
			grid[idx.r][idx.c+1] = boxOpen
			return true
		}
	} else {
		var other index
		if grid[idx.r][idx.c] == boxOpen {
			other = index{r: idx.r, c: idx.c + 1}
		} else if grid[idx.r][idx.c] == boxClose {
			other = index{r: idx.r, c: idx.c - 1}
		}
		
		nextIdxOther := index{r: other.r + dir.r, c: other.c + dir.c}
		nextIdx := index{r: idx.r + dir.r, c: idx.c + dir.c}
		if isWithinBounds(nextIdxOther, len(grid), len(grid[0])) && canMoveToV2(grid, nextIdxOther, dir) &&
			isWithinBounds(nextIdx, len(grid), len(grid[0])) && canMoveToV2(grid, nextIdx, dir) {
			grid[nextIdx.r][nextIdx.c] = grid[idx.r][idx.c]
			grid[nextIdxOther.r][nextIdxOther.c] = grid[other.r][other.c]
			grid[idx.r][idx.c] = space
			grid[other.r][other.c] = space
			return true
		}
	}
	
	return false
}

func printGrid(grid [][]rune) {
	for r := range grid {
		for c := range grid[r] {
			fmt.Print(string(grid[r][c]))
		}
		fmt.Println()
	}
}

func printMoves(moves []rune) {
	for i := range moves {
		fmt.Print(string(moves[i]))
	}
	fmt.Println()
}

func isWithinBounds(i index, maxRow, maxCol int) bool {
	return i.r >= 0 && i.c >= 0 && i.r < maxRow && i.c < maxCol
}
