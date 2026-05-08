package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

//directions
var dirs = map[byte][2]int {
	'^': {-1, 0},
	'v': {1, 0},
	'<': {0, -1},
	'>': {0, 1},
}

func readGrid(filename string) [][]byte {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var grid [][]byte

	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []byte(line))
	}
	return grid
}

func readDirections(filename string) string{
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var moves strings.Builder

	for scanner.Scan() {
		moves.WriteString(scanner.Text())
	}

	return moves.String()
}

func move(grid [][]byte, r, c int, dir byte)(int , int) {
	dr := dirs[dir][0]
	dc := dirs[dir][1]

	nr := r + dr
	nc := c + dc

	// wall
	if grid[nr][nc] == '#' {
		return r, c
	}

	// empty
	if grid[nr][nc] == '.' {
		grid[r][c] = '.'
		grid[nr][nc] = '@'
		return nr, nc
	}

	// pushing box
	if grid[nr][nc] == 'O' {
		cr, cc := nr, nc
		
		// find end of the box chain
		for grid[cr][cc] == 'O' {	
			cr += dr
			cc += dc
		}	
		
		//blocked by wall
		if grid[cr][cc] == '#' {
			return r, c
		}

		// must be empty now
		if grid[cr][cc] == '.' {
			// move last box
			grid[cr][cc] = 'O'

			// move robot
			grid[nr][nc] = '@'
			grid[r][c] = '.'

			return nr, nc
		}
	}
	return r, c
}

func gpsSum(grid [][]byte) int {
	sum := 0

	for r := range grid {
		for c := range grid[r] {
			if grid[r][c] == 'O' {
				sum += 100*r + c
			}
		}
	}

	return sum
}

func main() {
	grid := readGrid("grid.txt")
	moves := readDirections("directions.txt")


	// first find the position of the robot
	var r, c int
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '@' {
				r, c = i, j
			}
		}
	}

	for i := 0; i < len(moves); i++ {
		r, c = move(grid, r, c, moves[i])
	}

	fmt.Println(gpsSum(grid))
}