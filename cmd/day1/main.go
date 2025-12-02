package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const STARTING_POS = 50
const RANGE = 100
const CLICK = 0
const LEFT = "L"
const RIGHT = "R"

func findNextPos(line string, pos int) (int, error) {

	if len(line) < 2 {
		return 0, fmt.Errorf("line with wrong format %s", line)
	}

	if !strings.HasPrefix(line, RIGHT) && !strings.HasPrefix(line, LEFT) {
		return 0, fmt.Errorf("line doesnt start with left or right %s", line)
	}

	num, err := strconv.Atoi(line[1:])
	if err != nil {
		return 0, fmt.Errorf("atoi failed %s", line)
	}

	if line[0] == RIGHT[0] {
		pos += num
	} else {
		pos -= num
	}

	pos = (pos%RANGE + RANGE) % RANGE
	fmt.Println("Pos is ", pos)
	return pos, nil
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
	count := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())

		currentPos, err = findNextPos(scanner.Text(), currentPos)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}

		if currentPos == CLICK {
			count++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "scanner error: ", err)
	}

	fmt.Println("The password is ", count)
}
