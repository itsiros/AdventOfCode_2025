package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func doCalc(arr []uint64, isAdd bool) uint64 {
	var res uint64 = 1
	if isAdd {
		res = 0
	}

	for a := range arr {
		fmt.Print(arr[a], " ")
		if isAdd {
			res += arr[a]
		} else {
			res *= arr[a]
		}
	}
	fmt.Println(" = ", res)
	return res
}

func doMathPart2() uint64 {

	buf, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	splited := bytes.Split(buf, []byte("\n"))
	splited = splited[:len(splited) - 1]

	numCols := len(splited)
	numRows := len(splited[0])

	var table [] uint64
	var sign [] byte

	for row := range numRows {
		var s []byte
		if splited[len(splited) - 1][row] != ' ' {
			sign = append(sign, splited[len(splited) - 1][row])
		}

		for col := range numCols - 1 {
			if splited[col][row] != ' ' {
				s = append(s, splited[col][row])
			}
		}

		if string(s) != "" {
			num, err := strconv.ParseUint(string(s), 10, 64)
			if err != nil {
				panic(err)
			}
			table = append(table, num)
		}
	}

	// for t := range table {
	// 	fmt.Println(table[t])
	// }
	
	i := numCols - 1
	var total uint64 = 0
	for j := 0; j * i + i <= len(table); j++ {
		isAdd := false
		if sign[j] == '+' {
			isAdd = true
		}
		total += doCalc(table[j*i:i*j+i], isAdd)
	}

	return total
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
