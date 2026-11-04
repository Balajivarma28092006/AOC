package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func dfs(row []int) int {
	allZero := true 

	for _, x := range row {
		if x != 0 {
			allZero = false
			break
		}
	}

	if allZero {
		return 0
	}

	diff := make([]int, len(row) - 1)
	for i := 0; i < len(row) - 1; i++ {
		diff[i] = row[i + 1] - row[i];
	}
	return row[len(row) - 1] + dfs(diff)
}

func main() {
	file, err := os.Open("inputs.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	
	total := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		row := make([]int, len(fields))
		for i, s := range fields {
			row[i], _ = strconv.Atoi(s)
		}
		total += dfs(row)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	fmt.Println(total)
}
