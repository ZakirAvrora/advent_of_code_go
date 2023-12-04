package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	input := loadInput("input.txt")

	sum := 0
	rows := len(input)

	m := make(map[int]int, rows)

	for i := 0; i < rows; i++ {
		m[i] = 1
	}

	for i := 0; i < rows; i++ {
		arr := strings.Split(input[i], "|")

		winning := make(map[string]bool)

		start := 0
		l := len(arr[0])

		for ; start < l; start++ {
			if arr[0][start] == ' ' {
				continue
			}

			end := start + 1

			for ; end < l; end++ {
				if arr[0][end] == ' ' {
					break
				}
			}

			winning[string(arr[0][start:end])] = true

			start = end
		}

		matchCount := 0
		start = 0
		l = len(arr[1])
		for ; start < l; start++ {
			if arr[1][start] == ' ' {
				continue
			}

			end := start + 1

			for ; end < l; end++ {
				if arr[1][end] == ' ' {
					break
				}
			}

			if winning[string(arr[1][start:end])] {
				matchCount++
			}

			start = end
		}

		for n := 0; n < m[i]; n++ {
			for k := 0; k < matchCount; k++ {
				m[i+k+1]++
			}
		}

	}

	for _, v := range m {
		sum += v
	}

	fmt.Println(sum)
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

func loadInput(fileName string) []string {
	input := []string{}

	file, err := os.Open(fileName)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		arr := strings.Split(scanner.Text(), ":")
		input = append(input, arr[1])
	}

	return input
}
