package main

import(
	"log"
	"os"
	"strconv"
	"strings"
	"bufio"
	"fmt"
	"math"
)

func main() {
	puzzleInput := parseInput("day-11.input")
	fmt.Println("Part One: ", solve(puzzleInput, 25))
	fmt.Println("Part Two: ", solve(puzzleInput, 75))
}

func parseInput(fileName string) map[int]int {
	stoneToCount := make(map[int]int)
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	
	scanner := bufio.NewScanner(file)
	
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		for _, s := range fields {
			temp, _ := strconv.Atoi(s)
			stoneToCount[temp]++
		}
	}
	return stoneToCount
}

func solve(stoneToCount map[int]int, blinkCount int) (sum uint64){
	for i := 0; i < blinkCount; i++ {
		updated := make(map[int]int)
		for stone, count := range stoneToCount {
			nDigits := int(math.Floor(math.Log10(float64(stone))) + 1)
			if stone == 0 {
				updated[1] += count
			} else if nDigits&1 == 0 {
				firstHalf := stone / power(10, nDigits/2)
				secondHalf := stone % power(10, nDigits/2)
				updated[firstHalf] += count
				updated[secondHalf] += count
			} else {
				updated[stone*2024] += count
			}
		}
		stoneToCount = updated
	}
	
	for _, count := range stoneToCount {
		sum += uint64(count)
	}
	
	return
}

func power(base, exp int) int {
	result := 1
	for exp > 0 {
		if exp%2 == 1 {
			result *= base
		}
		base *= base
		exp /= 2
	}
	return result
}
