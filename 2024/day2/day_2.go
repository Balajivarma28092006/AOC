package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func SafeorNot(levels []int) bool {
	if len(levels) < 2 {
		return true
	}

	direction := 0 //+1 for positive and -1 for negative

	for i := 1; i < len(levels); i++ {
		diff := levels[i] - levels[i-1]

		if diff == 0 || diff < -3 || diff > 3 {
			return false // no change or too large difference
		}

		if direction == 0 {
			if diff > 0 {
				direction = 1
			} else {
				direction = -1
			}
		} else {
			if (direction == 1 && diff < 0) || (direction == -1 && diff > 0) {
				return false
			}
		}
	}
	return true
}

func TryWithRemoving(levels []int) bool {
	if SafeorNot(levels) {
		return true
	}

	for i := range levels {
		//check by removing each i value
		newLevels := append([]int{}, levels[:i]...)    //values before the ith value
		newLevels = append(newLevels, levels[i+1:]...) //after the ith value
		if SafeorNot(newLevels) {
			return true
		}
	}
	return false
}

func main() {
	data, err := os.ReadFile("day2_inputs.txt")
	if err != nil {
		fmt.Println("something happened")
	}

	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	safeCount := 0

	for _, line := range lines {
		fields := strings.Fields(line)
		var nums []int
		for _, j := range fields {
			n, _ := strconv.Atoi(j)
			nums = append(nums, n)
		}

		if SafeorNot(nums) {
			safeCount++
		} else if TryWithRemoving(nums) {
			safeCount++
		}
	}
	fmt.Printf("Safe count is : %v", safeCount)
}
