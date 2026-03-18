package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	data, err := os.ReadFile("day3_inputs.txt")
	if err != nil {
		fmt.Printf("Error opening file")
	}
	input := string(data)

	//using regex
	// re := regexp.MustCompile(`mul\((\d+),(\d+)\)`) for the part 1
	// matches := re.FindAllSubString(input, -1) for the part 1
	re := regexp.MustCompile(`do\(\)|don't\(\)|mul\(\d+,\d+\)`)
	matches := re.FindAllString(input, -1)

	enabled := true
	sum := 0
	for _, token := range matches {
		switch token {
		case "do()":
			enabled = true
		case "don't()":
			enabled = false
		default:
			if enabled {
				nums := regexp.MustCompile(`\d+`).FindAllString(token, -1)
				if len(nums) == 2 {
					n, _ := strconv.Atoi(nums[0])
					m, _ := strconv.Atoi(nums[1])
					sum += n * m
				}
			}
		}
	}
	fmt.Printf("The whatever: %v", sum)
}
