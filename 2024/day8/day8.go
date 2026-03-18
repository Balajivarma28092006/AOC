package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pair struct {
	First  int
	Second int
}

func gcd(a, b int) int {
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func Ingrid(r, c, H, W int) bool {
	return r >= 0 && r < H && c >= 0 && c < W
}

func GetData(filename string) (map[rune][]Pair, []string) {
	fp, err := os.Open(filename)
	if err != nil {
		fmt.Println("Erro opening the file", err)
		os.Exit(1)
	}
	defer fp.Close()

	var text []string
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
	data := make(map[rune][]Pair)

	for r := 0; r < len(text); r++ {
		for c, ch := range text[r] {
			if ch != '.' {
				data[ch] = append(data[ch], Pair{r, c})
			}
		}
	}
	return data, text
}

func main() {
	if len(os.Args) <= 1 {
		fmt.Printf("Usage: go run day8.go filename")
		return
	}
	groups, grid := GetData(os.Args[1])

	H := len(grid)
	W := len(grid[0])

	antinodes1 := make(map[[2]int]bool)

	for _, pt := range groups {
		for i := 0; i < len(pt); i++ {
			for j := i + 1; j < len(pt); j++ {

				p1 := pt[i]
				p2 := pt[j]

				dr := p2.First - p1.First
				dc := p2.Second - p1.Second

				a1r := p1.First - dr
				a1c := p1.Second - dc

				a2r := p2.First + dr
				a2c := p2.Second + dc

				if Ingrid(a1r, a1c, H, W) {
					antinodes1[[2]int{a1r, a1c}] = true
				}
				if Ingrid(a2r, a2c, H, W) {
					antinodes1[[2]int{a2r, a2c}] = true
				}
			}
		}
	}

	//part 2
	antinodes2 := make(map[[2]int]bool)

	for _, pt := range groups {
		for i := 0; i < len(pt); i++ {
			for j := i + 1; j < len(pt); j++ {
				p1 := pt[i]
				p2 := pt[j]

				dr := p2.First - p1.First
				dc := p2.Second - p1.Second

				g := gcd(dr, dc)
				stepR := dr / g
				stepC := dc / g

				antinodes2[[2]int{p1.First, p1.Second}] = true
				antinodes2[[2]int{p2.First, p2.Second}] = true

				r := p2.First
				c := p2.Second

				for {
					r += stepR
					c += stepC
					if !Ingrid(r, c, H, W) {
						break
					}
					antinodes2[[2]int{r, c}] = true
				}

				r = p1.First
				c = p1.Second
				for {
					r -= stepR
					c -= stepC
					if !Ingrid(r, c, H, W) {
						break
					}
					antinodes2[[2]int{r, c}] = true
				}
			}
		}
	}
	fmt.Println("Part1: ", len(antinodes1))
	fmt.Println("Part2 : ", len(antinodes2))
}
