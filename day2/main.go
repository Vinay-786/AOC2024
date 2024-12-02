package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal("Unable to read file")
	}
	defer file.Close()

	safeListCount := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		testlist := []int{}
		list := (strings.Split(scanner.Text(), " "))
		for _, i := range list {
			j, err := strconv.Atoi(i)
			if err != nil {
				log.Fatal("str to int conversion failed")
			}
			testlist = append(testlist, j)
		}
		result := CheckListPart2(testlist)
		// fmt.Println(result, ": ", testlist) //debugging
		if result {
			safeListCount++
		}
		testlist = nil
	}
	if err := scanner.Err(); err != nil {
		log.Fatal("Unable to parse the file")
	}

	fmt.Println("safeListCount: ", safeListCount)
}

func CheckListPart1(a []int) bool {
	if len(a) < 2 {
		return true // single element edge case
	}

	var diff float64
	increasing, decreasing := true, true
	for i := 0; i < len(a)-1; i++ {
		diff = float64(a[i+1] - a[i])
		if math.Abs(diff) < 1 || math.Abs(diff) > 3 {
			return false // Difference constraint violated
		}
		// fmt.Println("diff: ", diff) //debugging
		if diff > 0 {
			decreasing = false
		} else if diff < 0 {
			increasing = false
		}
	}
	return (increasing || decreasing) && !(increasing && decreasing)
}

func CheckListPart2(a []int) bool {
	if CheckListPart1(a) {
		return true
	}

	for i := 0; i < len(a); i++ {
		if CheckListPart1(skipElement(a, i)) {
			return true
		}
	}
	return false
}

func skipElement(a []int, i int) []int {
	result := make([]int, 0, len(a)-1)
	result = append(result, a[:i]...)
	result = append(result, a[i+1:]...)
	return result
}
