package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const END = 12

func findJoltsPart2(str string) uint64 {

	var total string
	start := 0

	for j := range END {
		var big byte = '0'
		bigIndex := 0
		step := len(str[start:]) - 11 + j
		until := start + step
		for i := start; i < until; i++ {
			if start >= len(str) || i >= len(str) {
				break
			}
			if str[i] > big {
				big = str[i]
				bigIndex = i - start
				start = i + 1
			}
		}

		step -= bigIndex
		if big >= '1' && big <= '9' {
			total += string(big)
		}
	}

	num, err := strconv.ParseUint(total, 10, 64)
	if err != nil {
		panic(err)
	}
	return num
}

func findJolts(str string) uint64 {

	res := 0
	strLen := len(str)

	for i := 0; i < strLen; i++ {
		for j := i + 1; j < strLen; j++ {

			num1, _ := strconv.Atoi(string(str[i]))
			num2, _ := strconv.Atoi(string(str[j]))

			if res < num1*10+num2 {
				res = num1*10 + num2
			}

		}
	}
	return uint64(res)
}

func main() {

	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "No file")
		return
	}

	file, fErr := os.Open(os.Args[1])
	if fErr != nil {
		fmt.Fprintln(os.Stderr, fErr)
		return
	}

	defer file.Close()

	var totalJoltagePart1 uint64 = 0
	var totalJoltagePart2 uint64 = 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		jolt := findJolts(scanner.Text())
		totalJoltagePart1 += jolt
		totalJoltagePart2 += findJoltsPart2(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	fmt.Println("Part 1 : ", totalJoltagePart1)
	fmt.Println("Part 2 : ", totalJoltagePart2)
}
