package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := loadInput("input.txt")
	sum := 0

	for _, records := range input {
		beginning := []int{}

		arr := records

		for len(arr) != 1 {
			l := len(arr)
			beginning = append(beginning, arr[0])
			for i := 0; i < l-1; i++ {
				arr[i] = arr[i+1] - arr[i]
			}

			arr = arr[:l-1]
		}

		next := 0
		for j := len(beginning) - 1; j >= 0; j-- {
			next = beginning[j] - next
		}

		sum += next

	}

	fmt.Println(sum)
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

	scanner := bufio.NewScanner(file)

	input := [][]int{}

	for scanner.Scan() {

		nums := strings.Split(scanner.Text(), " ")
		records := make([]int, 0, len(nums))

		for _, num := range nums {
			n, err := strconv.Atoi(num)
			check(err)
			records = append(records, n)
		}

		input = append(input, records)
	}

	return input
}
