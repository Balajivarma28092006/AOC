module main

import os

struct Warehouse {
mut:
	grid [][]u8
	moves string
	row int
	col int
}

fn read_grid() [][]u8 {
    lines := os.read_lines('grid.txt') or {
        panic('Cannot open grid.txt')
    }

    mut grid := [][]u8{}

    for line in lines {
        if line != '' {
            grid << line.bytes()
        }
    }

    return grid
}

fn read_moves() string {
    text := os.read_file('directions.txt') or {
        panic('Cannot open directions.txt')
    }

    mut result := ''

    for ch in text {
        if ch in [`^`, `v`, `<`, `>`] {
            result += ch.ascii_str()
        }
    }

    return result
}

fn (mut w Warehouse) find_robot() {
    for r in 0 .. w.grid.len {
        for c in 0 .. w.grid[r].len {
            if w.grid[r][c] == `@` {
                w.row = r
                w.col = c
                return
            }
        }
    }
	panic('Robot "@" not found in grid')
}

fn (mut w Warehouse) move(dir u8) {
    directions := {
        `^`: [-1, 0]
        `v`: [1, 0]
        `<`: [0, -1]
        `>`: [0, 1]
    }

    dr := directions[dir][0]
    dc := directions[dir][1]

    nr := w.row + dr
    nc := w.col + dc

    if w.grid[nr][nc] == `#` {
        return
    }

    if w.grid[nr][nc] == `.` {
        w.grid[w.row][w.col] = `.`
        w.grid[nr][nc] = `@`

        w.row = nr
        w.col = nc
        return
    }

    if w.grid[nr][nc] == `O` {
        mut br := nr
        mut bc := nc

        for w.grid[br][bc] == `O` {
            br += dr
            bc += dc
        }

        if w.grid[br][bc] == `.` {
            w.grid[br][bc] = `O`
            w.grid[nr][nc] = `@`
            w.grid[w.row][w.col] = `.`

            w.row = nr
            w.col = nc
        }
    }
}

fn (w Warehouse) gps_sum() int {
    mut total := 0

    for r in 0 .. w.grid.len {
        for c in 0 .. w.grid[r].len {
            if w.grid[r][c] == `O` {
                total += 100 * r + c
            }
        }
    }

    return total
}

fn main() {
    mut warehouse := Warehouse{
        grid: read_grid()
        moves: read_moves()
    }

    warehouse.find_robot()

    for ch in warehouse.moves {
        warehouse.move(ch)
    }

    println(warehouse.gps_sum())
}