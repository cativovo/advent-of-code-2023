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
			if num, err := strconv.Atoi(str[i : i+1]); err == nil {
				f = strconv.Itoa(num)
			}
		}

		if s == "" {
			if num, err := strconv.Atoi(str[j : j+1]); err == nil {
				s = strconv.Itoa(num)
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
