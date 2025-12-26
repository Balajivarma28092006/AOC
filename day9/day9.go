package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func getInput(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Error, bruh")
	}
	defer file.Close()

	var data []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	if err = scanner.Err(); err != nil {
		log.Fatal("Error bruh")
	}
	return data
}

func file_sanitize(data []string) []int {
	line := data[0] //remembered it was a line lol

	var disk []int
	fileID := 0

	for i, c := range line {
		size := int(c - '0')
		if i%2 == 0 {
			for j := 0; j < size; j++ {
				disk = append(disk, fileID)
			}
			fileID++
		} else {
			for j := 0; j < size; j++ {
				disk = append(disk, -1)
			}
		}
	}
	return disk
}

func compaction(disk []int) []int {
	l := 0
	r := len(disk) - 1

	for l < r {
		for l < r && disk[l] != -1 {
			l += 1
		}
		for l < r && disk[r] == -1 {
			r -= 1
		}

		if l < r {
			disk[l], disk[r] = disk[r], disk[l]
			l += 1
			r -= 1
		}
	}
	return disk
}

// part 1
func checksum(disk []int) int {
	sum := 0
	for i, v := range disk {
		if v != -1 {
			sum += i * v
		}
	}
	return sum
}

func main() {
	outer := checksum(compaction(file_sanitize(getInput("day9_inputs.txt"))))
	fmt.Print(outer)
}
