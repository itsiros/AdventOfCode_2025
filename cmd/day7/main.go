package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var mem [][]int

func goDown(y, x int, tree []string, debug [][]byte) int {

	if y < 0 || x < 0 || y >= len(tree) || x >= len(tree[0]) {
		return 0
	}

	for _, m := range mem {
		if m[0] == y && m[1] == x {
			return 0
		}
	}

	mem = append(mem, []int{y, x})
	debug[y][x] = '|'

	if tree[y][x] == '^' {
		return 1 + goDown(y, x-1, tree, debug) + goDown(y, x+1, tree, debug)
	}

	if tree[y][x] == '.' {
		return goDown(y+1, x, tree, debug)
	}

	return 0
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

	var tree []string
	var debug [][]byte
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		tree = append(tree, scanner.Text())
		debug = append(debug, []byte(scanner.Text()))
	}

	start := strings.Index(tree[0], "S")
	count := goDown(1, start, tree, debug)

	for line := range debug {
		fmt.Println(string(debug[line]))
	}

	fmt.Println("The solution for part 1 is: ", count)
}
