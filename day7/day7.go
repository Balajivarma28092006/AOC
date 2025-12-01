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
	//part2 := true
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		target, nums := parseLine(line)
		if canReach(target, nums) {
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

func canReach(target int, nums []int) bool {
	return helper(nums[0], nums[1:], target)
}

func helper(current int, remaining []int, target int) bool {
	if len(remaining) == 0 {
		return current == target
	}

	next := remaining[0]
	rest := remaining[1:]

	if helper(current+next, rest, target) {
		return true
	}

	if helper(current*next, rest, target) {
		return true
	}

	conact := concatInt(current, next)
	if helper(conact, rest, target) {
		return true
	}

	return false
}

func concatInt(current, next int) int {
	s := strconv.Itoa(current) + strconv.Itoa(next)
	v, _ := strconv.Atoi(s)
	return v
}
