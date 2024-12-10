package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type hole struct {
	start int
	size  int
}

type file struct {
	id    int
	start int
	size  int
}

func main() {
	disk, _, _ := parseInput("day-9.example")
	fmt.Println("Part One:", partOne(disk))
	disk, holes, files := parseInput("day-9.input")
	fmt.Println("Part Two:", partTwo(disk, holes, files))
}

func parseInput(fileName string) (disk []int, holes []hole, files []file) {
	inputFile, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	for scanner.Scan() {
		disk = make([]int, 0, len(scanner.Text()))
		id := 0
		for i, r := range scanner.Text() {
			times := int(r - '0')
			if i&1 == 1 { // is a vacancy
				holes = append(holes, hole{start: len(disk), size: times})
				for j := 0; j < times; j++ {
					disk = append(disk, -1)
				}
			} else { // is a file
				files = append(files, file{id: id, start: len(disk), size: times})
				for j := 0; j < times; j++ {
					disk = append(disk, id)
				}
				id++
			}
		}
	}
	return
}

func partOne(disk []int) (sum int) {
	for l, r := 0, len(disk)-1; l <= r; {
		for disk[l] != -1 {
			l++
		}
		for disk[r] == -1 {
			r--
		}
		if l < r {
			disk[l], disk[r] = disk[r], disk[l]
		}
	}

	for i := 0; disk[i] != -1; i++ {
		sum += disk[i] * i
	}
	return
}

func partTwo(disk []int, holes []hole, files []file) (sum int) {
	for i := len(files) - 1; i >= 0; i-- {
		currFile := files[i]
		for j := range holes {
			if holes[j].size >= currFile.size && holes[j].start < currFile.start {
				for k := 0; k < currFile.size; k++ {
					disk[k + holes[j].start] = currFile.id
					disk[k + currFile.start] = -1
				}
				holes[j].size -= currFile.size
				holes[j].start += currFile.size
				break
			}
		}
	}
	for i := range disk {
		if disk[i] != -1 {
			sum += disk[i] * i
		}
	}
	return
}

func isVacant(i int) bool {
	return i&1 == 1
}
