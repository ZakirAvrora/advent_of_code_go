package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Node struct {
	Left  string
	Right string
}

func main() {
	instructions, directions, startingPoints := loadInput("input.txt")

	stepsArr := []int{}

	for _, r := range startingPoints {
		found := false
		next := r
		steps := 0

		for !found {
			for _, r := range instructions {
				steps++
				if r == 'L' {
					next = directions[next].Left
				} else if r == 'R' {
					next = directions[next].Right
				}

				if next[len(next)-1] == 'Z' {
					found = true
					break
				}
			}
		}

		stepsArr = append(stepsArr, steps)
	}

	fmt.Println(LCM(stepsArr))
}

func check(e error) {
	if e != nil {
		log.Panic(e)
	}
}

func loadInput(fileName string) (string, map[string]Node, []string) {
	file, err := os.Open(fileName)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	instructions := ""
	if scanner.Scan() {
		instructions = scanner.Text()
	}

	m := make(map[string]Node)
	startingPoints := []string{}

	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}

		arr := strings.Split(scanner.Text(), "=")
		k := strings.TrimSpace(arr[0])
		if k[len(k)-1] == 'A' {
			startingPoints = append(startingPoints, k)
		}

		dirs := strings.Split(arr[1], ",")
		left := strings.TrimSpace(dirs[0])[1:]
		right := strings.TrimSpace(dirs[1])
		right = right[:len(right)-1]

		m[k] = Node{Left: left, Right: right}

	}

	return instructions, m, startingPoints
}

// Least Common Multiple
func LCM(integers []int) int {
	if len(integers) < 2 {
		return integers[0]
	}
	result := integers[0]

	for i := 1; i < len(integers); i++ {
		result = lcm(result, integers[i])
	}

	return result
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}
