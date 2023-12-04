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

	var totalPoints int

	for scanner.Scan() {
		numbers := strings.Split(scanner.Text(), "|")
		winningNumbers := strings.Split(numbers[0], " ")

		var points int

		myNumbers := strings.Split(numbers[1], " ")

		for i := 0; i < len(myNumbers); i++ {
			num := strings.Trim(myNumbers[i], " ")

			if num != "" && slices.Contains(winningNumbers, num) {
				if points == 0 {
					points = 1
				} else {
					points *= 2
				}
			}
		}

		totalPoints += points
	}

	fmt.Println(totalPoints)
}
