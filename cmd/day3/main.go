package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

const END = 12

var log [4]int

func findJoltsPart2(str string) uint64 {

	var total uint64 = 0
	var num int
	start := 0
	step := 4
	var err error

	for j := END; j >= 0; j-- {
		big := 0
		for i := start; i < start+step; i++ {
			if start >= len(str) || i >= len(str) {
				break
			}
			fmt.Println("i = ", i)
			num, err = strconv.Atoi(string(str[i]))
			if err != nil {
				panic(err)
			}
			if num > big {
				big = num
				start += i + 1
				step += i

			}
		}
		fmt.Println("j is: ", j)
		fmt.Println("big: ", big)
		powOfJ := math.Pow(float64(10), float64(j))
		total += uint64(big * int(powOfJ))

		fmt.Println("Total is: ", total)
	}

	return total
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
