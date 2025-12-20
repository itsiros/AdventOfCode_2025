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

var data [][]float64
var solution [][][]float64

func findDistance(v1, v2 []float64) float64 {
	return math.Sqrt((v2[0]-v1[0])*(v2[0]-v1[0]) +
		(v2[1]-v1[1])*(v2[1]-v1[1]) +
		(v2[2]-v1[2])*(v2[2]-v1[2]))
}

// func isEqual(vect1, vect2 []float64) bool {
// 	if len(vect1) != len(vect2) {
// 		return false
// 	}
//
// 	for i := range vect1 {
// 		if vect1[i] != vect2[i] {
// 			return false
// 		}
// 	}
// 	return true
// }

func findMinVectors() {
	for i := range data {
		for j := i + 1; j < len(data); j++ {
			num := findDistance(data[i], data[j])
			solution = append(solution, [][]float64{{num}, data[i], data[j]})
		}
	}
}

func solvePart1() {
	findMinVectors()
	sort.Slice(solution, func(i, j int) bool {
		return solution[i][0][0] < solution[j][0][0]
	})

	for s := range solution {
		fmt.Println(solution[s])
	}

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
		var vect []float64
		split := strings.SplitSeq(scanner.Text(), ",")
		for s := range split {
			num, _ := strconv.ParseFloat(s, 64)
			vect = append(vect, num)
		}
		data = append(data, vect)
	}

	solvePart1()
}
