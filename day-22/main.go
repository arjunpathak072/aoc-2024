package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	puzzleInput := parseInput("day-22.input")
	fmt.Println("Part One:", partOne(puzzleInput))
	fmt.Println("Part Two:", partTwo(puzzleInput))
}

func parseInput(name string) (nums []int) {
	file, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		temp, _ := strconv.Atoi(scanner.Text())
		nums = append(nums, temp)
	}
	return
}

func partOne(nums []int) (res int) {
	for _, num := range nums {
		temp := getNthSecretNumber(num, 2000)
		res += temp
	}
	return
}

func getNthSecretNumber(num int, n int) int {
	for ; n > 0; n-- {
		num = ((num << 6) ^ num) & (16777216 - 1)
		num = ((num >> 5) ^ num) & (16777216 - 1)
		num = ((num << 11) ^ num) & (16777216 - 1)
	}
	return num
}

type pair struct {
	delta   int
	bananas int
}

type window struct {
	a, b, c, d int
}

func updateCache(num int, n int, cache *map[window]int) (delBan []pair) {
	visited := make(map[window]bool)
	delBan = make([]pair, 0, n-1)
	currWindow := make([]int, 0, 8)
	prev := num % 10
	
	for ; n > 1; n-- {
		num = ((num << 6) ^ num) & (16777216 - 1)
		num = ((num >> 5) ^ num) & (16777216 - 1)
		num = ((num << 11) ^ num) & (16777216 - 1)

		delta, bananas := num%10-prev, num%10
		delBan = append(delBan, pair{delta: delta, bananas: bananas})
		currWindow = append(currWindow, delta)

		if len(currWindow) == 4 {
			key := window{a: currWindow[0], b: currWindow[1], c: currWindow[2], d: currWindow[3]}
			if !visited[key] {
				(*cache)[key] += num % 10
				visited[key] = true
			}
			currWindow = currWindow[1:]
		}
		prev = num % 10
	}
	return
}

func partTwo(nums []int) (res int) {
	cache := make(map[window]int)
	for _, num := range nums {
		updateCache(num, 2000, &cache)
	}

	for _, value := range cache {
		res = max(res, value)
	}

	return res
}
