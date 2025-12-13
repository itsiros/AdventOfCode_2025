package main

import (
	"bytes"
	"fmt"
	"os"
)

func countValid(input [][]byte, r, c int) int {
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

func part2(input [][]byte) int {

	temp := input
	total := 0

	for {
		count := 0

		for row := range input {
			for col := 0; col < len(input[row]); col++ {
				if input[row][col] == '.' {
					fmt.Print(".")
					continue
				}
				if countValid(input, row, col) == 1 {
					fmt.Print("x")
					temp[row][col] = '.'
					count++
				} else {
					fmt.Printf("%c", input[row][col])
				}
			}

			fmt.Print("\n")
		}
		if count == 0 {
			break
		}
		total += count
		fmt.Print("\n")

	}
	return total
}

func part1(input [][]byte) int {
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
	return count
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

	input := bytes.Fields(buf)

	part1 := part1(input)
	part2 := part2(input)
	fmt.Println("Solution to part 1 is: ", part1)
	fmt.Println("Solution to part 2 is: ", part2)
}
