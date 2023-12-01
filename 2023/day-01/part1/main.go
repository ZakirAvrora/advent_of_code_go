package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	input := loadInput("input.txt")

	sum := 0

	for _, str := range input {
		l := len(str)
		start, end := 0, l-1

		for start <= end {
			if isDigit(str[start]) && isDigit(str[end]) {
				break
			}

			if isDigit(str[start]) {
				end--
				continue
			}

			if isDigit(str[end]) {
				start++
				continue
			}

			end--
			start++

		}

		num, err := strconv.Atoi(string(str[start]) + string(str[end]))
		check(err)

		sum += num
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
		input = append(input, scanner.Text())
	}

	return input
}
