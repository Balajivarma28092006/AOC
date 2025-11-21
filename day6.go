package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Pos struct {
	r, c int
}

type Step struct {
	r, c, dir int
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("usage: program <input-file>")
	}

	fp := os.Args[1]
	file, err := os.Open(fp)
	if err != nil {
		log.Fatalf("Unable to open the file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	grid := [][]rune{} // for . = 46, for # = 35

	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []rune(line))
	}

	var gr int
	var gc int

	dir := 0 //0 for up 1=right 2=down and 3=left

	found := false
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[r]); c++ {
			switch grid[r][c] {
			case '^':
				gr, gc, dir = r, c, 0
				found = true
			case '>':
				gr, gc, dir = r, c, 1
				found = true
			case 'v':
				gr, gc, dir = r, c, 2
				found = true
			case '<':
				gr, gc, dir = r, c, 3
				found = true
			}
			if found {
				break
			}
		}
	}

	dr := []int{-1, 0, 1, 0}
	dc := []int{0, 1, 0, -1}

	count := 0

	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			if grid[r][c] != '.' {
				continue
			}

			grid[r][c] = '#'
			if check(grid, gr, gc, dir, dr, dc) {
				count++
			}

			grid[r][c] = '.'
		}
	}
	fmt.Println(count)
}

func check(grid [][]rune, gr, gc, dir int, dr, dc []int) bool {
	visited := map[Step]bool{}

	r, c := gr, gc
	d := dir

	for {
		st := Step{r, c, d}
		if visited[st] {
			return true
		}
		visited[st] = true
		nr := r + dr[d]
		nc := c + dc[d]

		if nr < 0 || nr >= len(grid) || nc < 0 || nc >= len(grid[0]) {
			return false
		}

		if grid[nr][nc] == '#' {
			d = (d + 1) % 4
			continue
		}

		r, c = nr, nc
	}
}
