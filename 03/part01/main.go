package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"unicode"
)

type info struct {
	num   int
	valid bool
	start int
	end   int
}

func main() {
	filename := "input.txt"
	rg := regexp.MustCompile(`\d+`)

	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("%s not found", filename)
	}

	scanner := bufio.NewScanner(f)

	// line: info
	numMap := make(map[int][]info)

	// line: pos
	symbolMap := make(map[int][]int)

	line := 1
	for scanner.Scan() {
		text := scanner.Text()
		runes := []rune(text)
		// get all nums, will return start and end(exclusive) index
		matches := rg.FindAllStringIndex(text, -1)

		// check if nums are valid parts
		for _, v := range matches {
			if num, err := strconv.Atoi(text[v[0]:v[1]]); err == nil {
				left := ' '
				if v[0] > 0 {
					left = runes[v[0]-1]
				}

				right := ' '
				if v[1] < len(runes) {
					right = runes[v[1]]
				}

				numMap[line] = append(numMap[line], info{
					num:   num,
					valid: isValid(left) || isValid(right),
					start: v[0],
					end:   v[1],
				})
			}
		}

		// get positions of symbols
		for _, r := range text {
			if r != '.' && !unicode.IsDigit(r) {
				positions := make([]int, 0)

				for i, v := range text {
					if v != '.' {
						positions = append(positions, i)
					}
				}

				symbolMap[line] = positions
			}
		}
		line++
	}

	for line, positions := range symbolMap {
		lineAbove, lineAboveOk := numMap[line-1]
		lineBelow, lineBelowOk := numMap[line+1]

		for _, v := range positions {
			if lineAboveOk {
				validateAndUpdate(lineAbove, v)
			}

			if lineBelowOk {
				validateAndUpdate(lineBelow, v)
			}
		}
	}

	var sum int
	for _, v := range numMap {
		for _, v := range v {
			if v.valid {
				sum += v.num
			}
		}
	}

	fmt.Println(sum)
}

func isValid(r rune) bool {
	return r != ' ' && r != '.' && !unicode.IsDigit(r)
}

func validateAndUpdate(infos []info, position int) {
	for i := 0; i < len(infos); i++ {
		v := infos[i]
		if v.valid {
			continue
		}

		infos[i].valid = position >= (v.start-1) && position <= v.end
	}
}
