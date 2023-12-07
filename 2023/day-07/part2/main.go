package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

var relativeStr = map[byte]int{
	'A': 12,
	'K': 11,
	'Q': 10,
	'T': 8,
	'9': 7,
	'8': 6,
	'7': 5,
	'6': 4,
	'5': 3,
	'4': 2,
	'3': 1,
	'2': 0,
	'J': -1,
}

type Card struct {
	Value    string
	Priority int
	Bid      int
}

func main() {
	cards, bids := loadInput("input.txt")
	l := len(cards)

	toSort := make([]Card, 0, l)

	for i, card := range cards {
		priority := 0
		if FiveOfKind(card) {
			priority = 6
		} else if FourIfKind(card) {
			priority = 5
		} else if FullHouse(card) {
			priority = 4
		} else if ThreeOfKind(card) {
			priority = 3
		} else if TwoPairs(card) {
			priority = 2
		} else if OnePair(card) {
			priority = 1
		} else if HighCard(card) {
			priority = 0
		} else {

			log.Panic("something wrong", card)
		}

		toSort = append(toSort, Card{
			Value:    card,
			Priority: priority,
			Bid:      bids[i],
		})
	}

	sort.SliceStable(toSort, func(i, j int) bool {

		if toSort[i].Priority == toSort[j].Priority {
			l := len(toSort[i].Value)
			for k := 0; k < l; k++ {
				if toSort[i].Value[k] == toSort[j].Value[k] {
					continue
				}

				return relativeStr[toSort[i].Value[k]] < relativeStr[toSort[j].Value[k]]
			}
		}

		return toSort[i].Priority < toSort[j].Priority
	})

	sum := 0

	for i := 0; i < l; i++ {
		sum += (i + 1) * toSort[i].Bid
	}

	fmt.Println(sum)
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

func loadInput(fileName string) ([]string, []int) {
	file, err := os.Open(fileName)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	cards := []string{}
	bids := []int{}

	for scanner.Scan() {
		arr := strings.Split(scanner.Text(), " ")
		cards = append(cards, arr[0])

		bid, err := strconv.Atoi(arr[1])
		check(err)
		bids = append(bids, bid)
	}

	return cards, bids
}

func FiveOfKind(s string) bool {
	m := make(map[byte]int)

	countJ := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'J' {
			countJ++
		} else {
			m[s[i]]++
		}
	}

	return len(m) <= 1
}

func FourIfKind(s string) bool {
	m := make(map[byte]int)

	countJ := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'J' {
			countJ++
		} else {
			m[s[i]]++
		}
	}

	for k := range m {
		if m[k]+countJ == 4 {
			return true
		}
	}

	return false
}

func FullHouse(s string) bool {
	m := make(map[byte]int)

	countJ := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'J' {
			countJ++
		} else {
			m[s[i]]++
		}
	}

	if len(m) == 2 {
		for k := range m {
			if m[k]+countJ == 3 || m[k]+countJ == 2 {
				return true
			}
		}
	}

	return false
}

func ThreeOfKind(s string) bool {
	m := make(map[byte]int)

	countJ := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'J' {
			countJ++
		} else {
			m[s[i]]++
		}
	}

	if len(m) == 3 {
		for k := range m {
			if m[k]+countJ == 3 {
				return true
			}
		}
	}

	return false
}

func TwoPairs(s string) bool {
	m := make(map[byte]int)

	countJ := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'J' {
			countJ++
		} else {
			m[s[i]]++
		}
	}

	if len(m) == 3 {
		for k := range m {
			if m[k]+countJ == 2 {
				return true
			}
		}
	}

	return false
}

func OnePair(s string) bool {
	m := make(map[byte]int)
	countJ := 0

	for i := 0; i < len(s); i++ {
		if s[i] == 'J' {
			countJ++
		} else {
			m[s[i]]++
		}
	}

	return len(m) == 4
}

func HighCard(s string) bool {
	m := make(map[byte]bool)

	for i := 0; i < len(s); i++ {
		if s[i] == 'J' {
			return false
		}

		if m[s[i]] {
			return false
		}

		m[s[i]] = true
	}

	return true
}
