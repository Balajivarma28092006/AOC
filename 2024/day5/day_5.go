package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func fixUpdates(update []int, rules map[int][]int) []int {
	inSet := make(map[int]bool)
	for _, v := range update {
		inSet[v] = true
	}

	graph := make(map[int][]int)
	indegree := make(map[int]int)

	for _, v := range update {
		indegree[v] = 0
	}

	for a, list := range rules {
		if !inSet[a] {
			continue
		}
		for _, b := range list {
			if inSet[b] {
				graph[a] = append(graph[a], b)
				indegree[b]++
			}
		}
	}

	queue := []int{}
	for node, deg := range indegree {
		if deg == 0 {
			queue = append(queue, node)
		}
	}

	order := []int{}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		order = append(order, cur)

		for _, nxt := range graph[cur] {
			indegree[nxt]--
			if indegree[nxt] == 0 {
				queue = append(queue, nxt)
			}
		}
	}

	return order
}

func toInt(s string) int {
	x, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("Invalid number: %s", err)
	}
	return x
}

func parseRules(ruleLines []string) map[int][]int {
	rules := make(map[int][]int)
	for _, line := range ruleLines {
		parts := strings.Split(line, "|")
		if len(parts) != 2 {
			log.Fatalf("Invalid rule line : %s", line)
		}

		a := toInt(parts[0])
		b := toInt(parts[1])

		rules[a] = append(rules[a], b)
	}
	return rules
}

func parseUpdates(updateLine []string) [][]int {
	updates := [][]int{}

	for _, line := range updateLine {
		nums := []int{}
		parts := strings.Split(line, ",")
		for _, p := range parts {
			nums = append(nums, toInt(p))
		}
		updates = append(updates, nums)
	}

	return updates
}

func main() {
	fp := os.Args[1]
	file, err := os.Open(fp)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var ruleLines []string   // like A|B
	var updateLines []string //like A, B etc..

	readingRules := true

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			readingRules = false // if we have a new line then switch to updatelines
			continue
		}

		if readingRules {
			ruleLines = append(ruleLines, line)
		} else {
			updateLines = append(updateLines, line)
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	isValid := func(update []int, rules map[int][]int) bool {
		pos := make(map[int]int)

		for i, v := range update {
			pos[v] = i
		}

		for a, list := range rules {
			for _, b := range list {
				ai, okA := pos[a] //position of the first
				bi, okB := pos[b] //the others

				if okA && okB {
					if ai > bi {
						return false
					}
				}
			}
		}
		return true
	}

	rules := parseRules(ruleLines)
	updates := parseUpdates(updateLines)

	//part1
	sum := 0
	for _, upd := range updates {
		if isValid(upd, rules) {
			mid := upd[len(upd)/2]
			sum += mid
		}
	}

	part2_sum := 0
	for _, upd := range updates {
		if !isValid(upd, rules) {
			corrected := fixUpdates(upd, rules)
			mid := corrected[len(corrected)/2]
			part2_sum += mid
		}
	}
	fmt.Println(part2_sum)
}
