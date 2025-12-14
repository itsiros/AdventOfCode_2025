package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func complexMath(c []string, isAdd bool) int {

	maxLen := 0
	for i:= range c {
		numLen := len(c[i])
		if maxLen < numLen {
			maxLen = numLen
		}
	}

	var table []string

	k := 1
	for i := maxLen - 1; i >= 0; i-- {
		var num strings.Builder
		for j := range c {

			numLen := len(c[j])
			if (isAdd && i >= numLen) || (!isAdd && numLen - k < 0) {
				continue
			}

			if isAdd {
				num.WriteByte(c[j][i])
			} else {
				num.WriteByte(c[j][numLen - k])
			}
		}
		k++
		table = append(table, num.String())
	}

	res := 0
	if !isAdd {
		res = 1
	}

	for t := range table {
		num, err := strconv.Atoi(table[t])
		if err != nil {
			panic(err)
		}
		if isAdd {
			res += num
		} else {
			res *= num
		}
	}

	// if !isAdd {
	fmt.Print(table)
		fmt.Println(" Total :",res)
	// }
	return res
}

func doMathPart2(table [][]string) int  {

	if len(table) == 0 {
		return 0
	}

	res := 0
	numRows := len(table)
	numCols := len(table[0])

	for col := range numCols {
		var Colums []string
		isAdd := false
		for row := 0; row < numRows-1; row++ {
			if table[numRows-1][col] == "+" {
				isAdd = true
			}
			Colums = append(Colums, table[row][col])
		}
		res += complexMath(Colums, isAdd)
	}

	return res
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

	total := doMath(math)
	num := doMathPart2(math)

	fmt.Println("The solution for part 1 is:", total)
	fmt.Println("The solution for part 2 is:", num)
}
