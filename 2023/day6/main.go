package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type DataRecord struct {
	Val int
	Dist int
}

func main(){
	file, err := os.Open("./test.txt")
	if err != nil {
		fmt.Printf("unable to open the file %q", err)
		return
	}
	defer file.Close()
	
	var tempVals, tempDists []int

	currentKey := ""
	
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

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
			tempVals = append(tempVals, num)
		}else if currentKey == "dist" {
			tempDists = append(tempDists, num)
		}
	}
	
	if err := scanner.Err(); err != nil {
		log.Fatalf("error scanning file: %s", err)
	}

	var records []DataRecord
	recordCount := len(tempDists)

	for i := range recordCount {
		records = append(records, DataRecord{
			Val: tempVals[i],
			Dist: tempDists[i],
		})
	}

	fmt.Println(records)
}
