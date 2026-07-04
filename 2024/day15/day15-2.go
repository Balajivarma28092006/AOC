package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var dirs = map[byte][2]int{
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

func readDirections(filename string) string {
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

func expandMapFromPreviousMap(grid [][]byte) [][]byte {
	var newGrid [][]byte

	for _, row := range grid {
		var newRow []byte
		for _, ch := range row {
			switch ch {
			case '#':
				newRow = append(newRow, '#', '#')
			case '.':
				newRow = append(newRow, '.', '.')
			case 'O':
				newRow = append(newRow, '[', ']')
			case '@':
				newRow = append(newRow, '@', '.')
			}
		}
		newGrid = append(newGrid, newRow)
	}
	return newGrid
}

func main() {
	grid := expandMapFromPreviousMap(readGrid("grid.txt"))
	moves := readDirections("directions.txt")
	fmt.Print(grid)
	fmt.Print(moves)
}
