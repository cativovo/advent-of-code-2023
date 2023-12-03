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

	var sumOfPower int
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		index := strings.Index(line, ":")

		cubes := line[index+1:]
		colorsHighestCount := make(map[string]int)

		for _, v1 := range strings.Split(cubes, ";") {
			for _, v2 := range strings.Split(v1, ",") {
				cube := strings.Split(strings.Trim(v2, " "), " ")
				countStr := cube[0]
				color := cube[1]

				if count, err := strconv.Atoi(countStr); err == nil {
					if count > colorsHighestCount[color] {
						colorsHighestCount[color] = count
					}
				} else {
					log.Fatalf("invalid count %s color %s", countStr, color)
				}
			}
		}

		power := 1
		for _, v := range colorsHighestCount {
			power *= v
		}

		sumOfPower += power

	}

	fmt.Println(sumOfPower)
}
