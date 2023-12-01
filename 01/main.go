package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("input.txt not found")
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var calibrationSum int
	for scanner.Scan() {
		calibrationSum += getCalibrationValue(scanner.Text())
	}

	fmt.Println(calibrationSum)
}

func getCalibrationValue(str string) int {
	var f string
	var s string

	for i := 0; i < len(str); i++ {
		j := len(str) - 1 - i

		if f != "" && s != "" {
			break
		}

		if f == "" {
			word := str[:i]
			num := wordToNum(word)

			if num > -1 {
				f = strconv.Itoa(num)
			} else {
				c := str[i : i+1]

				if _, err := strconv.Atoi(c); err == nil {
					f = c
				}
			}
		}

		if s == "" {
			word := str[j:]
			num := wordToNum(word)

			if num > -1 {
				s = strconv.Itoa(num)
			} else {
				c := str[j : j+1]

				if _, err := strconv.Atoi(c); err == nil {
					s = c
				}
			}
		}
	}

	resultStr := f + s

	if f == "" {
		resultStr = strings.Repeat(s, 2)
	}

	if s == "" {
		resultStr = strings.Repeat(f, 2)
	}

	result, err := strconv.Atoi(resultStr)
	if err != nil {
		return 0
	}

	return result
}

func wordToNum(word string) int {
	numsInWords := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	i := slices.IndexFunc(numsInWords, func(s string) bool {
		return strings.Contains(word, s)
	})

	if i > -1 {
		return i + 1
	}

	return i
}
