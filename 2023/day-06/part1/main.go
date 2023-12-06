package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type Mapping struct {
	Destination int
	Source      int
	Range       int
}

func main() {
	input := loadInput("input.txt")
	times := input[0]
	distances := input[1]
	timeLength := len(times)

	maxTime := times[timeLength-1]
	count := 1

	poss := make([]int, timeLength)

	for i := 0; i < maxTime; i++ {
		v := i

		for j, r := range times {
			if i <= r && v*(r-i) > distances[j] {
				poss[j]++
			}
		}
	}

	for _, e := range poss {
		count *= e
	}

	fmt.Println(count)
}

func isDigit(x byte) bool {
	if x >= '0' && x <= '9' {
		return true
	}

	return false
}

func check(e error) {
	if e != nil {
		log.Panic(e)
	}
}

func loadInput(fileName string) [][]int {
	file, err := os.Open(fileName)
	check(err)
	defer file.Close()

	str, err := io.ReadAll(file)
	check(err)

	pairs := strings.Split(string(str), "\n")

	times := strings.Split(strings.TrimSpace(strings.Split(pairs[0], ":")[1]), " ")
	distances := strings.Split(strings.TrimSpace(strings.Split(pairs[1], ":")[1]), " ")

	input := [][]int{}

	arr1 := []int{}
	for _, r := range times {
		if r == "" {
			continue
		}
		time, err := strconv.Atoi(r)
		check(err)
		arr1 = append(arr1, time)
	}

	input = append(input, arr1)

	arr2 := []int{}
	for _, r := range distances {
		if r == "" {
			continue
		}
		time, err := strconv.Atoi(r)
		check(err)
		arr2 = append(arr2, time)
	}

	input = append(input, arr2)

	return input
}
