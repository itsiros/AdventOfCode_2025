package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const STARTING_POS = 50
const RANGE = 100
const CLICK = 0
const LEFT = "L"
const RIGHT = "R"

func findNextPos(line string, pos int) (int, int, int, error) {

	if len(line) < 2 {
		return 0, 0, 0, fmt.Errorf("line with wrong format %s", line)
	}

	if !strings.HasPrefix(line, RIGHT) && !strings.HasPrefix(line, LEFT) {
		return 0, 0, 0, fmt.Errorf("line doesnt start with left or right %s", line)
	}

	num, err := strconv.Atoi(line[1:])
	if err != nil {
		return 0, 0, 0, fmt.Errorf("atoi failed %s", line)
	}

	if line[0] != RIGHT[0] {
		num = -num
	}

	newPos := pos + num
	crosses := newPos / RANGE
	if newPos < 0 && newPos%RANGE != 0 {
		crosses--
	}

	crosses = int(math.Abs(float64(crosses)))

	if pos == 0 && line[0] != RIGHT[0] {
		crosses--
	}

	part1 := 0
	pos = (newPos%RANGE + RANGE) % RANGE
	if pos == CLICK {
		part1++
		if line[0] != RIGHT[0] {
			crosses++
		}
	}

	return pos, part1, crosses, nil
}

func main() {

	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "input file is needed")
		return
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, "couldnt open file")
		return
	}
	defer file.Close()

	currentPos := STARTING_POS
	totalPart1 := 0
	totalPart2 := 0
	countPart1 := 0
	countPart2 := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		currentPos, countPart1, countPart2, err = findNextPos(scanner.Text(), currentPos)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}

		totalPart1 += countPart1
		totalPart2 += countPart2
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "scanner error: ", err)
	}

	fmt.Println("The password for part1 is: ", totalPart1)
	fmt.Println("The password for part2 is: ", totalPart2)
}
