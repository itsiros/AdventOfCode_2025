package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const STARTING_POS = 50
const MIN_POS = 0
const MAX_POS = 99
const CLICK = 0
const ERR_IN_FILE = -1
const LEFT = "L"
const RIGHT = "R"

func findNextPos(line string, pos int) int {

	if !strings.HasPrefix(line, RIGHT) && strings.HasPrefix(line, RIGHT) {
		fmt.Fprintln(os.Stderr, "line doesnt start with left or right ", line)
		return ERR_IN_FILE
	}

	num, err := strconv.Atoi(line[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, "atoi failed maybe because of new line ", line)
		return ERR_IN_FILE
	}

	if line[0] == RIGHT[0] {
		pos += num
	} else {
		pos -= num
	}

	for pos < MIN_POS || pos > MAX_POS {
		if pos > MAX_POS {
			pos %= MAX_POS + 1
		} else if pos < MIN_POS {
			pos += MAX_POS + 1
		}
	}

	fmt.Println("Pos is ", pos)
	return pos
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

		currentPos = findNextPos(string(scanner.Text()), currentPos)

		if currentPos == ERR_IN_FILE {
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
