package main

import (
	"fmt"
	"os"
	"strings"
)

func countValid(input []string, r, c int) int {
	count := 0
	rows := len(input)

	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			nr := r + i
			nc := c + j

			if nr < 0 || nr >= rows {
				continue
			}
			if nc < 0 || nc >= len(input[nr]) {
				continue
			}
			if i == 0 && j == 0 {
				continue
			}

			if input[nr][nc] == '@' {
				count++
			}
		}
	}

	if count < 4 {
		return 1
	}
	return 0
}

func main() {

	if len(os.Args) != 2 {
		return
	}

	buf, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	input := strings.Split(string(buf), "\n")

	count := 0
	for row := range input {
		for col := 0; col < len(input[row]); col++ {
			if input[row][col] == '.' {
				fmt.Print(".")
				continue
			}
			if countValid(input, row, col) == 1 {
				fmt.Print("x")
				count++
			} else {
				fmt.Printf("%c", input[row][col])
			}
		}
		fmt.Print("\n")
	}

	fmt.Println("Solution to part 1 is: ", count)
}
