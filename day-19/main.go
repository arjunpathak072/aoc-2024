package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	towels, designs := parseInput("day-19.input")
	p1, p2 := solve(towels, designs)
	fmt.Println("Part One:", p1)
	fmt.Println("Part Two:", p2)
}

func parseInput(name string) (towels map[string]bool, designs []string) {
	file, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	towels = make(map[string]bool)
	scanner.Scan()
	for _, towel := range strings.Split(scanner.Text(), ", ") {
		towels[towel] = true
	}
	scanner.Scan()
	for scanner.Scan() {
		designs = append(designs, scanner.Text())
	}
	return
}

func solve(towels map[string]bool, designs []string) (p1, p2 int) {
	for _, design := range designs {
		countFormations := make([]int, len(design)+1)
		countFormations[0] = 1

		for i := 1; i <= len(design); i++ {
			for j := 0; j < i; j++ {
				if towels[design[j:i]] {
					countFormations[i] += countFormations[j]
				}
			}
		}
		
		ways := countFormations[len(design)]
		if ways != 0 {
			p1++
		}
		p2 += ways
	}
	return
}
