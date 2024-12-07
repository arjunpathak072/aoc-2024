package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	puzzleInput := parseInput("day-7.example")
	fmt.Println("Part One: ", partOne(puzzleInput))
}

func partOne(puzzleInput map[uint64][]uint64) (total uint64) {
	for result, operands := range puzzleInput {
		if evaluate(0, operands, 0, result) {
			total += result
		}
	}
	return total
}

func concatenateValues(a, b uint64) uint64 {
	first := strconv.Itoa(int(a))
	second := strconv.Itoa(int(b))

	value, _ := strconv.ParseUint(first+second, 10, 64)
	return value
}

func evaluate(at int, operands []uint64, value, result uint64) bool {
	if at == len(operands) || value > result {
		return value == result
	}
	return evaluate(at+1, operands, value+operands[at], result) || evaluate(at+1, operands, value*operands[at], result)
}

func parseInput(fileName string) (puzzleInput map[uint64][]uint64) {
	puzzleInput = make(map[uint64][]uint64)
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ":")
		result, _ := strconv.ParseUint(line[0], 10, 64)

		fields := strings.Fields(line[1])
		puzzleInput[result] = make([]uint64, 0, len(fields))

		for _, str := range fields {
			value, _ := strconv.ParseUint(str, 10, 64)
			puzzleInput[result] = append(puzzleInput[result], value)
		}
	}

	return puzzleInput
}
