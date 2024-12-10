package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

func main() {
	comesBefore, updates := parseInput("day-5.input")
	fmt.Println("Part One: ", partOne(comesBefore, updates))
	comesBefore, updates = parseInput("day-5.input")
	fmt.Println("Part Two: ", partTwo(comesBefore, updates))
}

func parseInput(fileName string) (map[int][]int, [][]int) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	comesBefore := make(map[int][]int)
	for scanner.Scan() && scanner.Text() != "" {
		fields := strings.Split(scanner.Text(), "|")
		num1, num2 := convertToInteger(fields[0]), convertToInteger(fields[1])
		comesBefore[num1] = append(comesBefore[num1], num2)
	}

	var updates [][]int
	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), ",")
		update := make([]int, 0, len(fields))
		for _, f := range fields {
			update = append(update, convertToInteger(f))
		}
		updates = append(updates, update)
	}

	return comesBefore, updates
}

func partOne(comesBefore map[int][]int, updates [][]int) (sum int) {
	for _, update := range updates {
		original := make([]int, len(update))
		copy(original, update)

		slices.SortFunc(update, func(a, b int) int {
			if slices.Contains(comesBefore[a], b) {
				return -1
			} else {
				return 1
			}
		})
		
		if slices.Equal(update, original) {
			sum += original[len(original)/2]
		}
	}
	return sum
}

func partTwo(comesBefore map[int][]int, updates [][]int) (sum int) {
	for _, update := range updates {
		original := make([]int, len(update))
		copy(original, update)

		slices.SortFunc(update, func(a, b int) int {
			if slices.Contains(comesBefore[a], b) {
				return -1
			} else {
				return 1
			}
		})
		
		if !slices.Equal(update, original) {
			sum += update[len(update)/2]
		}
	}
	return sum
}

func convertToInteger(str string) (num int) {
	for _, r := range str {
		num *= 10
		num += int(r - '0')
	}
	return
}
