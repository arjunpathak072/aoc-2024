package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"os"
)

type freeSpace struct {
	start int
	end   int
}

type Item struct {
	fs       freeSpace
	priority int
	index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) update(item *Item, fs freeSpace, priority int) {
	item.fs = fs
	item.priority = priority
	heap.Fix(pq, item.index)
}

func main() {
	disk := parseInput("day-9.example")
	fmt.Println("Part One:", partOne(disk))
}

func parseInput(fileName string) (disk []int) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		disk = make([]int, 0, len(scanner.Text()))
		id := 0
		for i, r := range scanner.Text() {
			times := int(r - '0')
			if i&1 == 1 { // is a vacancy
				for j := 0; j < times; j++ {
					disk = append(disk, -1)
				}
			} else { // is a file
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

func partTwo(disk []int) (sum int) {
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

func isVacant(i int) bool {
	return i&1 == 1
}
