package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func isFresh(ingredient string, database [][]string) int {

	for data := range database {
		if len(ingredient) < len(database[data][0]) || len(ingredient) > len(database[data][1]) {
			continue
		}

		if ingredient >= database[data][0] && ingredient <= database[data][1] {
			return 1
		}
	}

	return 0
}

func main() {
	if len(os.Args) != 2 {
		return
	}

	file, fErr := os.Open(os.Args[1])
	if fErr != nil {
		fmt.Fprintln(os.Stderr, fErr)
		return
	}
	defer file.Close()

	var database [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}

		// fmt.Println(scanner.Text())
		split := strings.Split(scanner.Text(), "-")
		database = append(database, split)

	}

	count := 0
	for scanner.Scan() {
		count += isFresh(scanner.Text(), database)
	}

	// for data := range database {
	// 	fmt.Println(database[data])
	// }

	fmt.Println("Solution for part 1 is :", count)
}
