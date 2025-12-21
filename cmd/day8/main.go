package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

var data [][]int
var pairs [][][]int

func findDistance(v1, v2 []int) int {
	return int(math.Sqrt(float64((v2[0]-v1[0])*(v2[0]-v1[0]) +
		(v2[1]-v1[1])*(v2[1]-v1[1]) +
		(v2[2]-v1[2])*(v2[2]-v1[2]))))
}

func isEqual(vect1, vect2 []int) bool {
	if len(vect1) != len(vect2) {
		return false
	}

	for i := range vect1 {
		if vect1[i] != vect2[i] {
			return false
		}
	}
	return true
}

func mergeList() {
	var temp [][][]int
	for i := range pairs {
		temp = append(temp, pairs[i][1:])
	}

	for i := range temp {
		if i == 0 {
			continue
		}

		for j := range temp[i] {
			if isEqual(temp[i][j], temp[i-1][j]) {
				temp[i-1] = append(temp[i-1], temp[i][j])
				temp[i] = temp[len(temp)-1]
				temp = temp[:len(temp)-1]
			}
		}
	}
	for _, t := range temp {
		fmt.Println(t)
	}
}

func findMinVectors() {
	for i := range data {
		for j := i + 1; j < len(data); j++ {
			num := findDistance(data[i], data[j])
			pairs = append(pairs, [][]int{{num}, data[i], data[j]})
		}
	}
}

func solvePart1() {
	findMinVectors()
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][0][0] < pairs[j][0][0]
	})

	mergeList()

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

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var vect []int
		split := strings.SplitSeq(scanner.Text(), ",")
		for s := range split {
			num, _ := strconv.ParseInt(s, 10, 32)
			vect = append(vect, int(num))
		}
		data = append(data, vect)
	}

	solvePart1()
}
