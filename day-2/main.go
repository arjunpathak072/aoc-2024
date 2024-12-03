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
	puzzleInput := parseFile("day-2.input")
	fmt.Println("PartOne: ", partOne(puzzleInput))
	fmt.Println("PartTwo: ", partTwo(puzzleInput))
}

func parseFile(fileName string) [][]int {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	reports := make([][]int, 0, 1000)

	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		var integerFields []int
		for _, s := range fields {
			integerFields = append(integerFields, stringToInt(s))
		}
		reports = append(reports, integerFields)
	}

	return reports
}

func partOne(reports [][]int) (safeReports int) {
	for _, report := range reports {
		if isValid(report) {
			safeReports++
		}
	}
	return safeReports
}

func partTwo(reports [][]int) (safeReports int) {
	for _, report := range reports {
		if isValid(report) {
			safeReports++
		} else if canBeMadeValid(report) {
			safeReports++
		}
	}
	return safeReports
}

func isValid(report []int) bool {
	isSafe := true

	for i := 1; i < len(report); i++ {
		diff := abs(report[i] - report[i-1])

		if diff > 3 || diff < 1 {
			isSafe = false
		}
	}

	if !isStrictlyIncreasing(report) && !isStrictlyDecreasing(report) {
		isSafe = false
	}

	return isSafe
}

func isStrictlyIncreasing(report []int) bool {
	for i := 1; i < len(report); i++ {
		if report[i-1] >= report[i] {
			return false
		}
	}
	return true
}

func isStrictlyDecreasing(report []int) bool {
	for i := 1; i < len(report); i++ {
		if report[i] >= report[i-1] {
			return false
		}
	}
	return true
}

func canBeMadeValid(report []int) bool {
	for i := range report {
		newSlice := make([]int, len(report))
		copy(newSlice, report)

		slices.Delete(newSlice, i, i+1)
		newSlice = newSlice[:len(newSlice)-1]

		if isValid(newSlice) {
			return true
		}
	}
	return false
}

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}

func stringToInt(str string) (num int) {
	for _, r := range str {
		num *= 10
		num += int(r - '0')
	}

	return num
}
