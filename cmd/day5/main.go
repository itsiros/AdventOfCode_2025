package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func isFresh(ingredient string, database [][]string) bool {

	for data := range database {
		if len(ingredient) < len(database[data][0]) || len(ingredient) > len(database[data][1]) {
			continue
		}

		if ingredient >= database[data][0] && ingredient <= database[data][1] {
			return true
		}
	}

	return false
}

func sortAndSmooth(data [][]string) uint64 {

	var ranges [][]uint64

	for _, d := range data {
		var s, e uint64
		fmt.Sscan(d[0], &s)
		fmt.Sscan(d[1], &e)
		ranges = append(ranges, []uint64{s, e})
	}

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i][0] < ranges[j][0]
	})

	var merged [][]uint64
	for _, r := range ranges {
		if len(merged) == 0 || r[0] > merged[len(merged)-1][1]+1 {
			merged = append(merged, r)
		} else if r[1] > merged[len(merged)-1][1] {
			merged[len(merged)-1][1] = r[1]

		}
	}

	var total uint64 = 0
	for i := range merged {
		total += merged[i][1] - merged[i][0] + 1
		fmt.Println(merged[i])
	}
	return total
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

		split := strings.Split(scanner.Text(), "-")
		database = append(database, split)

	}

	count := 0
	for scanner.Scan() {
		if isFresh(scanner.Text(), database) {
			count++
		}
	}

	part2 := sortAndSmooth(database)
	fmt.Println("Solution for part 1 is :", count)
	fmt.Println("Solution for part 2 is :", part2)
}
