package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

type adjList map[string][]string

func main() {
	al := parseInput("day-23.input")
	fmt.Println("Part One:", partOne(al))
	fmt.Println("Part Two:", partTwo(al))
}

func parseInput(name string) adjList {
	file, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	al := make(adjList)

	for scanner.Scan() {
		temp := strings.Split(scanner.Text(), "-")
		a, b := temp[0], temp[1]
		al[a] = append(al[a], b)
		al[b] = append(al[b], a)
	}

	return al
}

func partOne(al adjList) int {
	cycles := make(map[string][]string)
	for n1 := range al {
		if n1[0] != 't' {
			continue
		}
		for _, n2 := range al[n1] {
			for _, n3 := range al[n2] {
				if slices.Contains(al[n3], n1) {
					cycle := []string{n1, n2, n3}
					slices.Sort(cycle)
					key := fmt.Sprintf("%s%s%s", cycle[0], cycle[1], cycle[2])
					cycles[key] = cycle
				}
			}
		}
	}
	return len(cycles)
}

func partTwo(al adjList) (password string) {
	nodesList := make([]string, 0, len(al))
	
	for node := range al {
		nodesList = append(nodesList, node)
	}
	
	var maxLenClique []string
	var maxLen int
	
	BronKerbosch([]string{}, nodesList, []string{}, &maxLenClique, &maxLen, al)
	
	slices.Sort(maxLenClique)
	for _, s := range maxLenClique {
		password += s + ","
	}
	
	return password[:len(password)-1]
}

func BronKerbosch(R, P, X []string, maxLenClique *[]string, maxLen *int, al adjList) {
	if len(P) == 0 && len(X) == 0 && *maxLen < len(R) {
		rCopy := slices.Clone(R)
		*maxLenClique = rCopy
		*maxLen = len(rCopy)
		return
	}
	
	pCopy := slices.Clone(P)
	for _, v := range pCopy {
		newR := append(R, v)
		neighbours := al[v]
		
		newP := intersect(P, neighbours)
		newX := intersect(X, neighbours)
		
		BronKerbosch(newR, newP, newX, maxLenClique, maxLen, al)
		
		vIdx := slices.Index(P, v)
		P = slices.Delete(P, vIdx, vIdx+1)
		
		X = append(X, v)
	}
}

func intersect(a, b []string) (res []string) {
	t := make(map[string]bool)
	for _, val := range a {
		t[val] = true
	}
	for _, val := range b {
		if t[val] {
			res = append(res, val)
		}
	}
	return
}