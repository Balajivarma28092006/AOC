package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func findRepetition(list []int, target int) int {
	total := 0
	for _, num := range list {
		if num == target {
			total++
		}
	}
	return total
}

func main() {
	locations, err := os.ReadFile("day1_inputs.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error unable to open the file")
	}

	lines := strings.Split(strings.TrimSpace(string(locations)), "\n")

	var (
		list1 []int
		list2 []int
	)

	for _, line := range lines {
		fields := strings.Fields(line)
		if len(fields) < 2 {
			continue
		}

		n1, err1 := strconv.Atoi(fields[0])
		n2, err2 := strconv.Atoi(fields[1])

		if err1 != nil || err2 != nil {
			fmt.Println("Skipping invalid line:", line)
			continue
		}

		list1 = append(list1, n1)
		list2 = append(list2, n2)
	}

	slices.Sort(list1)
	slices.Sort(list2)

	total := 0

	for i := range list1 {
		diff := list1[i] - list2[i]
		if diff < 0 {
			diff = -diff
		}
		total += diff
	}

	duplicates := 0
	for i := range list1 {
		times := list1[i] * findRepetition(list2, list1[i])
		duplicates += times
	}
	fmt.Printf("First Answer: %v\n", total)
	fmt.Printf("Duplicates count: %v\n", duplicates)
}
