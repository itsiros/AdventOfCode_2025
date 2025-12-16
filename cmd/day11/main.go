package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

var data = make(map[string][]string)
var visited = make(map[string]int)
var mem [][]string
var memorazation []string

const YOU = "you"
const OUT = "out"
const SVR = "svr"
const DAC = "dac"
const FFT = "fft"

func idea(node string) {

	if node == SVR {
		return
	}

	
	for q, d := range data {
		if slices.Contains(d, node) {
			memorazation = append(memorazation, q)
		}
	}

	// for _, m := range memorazation {
	// 	idea(m)
	// }
}

func findPathsPart2(node string, path []string) {
	if slices.Contains(path, node) {
		return
	}
	path = append(path, node)

	if node == OUT {
		mem = append(mem, path)
		return
	}

	for _, next := range data[node] {

		if visited[node] < len(data[node]) {
			visited[node]++
			findPathsPart2(next, path)
		}
	}

	fmt.Println(path)
}

func findPathsPart1(path []string) int {

	count := 0
	if slices.Contains(path, OUT) {
		return 1
	}

	for _, p := range path {
		count += findPathsPart1(data[p])
	}
	return count
}

func main() {

	if len(os.Args) != 2 {
		return
	}

	file, fErr := os.Open(os.Args[1])
	if fErr != nil {
		fmt.Fprintln(os.Stderr, fErr)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) < 2 {
			continue
		}

		key := strings.Trim(fields[0], ":")
		data[key] = append(data[key], fields[1:]...)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, fErr)
		return
	}

	// for key, _ := range data {
	// 	fmt.Println(key, data[key])
	// }

	// part1 := findPathsPart1(data[YOU])
	// fmt.Println("The solution for part 1 is: ", part1)

	findPathsPart2(SVR, []string{})
	part2 := 0
	fmt.Println("--------------------------------")
	for _, s := range mem {
		fmt.Println(s)
		if slices.Contains(s, FFT) && slices.Contains(s, DAC) {
			part2++
		}
	}
	fmt.Println("The solution for part 2 is: ", part2)

	// idea(OUT)
	// for _, m := range memorazation {
	// 	fmt.Println(m)
	// }
}
