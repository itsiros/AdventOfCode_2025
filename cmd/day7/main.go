package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var mem [][]int
var tree []string
var rows int
var cols int

func goDownPart2() int {
	// DP table to store path counts to each position
	dp := make([][]int, rows)
	for i := range dp {
		dp[i] = make([]int, cols)
	}

	// Start from the 'S' position
	for c, ch := range tree[0] {
		if ch == 'S' {
			dp[0][c] = 1
			break
		}
	}

	// Process each row, and update path counts for each position
	for r := 0; r < rows-1; r++ {
		for c := 0; c < cols; c++ {
			if dp[r][c] > 0 { // If there are any paths to this position
				// If current cell is '^', split paths
				if tree[r][c] == '^' {
					// Move left
					if c-1 >= 0 {
						dp[r+1][c-1] += dp[r][c]
					}
					// Move right
					if c+1 < cols {
						dp[r+1][c+1] += dp[r][c]
					}
				} else {
					// Move straight down (for '.' or 'S')
					dp[r+1][c] += dp[r][c]
				}
			}
		}
	}

	// Sum all paths that reach the last row (bottom of the grid)
	totalPaths := 0
	for c := 0; c < cols; c++ {
		totalPaths += dp[rows-1][c]
	}

	return totalPaths
}

func goDown(y, x int) int {

	if y < 0 || x < 0 || y >= len(tree) || x >= len(tree[0]) {
		return 0
	}

	for _, m := range mem {
		if m[0] == y && m[1] == x {
			return 0
		}
	}

	mem = append(mem, []int{y, x})

	if tree[y][x] == '^' {
		return 1 + goDown(y, x-1) + goDown(y, x+1)
	}

	if tree[y][x] == '.' {
		return goDown(y+1, x)
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

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		tree = append(tree, scanner.Text())
	}

	rows = len(tree)
	cols = len(tree[0])
	start := strings.Index(tree[0], "S")
	count := goDown(1, start)
	part2 := goDownPart2()

	fmt.Println("The solution for part 1 is: ", count)
	fmt.Println("The solution for part 2 is: ", part2)
}
