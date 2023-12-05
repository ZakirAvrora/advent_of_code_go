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

type Mapping struct {
	Destination int
	Source      int
	Range       int
}

func main() {
	input, mappings := loadInput("input.txt")
	l := len(input)

	min := math.MaxInt32

	for i := 0; i < l; i++ {
		s := input[i]

		for _, m := range mappings {
			for j := 0; j < len(m); j++ {
				if s >= m[j].Source && s < m[j].Source+m[j].Range {
					s = m[j].Destination + (s - m[j].Source)
					break
				}
			}
		}

		if s < min {
			min = s
		}
	}

	fmt.Println("Min", min)
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

func loadInput(fileName string) ([]int, [][]Mapping) {
	file, err := os.Open(fileName)
	check(err)
	defer file.Close()

	seeds := []int{}
	mappings := [][]Mapping{}

	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		arr := strings.Split(scanner.Text(), ":")
		nums := strings.Split(strings.TrimSpace(arr[1]), " ")
		for _, r := range nums {
			seed, err := strconv.Atoi(r)
			check(err)
			seeds = append(seeds, seed)
		}
	}

	m := []Mapping{}
	for scanner.Scan() {
		if scanner.Text() == "" {
			continue

		}

		if strings.Contains(scanner.Text(), ":") {
			for scanner.Scan() {
				if scanner.Text() == "" {
					break

				}

				nums := strings.Split(strings.TrimSpace(scanner.Text()), " ")

				dest, err := strconv.Atoi(nums[0])
				check(err)
				sour, err := strconv.Atoi(nums[1])
				check(err)
				r, err := strconv.Atoi(nums[2])
				check(err)

				m = append(m, Mapping{Destination: dest, Source: sour, Range: r})

			}

			mappings = append(mappings, m)
			m = []Mapping{}
		}

	}

	return seeds, mappings
}
