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

const (
	trailhead = 0
	trailend  = 9
)

var directions = []index{
	{-1, 0},
	{0, -1},
	{1, 0},
	{0, 1},
}

func main() {
	puzzleInput := parseInput("day-10.input")
	fmt.Println("Part One: ", partOne(puzzleInput))
	fmt.Println("Part Two: ", partTwo(puzzleInput))
}

func parseInput(fileName string) (puzzleInput [][]int) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		temp := make([]int, 0, len(scanner.Text()))
		for _, r := range scanner.Text() {
			temp = append(temp, int(r-'0'))
		}
		puzzleInput = append(puzzleInput, temp)
	}
	return
}

func partOne(topoMap [][]int) (sum int) {
	for r := range topoMap {
		for c := range topoMap {
			if topoMap[r][c] == trailhead {
				visited := make(map[index]bool)
				calculateScore(topoMap, index{r: r, c: c}, &sum, visited)
			}
		}
	}
	return
}

func calculateScore(topoMap [][]int, curr index, sum *int, visited map[index]bool) {
	currValue := topoMap[curr.r][curr.c]
	if currValue == trailend {
		if !visited[curr] {
			(*sum)++
			visited[curr] = true
		}
		return
	}

	maxRow, maxCol := len(topoMap), len(topoMap[0])
	for _, dir := range directions {
		nextIndex := index{r: curr.r + dir.r, c: curr.c + dir.c}
		if isWithinBounds(nextIndex, maxRow, maxCol) &&
			topoMap[nextIndex.r][nextIndex.c] == currValue+1 {
			calculateScore(topoMap, nextIndex, sum, visited)
		}
	}
}

func partTwo(topoMap [][]int) (sum int) {
	for r := range topoMap {
		for c := range topoMap {
			if topoMap[r][c] == trailhead {
				calculateRating(topoMap, index{r: r, c: c}, &sum)
			}
		}
	}
	return
}

func calculateRating(topoMap [][]int, curr index, sum *int) {
	currValue := topoMap[curr.r][curr.c]

	if currValue == trailend {
		(*sum)++
		return
	}

	maxRow, maxCol := len(topoMap), len(topoMap[0])
	for _, dir := range directions {
		nextIndex := index{r: curr.r + dir.r, c: curr.c + dir.c}
		if isWithinBounds(nextIndex, maxRow, maxCol) &&
			topoMap[nextIndex.r][nextIndex.c] == currValue+1 {
			calculateRating(topoMap, nextIndex, sum)
		}
	}
}

func isWithinBounds(i index, maxRow, maxCol int) bool {
	return i.r >= 0 && i.c >= 0 && i.r < maxRow && i.c < maxCol
}
