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
	input, arr := loadInput("input.txt")

	sum := 0
	for i := 0; i < len(input); i++ {
		sum += count(input[i], arr[i], 0, 0)
	}

	fmt.Println(sum)
}

func count(str string, nums []int, i, j int) int {
	if i >= len(str) {
		if j >= len(nums) {
			return 1
		}

		return 0
	}

	if j >= len(nums) {
		if strings.Contains(str[i:], "#") {
			return 0
		} else {
			return 1
		}
	}

	result := 0

	if str[i] == '.' || str[i] == '?' {
		result += count(str, nums, i+1, j)
	}

	if str[i] == '#' || str[i] == '?' {
		if nums[j] <= len(str[i:]) && !strings.Contains(str[i:i+nums[j]], ".") && (nums[j] == len(str[i:]) || str[i+nums[j]] != '#') {
			if nums[j] == len(str[i:]) {
				result += count(str, nums, i+nums[j], j+1)
			} else {
				result += count(str, nums, i+nums[j]+1, j+1)
			}
		}
	}

	return result
}

func check(e error) {
	if e != nil {
		log.Panic(e)
	}
}

func loadInput(fileName string) ([]string, [][]int) {
	file, err := os.Open(fileName)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	input := []string{}
	arrange := [][]int{}

	for scanner.Scan() {
		str := scanner.Text()

		arr := strings.Split(str, " ")
		input = append(input, arr[0])

		x := []int{}

		for _, e := range strings.Split(arr[1], ",") {
			n, err := strconv.Atoi(e)
			check(err)
			x = append(x, n)
		}

		arrange = append(arrange, x)
	}

	return input, arrange
}
