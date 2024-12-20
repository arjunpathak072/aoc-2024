package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const (
	start = 'S'
	end   = 'E'
	track = '.'
	wall  = '#'
)

type index struct {
	r, c int
}

type direction struct {
	dr, dc int
}

var directions = []direction{
	{-1, 0},
	{0, -1},
	{1, 0},
	{0, 1},
}

func main() {
	grid, sIdx, eIdx := parseInput("day-20.example")
	fmt.Println("Part One:", partOne(grid, sIdx, eIdx))
	fmt.Println("Part Two:", partTwo(grid, sIdx, eIdx))
}

func parseInput(name string) (grid [][]rune, sIdx, eIdx index) {
	file, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		grid = append(grid, make([]rune, 0, len(scanner.Text())))
		r := len(grid) - 1
		for c, block := range scanner.Text() {
			if block == start {
				sIdx = index{r: r, c: c}
			}
			if block == end {
				eIdx = index{r: r, c: c}
			}
			grid[r] = append(grid[r], block)
		}
	}
	return
}

func partOne(grid [][]rune, sIdx, eIdx index) (count int) {
	distanceFromStart := BFS(grid, sIdx)
	distanceFromEnd := BFS(grid, eIdx)

	maxRow, maxCol := len(grid), len(grid[0])
	minDistanceWithoutCheating := distanceFromStart[eIdx]

	for idx, ds := range distanceFromStart {
		for _, cIdx := range getCheatIndices(idx, maxRow, maxCol) {
			if de, ok := distanceFromEnd[cIdx]; ok {
				dist := ds + de + getManhattenDistance(idx, cIdx)
				if minDistanceWithoutCheating-dist >= 100 {
					count++
				}
			}
		}
	}
	return
}

func partTwo(grid [][]rune, sIdx, eIdx index) (count int) {
	distanceFromStart := BFS(grid, sIdx)
	distanceFromEnd := BFS(grid, eIdx)

	maxRow, maxCol := len(grid), len(grid[0])
	minDistanceWithoutCheating := distanceFromStart[eIdx]

	for r := range grid {
		for c := range grid[r] {
			idx := index{r: r, c: c}
			if ds, ok := distanceFromStart[idx]; ok {
				for r := max(0, idx.r-20); r <= min(idx.r+20, maxRow-1); r++ {
					for c := max(0, idx.c-20); c <= min(idx.c+20, maxCol-1); c++ {
						cIdx := index{r: r, c: c}
						md := getManhattenDistance(idx, cIdx)
						if de, ok := distanceFromEnd[cIdx]; ok && md <= 20 {
							dist := ds + de + getManhattenDistance(idx, cIdx)
							if minDistanceWithoutCheating-dist == 50 {
								count++
							}
						}
					}
				}
			}
		}
	}
	return
}

func getCheatIndices(idx index, maxRow, maxCol int) (cheatIndices []index) {
	for i := 1; i <= 2; i++ {
		for _, dir := range directions {
			nIdx := index{r: idx.r + i*dir.dr, c: idx.c + i*dir.dc}
			if isWithinBounds(nIdx, maxRow, maxCol) {
				cheatIndices = append(cheatIndices, nIdx)
			}
		}
	}
	return
}

func getManhattenDistance(i1 index, i2 index) int {
	return abs(i1.r-i2.r) + abs(i1.c-i2.c)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func BFS(grid [][]rune, sIdx index) map[index]int {
	type qNode struct {
		idx      index
		distance int
	}
	maxRow, maxCol := len(grid), len(grid[0])

	distMap := make(map[index]int)
	queue := []qNode{{idx: sIdx, distance: 0}}
	distMap[sIdx] = 0

	for len(queue) > 0 {
		front := queue[0]
		queue = queue[1:]

		distMap[front.idx] = front.distance

		for _, dir := range directions {
			nIdx := index{r: front.idx.r + dir.dr, c: front.idx.c + dir.dc}
			if _, ok := distMap[nIdx]; !ok && isWithinBounds(nIdx, maxRow, maxCol) && grid[nIdx.r][nIdx.c] != wall {
				queue = append(queue, qNode{idx: nIdx, distance: front.distance + 1})
			}
		}
	}
	return distMap
}

func isWithinBounds(idx index, maxRow, maxCol int) bool {
	return idx.r >= 0 && idx.c >= 0 && idx.r < maxRow && idx.c < maxCol
}
