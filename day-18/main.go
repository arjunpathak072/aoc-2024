package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type position struct {
	x, y int
}

var directions = []position{
	{-1, 0},
	{0, -1},
	{1, 0},
	{0, 1},
}

var max = 71

func main() {
	corrupted, corruptedPositions := parseInput("day-18.input")
	fmt.Println("Part One:", partOne(corrupted))
	fmt.Println("Part Two:", partTwo(corruptedPositions, corrupted))
}

func parseInput(name string) (map[position]bool, []position) {
	file, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var corruptedPositions []position
	corrupted := make(map[position]bool)
	scanner := bufio.NewScanner(file)
	
	for i := 0; scanner.Scan(); i++ {
		temp := strings.Split(scanner.Text(), ",")
		x, _ := strconv.Atoi(temp[0])
		y, _ := strconv.Atoi(temp[1])
		pos := position{x, y}
		if i < 1024 {
			corrupted[pos] = true
		}
		corruptedPositions = append(corruptedPositions, pos)
	}
	return corrupted, corruptedPositions
}

func partOne(corrupted map[position]bool) int {
	type qNode struct {
		pos   position
		steps int
	}
	sPos, ePos := position{0, 0}, position{max-1, max-1}
	queue := []qNode{{pos: sPos}}
	visited := make(map[position]bool)
	visited[sPos] = true

	for len(queue) > 0 {
		front := queue[0]
		queue = queue[1:]
		
		if front.pos == ePos {
			return front.steps
		}
		
		for _, dir := range directions {
			nPos := position{x: front.pos.x + dir.x, y: front.pos.y + dir.y}
			if !isWithinBounds(nPos) || visited[nPos] || corrupted[nPos] {
				continue
			}
			node := qNode{pos: nPos, steps: front.steps + 1}
			visited[nPos] = true
			queue = append(queue, node)
		}
	}
	return -1
}

func partTwo(corruptedPositions []position, corrupted map[position]bool) position {
	for i := 1024; i < len(corruptedPositions); i++ {
		pos := corruptedPositions[i]
		corrupted[position{pos.x, pos.y}] = true
		if val := partOne(corrupted); val == -1 {
			return pos
		}
	}
	panic("no such byte exists")
}

func isWithinBounds(pos position) bool {
	return pos.x >= 0 && pos.y >= 0 && pos.x < max && pos.y < max
}