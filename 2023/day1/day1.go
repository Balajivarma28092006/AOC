package main

import (
	"bufio"
	"fmt"
	"os"
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
}
