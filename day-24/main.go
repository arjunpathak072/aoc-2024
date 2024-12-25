package main

import (
	"bufio"
	"bytes"
	"cmp"
	"fmt"
	"log"
	"maps"
	"os"
	"regexp"
	"slices"
	"strconv"
)

type dependency struct {
	w1, w2 string
	op     string
}

func main() {
	value, dependencies := parseInput("day-24.input")
	fmt.Println("Part One:", partOne(value, dependencies))
	fmt.Println("Part Two:", partTwo(dependencies))
}

func parseInput(name string) (map[string]int8, map[string]dependency) {
	instrRegex := regexp.MustCompile(`([a-z0-9]*) ([A-Z]*) ([a-z0-9]*) -> ([a-z0-9]*)`)
	wireValueRegex := regexp.MustCompile(`([a-zA-Z0-9]*): ([0-9])`)

	file, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)

	value := make(map[string]int8)
	dependencies := make(map[string]dependency)

	for sc.Scan() && sc.Text() != "" {
		matches := wireValueRegex.FindStringSubmatch(sc.Text())
		w := matches[1]
		v := int8(matches[2][0] - '0')
		value[w] = v
	}

	for sc.Scan() {
		matches := instrRegex.FindStringSubmatch(sc.Text())
		w := matches[4]
		op := matches[2]
		w1, w2 := matches[1], matches[3]

		dependencies[w] = dependency{
			w1: w1,
			w2: w2,
			op: op,
		}
	}

	return value, dependencies
}

func partOne(value map[string]int8, dependencies map[string]dependency) (res uint64) {
	var resolve func(string) int8

	resolve = func(curr string) int8 {
		if v, ok := value[curr]; ok {
			return v
		}

		d := dependencies[curr]
		v1 := resolve(d.w1)
		v2 := resolve(d.w2)

		switch d.op {
		case "XOR":
			value[curr] = v1 ^ v2
		case "AND":
			value[curr] = v1 & v2
		case "OR":
			value[curr] = v1 | v2
		}

		return value[curr]
	}

	for n := range dependencies {
		resolve(n)
	}

	for n, v := range value {
		if n[0] == 'z' {
			temp, _ := strconv.Atoi(n[1:])
			res |= uint64(v) << temp
		}
	}

	return
}

func generateGraph(value map[string]int8, dependencies map[string]dependency) {
	file, err := os.Create("./temp/.dot")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var xBuf, yBuf, buf bytes.Buffer

	inputWires := slices.Collect(maps.Keys(value))
	slices.Sort(inputWires)

	for _, w := range inputWires {
		if w[0] == 'x' {
			xBuf.Write([]byte(fmt.Sprintf("%s -> ", w)))
		} else if w[0] == 'y' {
			yBuf.Write([]byte(fmt.Sprintf("%s -> ", w)))
		}
	}

	xBuf.Write([]byte("\n"))
	yBuf.Write([]byte("\n"))

	buf.Write([]byte("digraph G {\n"))
	buf.Write(xBuf.Bytes())
	buf.Write(yBuf.Bytes())

	depList := make([][]string, 0, len(dependencies))
	for w, d := range dependencies {
		fmt.Println(w)
		depList = append(depList, []string{d.w1, w, d.op})
		depList = append(depList, []string{d.w2, w, d.op})
	}

	slices.SortFunc(depList, func(a []string, b []string) int {
		return cmp.Compare(a[1], b[1])
	})

	for _, pair := range depList {
		var color string
		switch pair[2] {
		case "XOR":
			color = "red"
		case "AND":
			color = "blue"
		case "OR":
			color = "green"
		}
		buf.Write([]byte(fmt.Sprintf("%s -> %s [color = %s];\n", pair[0], pair[1], color)))
	}
	buf.Write([]byte("}"))

	file.Write(buf.Bytes())
}

func partTwo(dependencies map[string]dependency) (res string) {
	temp := make(map[string]bool)

	for w, d := range dependencies {
		if w[0] == 'z' {
			val, _ := strconv.Atoi(w[1:])
			if d.op != "XOR" && val != 45 {
				temp[w] = true
			}
		} else if !isXOrY(d.w1) && !isXOrY(d.w2) && d.w1[0] != d.w2[0] && d.op == "XOR" {
			temp[w] = true
		}

		if d.op == "XOR" && isXOrY(d.w1) && isXOrY(d.w2) && d.w1[0] != d.w2[0] {
			isValid := false
			for _, dp := range dependencies {
				if dp.op == "XOR" && (dp.w1 == w || dp.w2 == w) {
					isValid = true
				}
			}
			if !isValid {
				temp[w] = true
			}
		}

		if d.op == "AND" && isXOrY(d.w1) && isXOrY(d.w2) && d.w1[0] != d.w2[0] {
			isValid := false
			for _, dp := range dependencies {
				if dp.op == "OR" && (dp.w1 == w || dp.w2 == w) {
					isValid = true
				}
			}
			if !isValid {
				temp[w] = true
			}
		}
	}
	ans := slices.Collect(maps.Keys(temp))
	slices.Sort(ans)

	for _, w := range ans {
		res += w + ","
	}

	return res[:len(res)-1]
}

func isXOrY(wire string) bool {
	temp, _ := strconv.Atoi(wire[1:])
	return (wire[0] == 'x' || wire[0] == 'y') && temp != 0
}
