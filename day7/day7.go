package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Error you need to provide file as args: ")
		return
	}
	fp := os.Args[1]
	file, err := os.Open(fp)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to open the file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	total := 0
	part2 := true
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		target, nums := parseLine(line)
		if canMatch(target, nums, part2) {
			total += target
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Scanner error:", err)
		return
	}
	fmt.Println("Answer:", total)
}

func parseLine(line string) (int, []int) {
	parts := strings.Split(line, ":")
	target, _ := strconv.Atoi(strings.TrimSpace(parts[0]))
	rawNums := strings.Fields(strings.TrimSpace(parts[1]))

	nums := make([]int, 0, len(rawNums))
	for _, n := range rawNums {
		val, _ := strconv.Atoi(n)
		nums = append(nums, val)
	}
	return target, nums
}

func canMatch(target int, nums []int, part2 bool) bool {
	return dfs(nums[0], nums[1:], target, part2)
}

func dfs(current int, remaining []int, target int, part2 bool) bool {
    sc = bufio.
}
