package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
)

type robot struct {
	x, y   int
	vx, vy int
}

func main() {
	puzzleInput := parseInput("day-14.input")
	res, _ := partOne(puzzleInput, 100)
	fmt.Println("Part One:", res)
	partTwo()
}

func parseInput(fileName string) (robots []robot) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	r := regexp.MustCompile(`p=([0-9]*),([0-9]*) v=(-?[0-9]*),(-?[0-9]*)`)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		matches := r.FindAllStringSubmatch(scanner.Text(), -1)
		x, _ := strconv.Atoi(matches[0][1])
		y, _ := strconv.Atoi(matches[0][2])
		vx, _ := strconv.Atoi(matches[0][3])
		vy, _ := strconv.Atoi(matches[0][4])
		robots = append(robots, robot{x: x, y: y, vx: vx, vy: vy})
	}
	return
}

func partOne(robots []robot, times int) (int, [][]bool) {
	maxRow, maxCol := 103, 101
	halfRow := maxRow / 2
	halfCol := maxCol / 2

	for i, r := range robots {
		robots[i].x = (r.x + r.vx*times + maxCol*times) % maxCol
		robots[i].y = (r.y + r.vy*times + maxRow*times) % maxRow
	}

	grid := make([][]bool, maxRow)
	for r := range grid {
		grid[r] = make([]bool, maxCol)
	}
	for _, r := range robots {
		grid[r.y][r.x] = true
	}

	quardrants := make([]int, 4)
	for _, r := range robots {
		if r.y < halfRow && r.x < halfCol {
			quardrants[0]++
		}
		if r.y > halfRow && r.x < halfCol {
			quardrants[1]++
		}
		if r.y < halfRow && r.x > halfCol {
			quardrants[2]++
		}
		if r.y > halfRow && r.x > halfCol {
			quardrants[3]++
		}
	}

	safetyScore := 1
	for _, q := range quardrants {
		safetyScore *= q
	}
	return safetyScore, grid
}

func partTwo() {
	for i := 0; i < 101*103; i++ {
		puzzleInput := parseInput("day-14.input")
		_, grid := partOne(puzzleInput, i)
		img := image.NewGray(image.Rect(0, 0, 4*len(grid[0]), 4*len(grid)))

		for y := 0; y < len(grid); y++ {
			for x := 0; x < len(grid[0]); x++ {
				var cellColor color.Gray
				if grid[y][x] {
					cellColor = color.Gray{255}
				} else {
					cellColor = color.Gray{0}
				}

				for dy := 0; dy < 4; dy++ {
					for dx := 0; dx < 4; dx++ {
						img.SetGray(x*4+dx, y*4+dy, cellColor)
					}
				}
			}
		}

		file, err := os.Create(filepath.Join("temp", fmt.Sprintf("%d", i)))
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		err = png.Encode(file, img)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}
