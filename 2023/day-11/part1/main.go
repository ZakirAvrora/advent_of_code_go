package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Cordinate struct {
	row, col int
}

func main() {
	input := loadInput("input.txt")

	rows := len(input)
	cols := len(input[0])

	nonEmptyRows := make(map[int]bool, rows)
	nonEmptyCols := make(map[int]bool, cols)

	galaxies := []Cordinate{}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if input[r][c] == '#' {
				nonEmptyRows[r] = true
				nonEmptyCols[c] = true

				galaxies = append(galaxies, Cordinate{row: r, col: c})
			}
		}
	}

	sum := 0
	for i := 0; i < len(galaxies)-1; i++ {

		for j := i + 1; j < len(galaxies); j++ {

			steps := 0
			minR, maxR := galaxies[i].row, galaxies[j].row
			minC, maxC := galaxies[i].col, galaxies[j].col

			if galaxies[i].row > galaxies[j].row {
				minR = galaxies[j].row
				maxR = galaxies[i].row
			}

			if galaxies[i].col > galaxies[j].col {
				minC = galaxies[j].col
				maxC = galaxies[i].col
			}

			for i := minC; i < maxC; i++ {
				if !nonEmptyCols[i] {
					steps++
				}
			}

			for i := minR; i < maxR; i++ {
				if !nonEmptyRows[i] {
					steps++
				}
			}

			steps = steps + (maxR - minR) + (maxC - minC)

			sum = sum + steps
		}

	}

	fmt.Println(sum)
}

func check(e error) {
	if e != nil {
		log.Panic(e)
	}
}

func loadInput(fileName string) []string {
	file, err := os.Open(fileName)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	input := []string{}

	for scanner.Scan() {
		str := scanner.Text()
		input = append(input, str)
	}

	return input
}
