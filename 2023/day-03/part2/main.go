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
	rows := len(input)

	seen := map[[2]int]bool{}

	for i := 0; i < rows; i++ {
		cols := len(input[i])

		for j := 0; j < cols; j++ {
			if isSymbol(input[i][j]) {
				pair := []int{}
				for dy := -1; dy < 2; dy++ {
					for dx := -1; dx < 2; dx++ {
						y := i + dy
						x := j + dx
						if dx == 0 && dy == 0 || x < 0 || x > cols || y < 0 || y > rows {
							continue
						}

						if isDigit(input[y][x]) {
							xStart, xEnd := x, x
							for xStart >= 0 {
								if isDigit(input[y][xStart]) {
									xStart--
								} else {
									break
								}
							}

							for xEnd < cols {
								if isDigit(input[y][xEnd]) {
									xEnd++
								} else {
									break
								}
							}

							if seen[[2]int{y, xStart}] {
								continue
							}

							n, err := strconv.Atoi(string(input[y][xStart+1 : xEnd]))
							check(err)
							pair = append(pair, n)
							seen[[2]int{y, xStart}] = true
						}
					}
				}

				if len(pair) == 2 {
					sum += pair[0] * pair[1]
				}
			}
		}
	}

	fmt.Println(sum)
}

func isDigit(x byte) bool {
	if x >= '0' && x <= '9' {
		return true
	}

	return false
}

func isSymbol(x byte) bool {
	return x == '*'
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
