package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

const (
	wall = '#'
	path = '.'
)

type index struct {
	x, y int
}

type direction struct {
	dx, dy int
}

type node struct {
	idx index
	dir direction
}

type state struct {
	reindeer node
	path     []index
	score    int
}

var directions = []direction{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
}

func main() {
	grid, sIdx, eIdx := parseInput("day-16.input")
	p1, p2 := BFS(grid, sIdx, eIdx)
	fmt.Println("Part One:", p1)
	fmt.Println("Part Two:", p2)
}

func parseInput(fileName string) (grid [][]rune, sIdx, eIdx index) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		for i, r := range []rune(scanner.Text()) {
			if r == 'S' {
				sIdx = index{x: len(grid), y: i}
			} else if r == 'E' {
				eIdx = index{x: len(grid), y: i}
			}
		}
		grid = append(grid, []rune(scanner.Text()))
	}
	return
}

func BFS(grid [][]rune, sIdx, eIdx index) (minScore int, bestSeatCount int) {
	minScore = math.MaxInt
	reindeer := node{idx: sIdx, dir: direction{dx: 0, dy: 1}}
	queue := []state{
		state{
			reindeer: reindeer,
			path: []index{sIdx},
			score: 0,
		},
	}
	visited := make(map[node]int)
	sizeToIndices := make(map[int][]index)

	for len(queue) > 0 {
		currState := queue[0]
		queue = queue[1:]

		if currState.score > minScore {
			continue
		}

		if currState.reindeer.idx == eIdx {
			if currState.score <= minScore {
				minScore = currState.score
				sizeToIndices[minScore] = append(sizeToIndices[minScore], currState.path...)
			}
			continue
		}

		for _, n := range getNeighbours(currState.reindeer) {
			if grid[n.idx.x][n.idx.y] == wall {
				continue
			}
			score := currState.score + 1
			if currState.reindeer.dir != n.dir {
				score += 1000
			}
			if previous, has := visited[n]; has {
				if previous < score {
					continue
				}
			}
			visited[n] = score
			
			nPath := make([]index, len(currState.path))
			copy(nPath, currState.path)
			
			queue = append(queue, state{
				reindeer: n,
				path: append(nPath, n.idx),
				score: score,
			})
		}
	}
	
	countMap := make(map[index]bool)
	for _, index := range sizeToIndices[minScore] {
		countMap[index] = true
	}

	return minScore, len(countMap)
}

func getNeighbours(reindeer node) (neighbours []node) {
	neighbours = make([]node, 0, 4)
	currDir, currIdx := reindeer.dir, reindeer.idx
	oppositeDir := direction{dx: -currDir.dx, dy: -currDir.dy}

	for _, dir := range directions {
		if dir == oppositeDir {
			continue
		}
		nIdx := index{x: currIdx.x + dir.dx, y: currIdx.y + dir.dy}
		neighbours = append(neighbours, node{idx: nIdx, dir: dir})
	}

	return
}