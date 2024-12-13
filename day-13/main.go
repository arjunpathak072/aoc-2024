package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	machines := parseInput("day-13.input", false)
	fmt.Println("Part One: ", solve(machines, false))
	
	machines = parseInput("day-13.input", true)
	fmt.Println("Part Two: ", solve(machines, true))
}

type equation struct {
	xc, yc, rhs int64
}

type pair struct {
	first, second equation
}

func parseInput(fileName string, partTwo bool) (machines []pair) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	r := regexp.MustCompile(`Button A: X\+([0-9]*), Y\+([0-9]*)\nButton B: X\+([0-9]*), Y\+([0-9]*)\nPrize: X=([0-9]*), Y=([0-9]*)`)

	content, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	matches := r.FindAllStringSubmatch(string(content), -1)

	for _, match := range matches {
		var p pair

		p.first.xc, _ = strconv.ParseInt(match[1], 10, 64)
		p.first.yc, _ = strconv.ParseInt(match[3], 10, 64)
		p.first.rhs, _ = strconv.ParseInt(match[5], 10, 64)

		p.second.xc, _ = strconv.ParseInt(match[2], 10, 64)
		p.second.yc, _ = strconv.ParseInt(match[4], 10, 64)
		p.second.rhs, _ = strconv.ParseInt(match[6], 10, 64)

		if partTwo {
			p.first.rhs += 10000000000000
			p.second.rhs += 10000000000000
		}

		machines = append(machines, p)
	}

	return machines
}

func solve(machines []pair, partTwo bool) (cost int64) {
	for _, p := range machines {
		cost += calculateCost(p, partTwo)
	}
	return cost
}

func calculateCost(p pair, partTwo bool) int64 {
	firstY, secondY := p.first.yc, p.second.yc

	p.second.xc *= firstY
	p.second.yc *= firstY
	p.second.rhs *= firstY

	p.first.xc *= secondY
	p.first.yc *= secondY
	p.first.rhs *= secondY

	xDelta := abs(p.first.xc - p.second.xc)
	rhsDelta := abs(p.first.rhs - p.second.rhs)

	if !(rhsDelta%xDelta == 0) {
		return 0
	}
	aPresses := rhsDelta / xDelta
	if !((p.first.rhs-aPresses*p.first.xc)%p.first.yc == 0) {
		return 0
	}
	bPresses := (p.first.rhs - aPresses*p.first.xc) / p.first.yc
	if partTwo {
		return aPresses*3 + bPresses
	} else if aPresses <= 100 && bPresses <= 100 {
		return aPresses*3 + bPresses
	}
	
	return 0
}

func abs(x int64) int64 {
	if x < 0 {
		return -1 * x
	}
	return x
}
