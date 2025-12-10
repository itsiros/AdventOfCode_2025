package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func doMath(table [][]string) int {

	if len(table) == 0 {
		return 0
	}

	res := 0
	numRows := len(table)
	numCols := len(table[0])

	for col := 0; col < numCols; col++ {
		total := 0
		isAdd := false
		for row := 0; row < numRows-1; row++ {
			if table[numRows-1][col] == "+" {
				isAdd = true
			}
			num, err := strconv.Atoi(table[row][col])
			if err != nil {
				panic(err)
			}
			if isAdd {
				total += num
			} else {
				if total == 0 {
					total = num
					continue
				}
				total *= num
			}
		}
		fmt.Println(total)
		res += total
	}

	return res
}

func main() {

	if len(os.Args) != 2 {
		return
	}

	var math [][]string

	file, fErr := os.Open(os.Args[1])
	if fErr != nil {
		fmt.Fprintln(os.Stderr, fErr)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		math = append(math, strings.Fields(scanner.Text()))
	}

	total := 0
	total = doMath(math)

	fmt.Println("The solution for part 1 is:", total)
}
