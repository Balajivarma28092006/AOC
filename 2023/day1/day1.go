package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func getData(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("something went wrong: %v", err)
	}
	defer file.Close()

	var values []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		values = append(values, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("something went wrong while reading: %v", err)
	}

	return values, nil
}

func main() {
	values, err := getData("inputs.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	ans := 0

	for _, s := range values {
		first, last := -1, -1

		for _, ch := range s {
			if ch >= '0' && ch <= '9' {
				d := int(ch - '0')
				if first == -1 {
					first = d
				}
				last = d
			}
		}

		if first != -1 { // line had at least one digit
			ans += first*10 + last
		}
	}

	fmt.Println(ans)

	numMap := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
		"1":     1, "2": 2, "3": 3, "4": 4, "5": 5,
		"6": 6, "7": 7, "8": 8, "9": 9,
	}

	ans = 0
	for _, str := range values {
		var firstStr, lastStr string
		for i := 0; i < len(str); i++ {
			for word := range numMap {
				if i+len(word) <= len(str) && str[i:i+len(word)] == word {
					if len(firstStr) == 0 {
						firstStr = word
					}
					lastStr = word
				}
			}
		}
		fmt.Println(firstStr, lastStr)
		number := numMap[firstStr]*10 + numMap[lastStr]
		ans += number
	}
	fmt.Println("Part2 answer is: ", ans)
	regexp.Match()
}
