package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func doCalc(arr []string) uint64 {
	var res uint64 = 1

	isAdd := false
	if arr[0] == "+" {
		isAdd = true
		res = 0
	}

	for a := 1; a < len(arr); a++ {
		num, err := strconv.ParseUint(arr[a], 10, 64)
		if err != nil {
			panic(err)
		}

		if isAdd {
			res += num
		} else {
			res *= num
		}
	}
	return res
}

func doMathPart2() uint64 {

	buf, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	splited := bytes.Split(buf, []byte("\n"))
	splited = splited[:len(splited)-1]

	numCols := len(splited)
	numRows := len(splited[0])

	var table []string

	for row := range numRows {
		var s []byte
		sign := splited[len(splited)-1][row]
		if sign != ' ' {
			table = append(table, string(sign))
		}

		for col := range numCols - 1 {
			if splited[col][row] != ' ' {
				s = append(s, splited[col][row])
			}
		}

		if string(s) != "" {
			table = append(table, string(s))
		}
	}

	// for t := range table {
	// 	fmt.Println(table[t])
	// }

	var total uint64 = 0
	start := 0

	for t := range table {
		if t == 0 {
			continue
		}

		if table[t] == "+" || table[t] == "*" {
			total += doCalc(table[start:t])
			start = t
		}

	}
	total += doCalc(table[start:])
	return total
}

func doMath(table [][]string) int {

	if len(table) == 0 {
		return 0
	}

	res := 0
	numRows := len(table)
	numCols := len(table[0])

	for col := range numCols {
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
