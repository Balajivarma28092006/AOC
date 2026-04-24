package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Range struct {
	start int
	end   int
}

func main() {
	file, _ := os.Open("inputs.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var ranges []Range
	var numbers []int

	reading_ranges := true

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			reading_ranges = false
			continue
		}

		if reading_ranges {
			var s, e int
			fmt.Sscanf(line, "%d-%d", &s, &e)
			ranges = append(ranges, Range{s, e})
		} else {
			var n int
			fmt.Sscanf(line, "%d", &n)
			numbers = append(numbers, n)
		}
	}

	// Part 1
	count := 0
	for _, n := range numbers {
		for _, r := range ranges {
			if n >= r.start && n <= r.end {
				count++
				break
			}
		}
	}

	fmt.Println("Part1:", count)

	// PART 2
	merged := mergeRanges(ranges)

	total := 0
	for _, r := range merged {
		total += r.end - r.start + 1
	}

	fmt.Println("Part2:", total)
}

func mergeRanges(ranges []Range) []Range {
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].start < ranges[j].start
	})

	merged := []Range{}

	for _, r := range ranges {
		if len(merged) == 0 {
			merged = append(merged, r)
		} else {
			last := &merged[len(merged)-1]

			if r.start <= last.end {
				if r.end > last.end {
					last.end = r.end
				}
			} else {
				merged = append(merged, r)
			}
		}
	}
	return merged
}
