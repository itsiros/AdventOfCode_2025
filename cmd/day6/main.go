package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func doMathPart2() uint64 {

	buf, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	splited := bytes.Split(buf, []byte("\n"))

	numCols := len(splited)
	numRows := len(splited[0])

	fmt.Println("num of cols: ", numCols)
	fmt.Println("num of rows: ", numRows)
	for row := range numRows {
		for col:= range numCols {
			if string(splited[row][col]) == "" {
				fmt.Print(" ")
			} else {

				fmt.Print(string(splited[row][col]))
			}
		}
		fmt.Println()
	}
	return 0
}

func doMath(table [][]string) int {

	if len(table) == 0 {
		return 0
	}

	res := 0
	numRows := len(table)
	numCols := len(table[0])

	for col := range numCols{
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

	// total := doMath(math)
	num := doMathPart2()

	// fmt.Println("The solution for part 1 is:", total)
	fmt.Println("The solution for part 2 is:", num)
}
