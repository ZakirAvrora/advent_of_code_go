package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var m = map[string]struct {
	exist bool
	val   string
}{
	"one":   {true, "1"},
	"two":   {true, "2"},
	"three": {true, "3"},
	"four":  {true, "4"},
	"five":  {true, "5"},
	"six":   {true, "6"},
	"seven": {true, "7"},
	"eight": {true, "8"},
	"nine":  {true, "9"},
}

func main() {
	input := loadInput("input.txt")

	sum := 0

	for _, s := range input {
		str := lettersToDigits(s)

		num, err := strconv.Atoi(str)
		check(err)

		sum += num
	}

	fmt.Println(sum)
}

func lettersToDigits(s string) string {
	l := len(s)
	firstDigit := ""
	secondDigit := ""

	i := 0

	for i < l {

		if isDigit(s[i]) {
			firstDigit += string(s[i])
			break
		}

		found := false
		for k := range m {
			end := len(k)
			if (i + end) > l {
				continue
			}
			if m[s[i:i+end]].exist {
				firstDigit += m[s[i:i+end]].val
				found = true
				break
			}
		}

		if found {
			break
		} else {
			i++
		}

	}

	j := l - 1

	for j >= 0 {

		if isDigit(s[j]) {
			secondDigit += string(s[j])
			break
		}

		found := false
		for k := range m {
			end := len(k)
			if (j - end + 1) < 0 {
				continue
			}

			word := s[j-end+1 : j+1]
			if m[word].exist {
				secondDigit += m[word].val
				found = true
				break
			}
		}

		if found {
			break
		} else {
			j--
		}

	}

	return firstDigit + secondDigit
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
