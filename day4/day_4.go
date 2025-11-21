package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("day4_inputs.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: opening file: %v", err)
		return
	}
	defer file.Close()

	var grid []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			grid = append(grid, line)
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	n, m := len(grid), len(grid[0])

	// dirs := [8][2]int{ // the directions as left right up down, diagnals
	// 	{-1, -1}, {-1, 0}, {-1, 1},
	// 	{0, -1}, {0, 1},
	// 	{1, -1}, {1, 0}, {1, 1},
	// }

	count := 0
	//brute force the shit out
	// for i := range n {
	// 	for j := range m {
	// 		for _, d := range dirs {
	// 			valid := true
	// 			for k := 0; k < len(word); k++ {
	// 				x := i + k*d[0]
	// 				y := j + k*d[1]
	// 				if x < 0 || y < 0 || x >= n || y >= len(grid[x]) || rune(grid[x][y]) != rune(word[k]) {
	// 					valid = false
	// 					break
	// 				}
	// 			}
	// 			if valid {
	// 				count++
	// 			}
	// 		}
	// 	}
	// }
	//
	//part 2: we need mas as x so a is in center and checkthe other diagnal blocks
	for i := 1; i < n-1; i++ {
		for j := 1; j < m-1; j++ {
			if grid[i][j] != 'A' {
				continue
			}

			//diagnoal 1
			d1 := string([]byte{grid[i-1][j-1], grid[i][j], grid[i+1][j+1]})
			//diagnal 2
			d2 := string([]byte{grid[i-1][j+1], grid[i][j], grid[i+1][j-1]})
			if (d1 == "MAS" || d1 == "SAM") && (d2 == "MAS" || d2 == "SAM") {
				count++
			}
		}
	}

	fmt.Printf("%v\n", count)
}
