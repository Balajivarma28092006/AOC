package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

var tempVals, tempDists []int

func part1(tempVals, tempDists []int) int {
	i := 0 // for distance iteration
	sum := 1 

	for _, val := range tempVals {
		count := 0
		for j := 1; j < val; j++ {
//			fmt.Println(j * (val - j))
			if j * (val - j) > tempDists[i] {
				count++
			}
		}
		i++
		sum *= count
	}
	return sum
}

func part2(start, end int) int {
	count := 0
	
	for i := 1; i < start; i++ {
		if i * (start - i) > end {
			count++
		}
	}
	return count
}


func main(){
	file, err := os.Open("./test.txt")
	if err != nil {
		fmt.Printf("unable to open the file %q", err)
		return
	}
	defer file.Close()
	
//	var tempVals, tempDists []int

	currentKey := ""
	
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	
	var startString, endString string
	

	for scanner.Scan() {
		fields := scanner.Text()
		if fields == "Time:" {
			currentKey = "time"
			continue
		}else if fields == "Distance:"{
			currentKey = "dist"
		}
		
		num, err := strconv.Atoi(fields)
		if err != nil {
			continue
		}

		if currentKey == "time" {
			startString += strconv.Itoa(num)
			tempVals = append(tempVals, num)
		}else if currentKey == "dist" {
			endString += strconv.Itoa(num)
			tempDists = append(tempDists, num)
		}
	}
	
	if err := scanner.Err(); err != nil {
		log.Fatalf("error scanning file: %s", err)
	}
	
	fmt.Println(startString, endString)

	start := time.Now()
	//fmt.Println(records)
	fmt.Println(part1(tempVals, tempDists))

	i, err := strconv.Atoi(startString)
	if err != nil {
		log.Fatalf("some error %q", err)
	}

	j, err := strconv.Atoi(endString)
	if err != nil {
		log.Fatalf("some error %q", err)
	}
	fmt.Println(part2(i, j))

	end := time.Now()
	
	fmt.Println(start.Sub(end))
}
