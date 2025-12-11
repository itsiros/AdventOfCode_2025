package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var table [][]int

func ft_abs(x int) int {
	if x < 0 {
		x *= -1
	}
	return x
}

func findBiggest(table [][]int) int {

	max := 0
	for i := 0; i < len(table); i++ {
		for j := 0; j < len(table); j++ {
			if table[i][0] == table[j][0] && table[i][1] == table[j][1] {
				continue
			}
			maxRect := (ft_abs(table[i][0]-table[j][0]) + 1) * (ft_abs(table[i][1]-table[j][1]) + 1)
			if maxRect > max {
				max = maxRect
			}
		}
	}
	return max
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

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		split := strings.Split(scanner.Text(), ",")
		x, _ := strconv.Atoi(split[0])
		y, _ := strconv.Atoi(split[1])
		table = append(table, []int{x, y})
	}

	res := findBiggest(table)

	fmt.Println("The solution for part 1 is: ", res)
}
