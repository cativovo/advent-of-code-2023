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
	filename := "input.txt"
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("%v not found", filename)
	}
	defer f.Close()

	cubeMaxCountMap := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	scanner := bufio.NewScanner(f)

	var gameIdSum int

ScannerLoop:
	for scanner.Scan() {
		line := scanner.Text()
		index := strings.Index(line, ":")

		cubes := line[index+1:]

		for _, v1 := range strings.Split(cubes, ";") {
			for _, v2 := range strings.Split(v1, ",") {
				cube := strings.Split(strings.Trim(v2, " "), " ")
				countStr := cube[0]
				color := cube[1]

				if count, err := strconv.Atoi(countStr); err == nil {
					if count > cubeMaxCountMap[color] {
						// go to next line since the current line is not valid
						continue ScannerLoop
					}
				} else {
					log.Fatalf("invalid count %s color %s", countStr, color)
				}
			}
		}

		if gameId, err := strconv.Atoi(strings.Split(line[:index], " ")[1]); err == nil {
			gameIdSum += gameId
		}
	}

	fmt.Println(gameIdSum)
}
