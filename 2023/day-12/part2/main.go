package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type pos struct {
	i, j int
}

func main() {
	input, arr := loadInput("input.txt")

	sum := 0
	for i := 0; i < len(input); i++ {
		//Use cache to speed up
		cache := make(map[pos]int)

		var sb strings.Builder
		sb.Grow(5*len(input[i]) + 4)
		for j := 0; j < 5; j++ {
			_, err := sb.WriteString(input[i])
			check(err)

			if j != 4 {
				_, err = sb.WriteString("?")
				check(err)
			}
		}

		nums := make([]int, 0, 5*len(arr[i]))
		for j := 0; j < 5; j++ {
			nums = append(nums, arr[i]...)
		}

		sum += count(sb.String(), nums, 0, 0, cache)
	}

	fmt.Println(sum)
}

func count(str string, nums []int, i, j int, cache map[pos]int) int {
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

	if result, ok := cache[pos{i, j}]; ok {
		return result
	}

	result := 0

	if str[i] == '.' || str[i] == '?' {
		result += count(str, nums, i+1, j, cache)
	}

	if str[i] == '#' || str[i] == '?' {
		if nums[j] <= len(str[i:]) && !strings.Contains(str[i:i+nums[j]], ".") {
			if nums[j] == len(str[i:]) {
				result += count(str, nums, i+nums[j], j+1, cache)
			} else if str[i+nums[j]] != '#' {
				result += count(str, nums, i+nums[j]+1, j+1, cache)
			}
		}
	}

	cache[pos{i, j}] = result

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
