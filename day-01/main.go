package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	fmt.Println("Part One: ", partOne())
	fmt.Println("Part Two: ", partTwo())
}

func partOne() int {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	
	scanner := bufio.NewScanner(file)
	
	list1, list2 := make([]int, 0, 1000), make([]int, 0, 1000)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		list1 = append(list1, convertToInt(fields[0]))
		list2 = append(list2, convertToInt(fields[1]))
	}
	
	slices.Sort(list1)
	slices.Sort(list2)
	
	sumOfDiffs := 0
	for i := 0; i < 1000; i++ {
		sumOfDiffs += getAbsoluteDifference(list1[i], list2[i])
	}
	
	return sumOfDiffs
}

func partTwo() int {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	
	scanner := bufio.NewScanner(file)
	
	list := make([]int, 0, 1000)
	frequencyMap := make(map[int]int)
	
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		list = append(list, convertToInt(fields[0]))
		frequencyMap[convertToInt(fields[1])]++
	}
	
	sumOfDiffs := 0
	for i := 0; i < 1000; i++ {
		sumOfDiffs += list[i] * frequencyMap[list[i]]
	}
	
	return sumOfDiffs
}


func getAbsoluteDifference(num1, num2 int) int {
	diff := num1 - num2
	if diff < 0 {
		return -1 * diff
	}
	return diff
}

func convertToInt(str string) (num int) {
	for _, r := range str {
		num *= 10
		num += int(r - '0')
	}
	return
}