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
	instructions, directions := loadInput("input.txt")

	found := false
	next := "AAA"
	steps := 0

	for !found {
		for _, r := range instructions {
			steps++
			if r == 'L' {
				next = directions[next].Left
			} else if r == 'R' {
				next = directions[next].Right
			}

			if next == "ZZZ" {
				found = true
				break
			}
		}
	}

	fmt.Println(steps)
}

func check(e error) {
	if e != nil {
		log.Panic(e)
	}
}

func loadInput(fileName string) (string, map[string]Node) {
	file, err := os.Open(fileName)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	instructions := ""
	if scanner.Scan() {
		instructions = scanner.Text()
	}

	m := make(map[string]Node)

	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}

		arr := strings.Split(scanner.Text(), "=")
		k := strings.TrimSpace(arr[0])

		dirs := strings.Split(arr[1], ",")
		left := strings.TrimSpace(dirs[0])[1:]
		right := strings.TrimSpace(dirs[1])
		right = right[:len(right)-1]

		m[k] = Node{Left: left, Right: right}

	}

	return instructions, m
}
