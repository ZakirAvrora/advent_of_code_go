package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type set struct {
	B int
	G int
	R int
}

func main() {
	input := loadInput("input.txt")

	count := 0
	for _, game := range input {

		maxRed, maxBlue, maxGreen := 0, 0, 0
		for _, set := range game {
			if set.G > maxGreen {
				maxGreen = set.G
			}

			if set.B > maxBlue {
				maxBlue = set.B
			}

			if set.R > maxRed {
				maxRed = set.R
			}
		}

		if (maxBlue * maxGreen * maxRed) == 0 {
			panic("Something wrong")
		}
		count += (maxBlue * maxGreen * maxRed)

	}

	fmt.Println(count)
}

func check(e error) {
	if e != nil {
		log.Panic(e)
	}
}

func loadInput(fileName string) [][]set {
	input := [][]set{}

	file, err := os.Open(fileName)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		game := strings.Split(scanner.Text(), ":")

		sets := strings.Split(game[1], ";")

		arr := []set{}

		for _, s := range sets {

			cubes := strings.Split(strings.TrimSpace(s), ",")

			newSet := set{}

			for _, c := range cubes {
				pair := strings.Split(strings.TrimSpace(c), " ")
				num, err := strconv.Atoi(pair[0])
				check(err)

				switch pair[1] {
				case "blue":
					newSet.B = num
				case "red":
					newSet.R = num
				case "green":
					newSet.G = num
				}
			}
			arr = append(arr, newSet)

		}

		input = append(input, arr)
	}

	return input
}

func isDigit(x byte) bool {
	if x >= '0' && x <= '9' {
		return true
	}

	return false
}
