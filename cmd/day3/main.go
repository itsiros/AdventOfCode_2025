package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

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
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		jolt := findJolts(scanner.Text())
		totalJoltagePart1 += jolt
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	fmt.Println("Part 1 : ", totalJoltagePart1)
}
