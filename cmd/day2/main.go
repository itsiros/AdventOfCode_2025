package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func sameDigitWholeString(num string) bool {

	numLen := len(num)

	if numLen < 2 {
		return false
	}

	for i := 1; i < numLen; i++ {
		if num[0] != num[i] {
			return false
		}
	}

	return true
}

func isNumInvalidPart2(num string) bool {

	strLen := len(num)
	for i := 2; i < strLen/2; i++ {
		if strLen%i == 0 {
			cuts := strLen / i
			for j := 1; j < cuts; j++ {
				if strings.HasPrefix(num[j*i:], num[0:i]) {
					if j == cuts-1 {
						return true
					}
					continue
				} else {
					break
				}
			}
		}
	}
	return false
}

func findInvalidPart2(num string) string {

	if sameDigitWholeString(num) || isNumInvalid(num) || isNumInvalidPart2(num) {
		return num
	}

	return ""
}

func isNumInvalid(num string) bool {
	strlen := len(num)
	if strlen%2 != 0 {
		return false
	}

	leftPart := num[:strlen/2]
	rightPart := num[strlen/2:]

	if leftPart == rightPart {
		return true
	}

	return false
}

func findInvalid(line string) (uint64, uint64) {

	split := strings.SplitN(line, "-", 2)
	start, startErr := strconv.ParseUint(split[0], 10, 64)
	end, endErr := strconv.ParseUint(split[1], 10, 64)

	if len(split) < 2 {
		fmt.Println("SplitN failed in line: ", line)
		os.Exit(-1)
	}

	if startErr != nil || endErr != nil || strings.HasPrefix(split[0], "0") || strings.HasPrefix(split[1], "0") {
		fmt.Println("ParseUint faild in line: ", line)
		fmt.Println("left: ", start, " right: ", end)
		fmt.Println("left: ", split[0], " right: ", split[1])
		os.Exit(-1)
	}

	var totalPart1 uint64 = 0
	var totalPart2 uint64 = 0

	for i := start; i <= end; i++ {

		str := strconv.FormatUint(i, 10)
		str2 := findInvalidPart2(str)
		if str2 != "" {
			part2, err := strconv.ParseUint(str2, 10, 64)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
			totalPart2 += part2
		}

		if !isNumInvalid(str) {
			continue
		}

		totalPart1 += i
	}

	return totalPart1, totalPart2
}

func main() {

	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "No file")
		return
	}

	buf, readErr := os.ReadFile(os.Args[1])
	if readErr != nil {
		fmt.Fprintln(os.Stderr, readErr)
		return
	}

	s := strings.Trim(string(buf), "\n")
	splitted := strings.Split(s, ",")

	var totalPart1 uint64 = 0
	var totalPart2 uint64 = 0
	for _, split := range splitted {
		part1, part2 := findInvalid(split)
		totalPart1 += part1
		totalPart2 += part2
	}

	fmt.Println("The result for part 1 is: ", totalPart1)
	fmt.Println("The result for part 2 is: ", totalPart2)
}
