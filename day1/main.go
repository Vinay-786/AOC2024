package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
)

func main() {
	f, err := os.Open("./input.csv")
	if err != nil {
		log.Fatal("Unable to read file")
	}
	defer f.Close()

	csvFile := csv.NewReader(f)
	parseCsv, err := csvFile.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse the file")
	}

	// Part one
	var column1, column2 []float64
	for _, record := range parseCsv {
		tempcol1, _ := strconv.ParseFloat(record[0], 64)
		column1 = append(column1, (tempcol1))
		tempcol2, _ := strconv.ParseFloat(record[1], 64)
		column2 = append(column2, (tempcol2))
	}

	slices.Sort(column1)
	slices.Sort(column2)

	var total float64 = 0
	for i := 0; i < len(column1); i++ {
		diff := column1[i] - column2[i]
		if diff < 0 {
			diff = -diff
		}
		total = total + diff
	}

	fmt.Println("total diff(part1): ", total)

	// Part two
	var occurancdSummary float64 = 0
	for _, i := range column1 {
		occurenceCount := findOccurance(i, column2)
		occurancdSummary += float64(occurenceCount) * i
	}

	fmt.Println("total Occurance(part2): ", math.Abs(occurancdSummary))
}

func findOccurance(a float64, b []float64) int {
	count := 0
	for _, i := range b {
		if i == a {
			count += 1
		}
	}
	return count
}
