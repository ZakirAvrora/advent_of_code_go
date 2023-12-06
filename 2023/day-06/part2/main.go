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
	time, distance := loadInput("input.txt")

	count := 0

	for i := 0; i < time; i++ {
		v := i

		if v*(time-i) > distance {
			count++
		}
	}

	fmt.Println(count)
}

func isDigit(x rune) bool {
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

func loadInput(fileName string) (int, int) {
	file, err := os.Open(fileName)
	check(err)
	defer file.Close()

	str, err := io.ReadAll(file)
	check(err)

	pairs := strings.Split(string(str), "\n")

	times := strings.Split(pairs[0], ":")[1]
	distances := strings.Split(pairs[1], ":")[1]

	t := []rune{}
	for _, r := range times {
		if isDigit(r) {
			t = append(t, r)
		}
	}

	d := []rune{}
	for _, r := range distances {
		if isDigit(r) {
			d = append(d, r)
		}
	}

	time, err := strconv.Atoi(string(t))
	check(err)
	distance, err := strconv.Atoi(string(d))
	check(err)

	return time, distance
}
