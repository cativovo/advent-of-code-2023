package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

func main() {
	filename := "input.txt"

	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("%s not found", filename)
	}

	scanner := bufio.NewScanner(f)

	cards := make([]string, 0)

	for scanner.Scan() {
		cards = append(cards, scanner.Text())
	}

	fmt.Println(countCards(cards))
}

func countCards(cards []string) int {
	cardsCount := len(cards)

	for i, card := range cards {
		points := validate(card)
		if points > 0 {
			cardsCount += countCards(cards[i+1 : i+1+points])
		}
	}

	return cardsCount
}

func validate(s string) int {
	numbers := strings.Split(s, "|")
	winningNumbers := strings.Split(numbers[0], " ")
	myNumbers := strings.Split(numbers[1], " ")

	var points int

	for _, num := range myNumbers {
		num = strings.Trim(num, " ")

		if num != "" && slices.Contains(winningNumbers, num) {
			points++
		}
	}

	return points
}
