package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	puzzleInput := parseInput("day-3.input")
	fmt.Println("Part One: ", partOne(puzzleInput))
	fmt.Println("Part Two: ", partTwo(puzzleInput))
}

func parseInput(fileName string) (puzzleInput []string) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		puzzleInput = append(puzzleInput, scanner.Text())
	}

	return
}

func partOne(input []string) (sum int) {
	r := regexp.MustCompile(`mul\(\d+,\d+\)`)
	for _, line := range input {
		for _, instruction := range r.FindAllString(line, -1) {
			sum += evaludateMul(instruction)
		}
	}
	return
}

func partTwo(input []string) (sum int) {
	r := regexp.MustCompile(`mul\(\d+,\d+\)|do\(\)|don't\(\)`)
	validMul := true
	for _, line := range input {
		for _, instruction := range r.FindAllString(line, -1) {
			if strings.Contains(instruction, "don't") {
				validMul = false
			} else if strings.Contains(instruction, "do") {
				validMul = true
			} else if (validMul) {
				sum += evaludateMul(instruction)
			}
		}
	}
	return sum
}

func evaludateMul(instruction string) int {
	r := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	matches := r.FindStringSubmatch(instruction)
	
	num1, _ := strconv.Atoi(matches[1])
	num2, _ := strconv.Atoi(matches[2])

	product := num1 * num2
	
	fmt.Println(instruction, "->", product)
	return num1 * num2
}

func convertToInt(str string) (num int) {
	for _, r := range str {
		num *= 10
		num += int(r - '0')
	}
	return
}
