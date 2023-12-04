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
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		s := scanner.Text()
		lines = append(lines, s)
	}

	count := 0
	for y, l := range lines {
		//fmt.Println(l)
		for x, c := range l {
			if !isNumeric(c) && c != '.' {
				nums := []int{}
				// Check up
				if y > 0 {
					if isNumeric(rune(lines[y-1][x])) {
						l, n := findNumber(lines[y-1], x)
						nums = append(nums, n)
						lines[y-1] = l
					}
				}

				// Check Up & Right
				if y > 0 && x < len(l)-1 {
					if isNumeric(rune(lines[y-1][x+1])) {
						l, n := findNumber(lines[y-1], x+1)
						nums = append(nums, n)
						lines[y-1] = l
					}
				}

				// Check right
				if x < len(l)-1 {
					if isNumeric(rune(lines[y][x+1])) {
						l, n := findNumber(lines[y], x+1)
						nums = append(nums, n)
						lines[y] = l
					}
				}

				// Check down & right
				if y < len(lines)-1 && x < len(l)-1 {
					if isNumeric(rune(lines[y+1][x+1])) {
						l, n := findNumber(lines[y+1], x+1)
						nums = append(nums, n)
						lines[y+1] = l
					}
				}

				// Check down
				if y < len(lines)-1 {
					if isNumeric(rune(lines[y+1][x])) {
						l, n := findNumber(lines[y+1], x)
						nums = append(nums, n)
						lines[y+1] = l
					}
				}

				// Check down & left
				if y < len(lines)-1 && x > 0 {
					if isNumeric(rune(lines[y+1][x-1])) {
						l, n := findNumber(lines[y+1], x-1)
						nums = append(nums, n)
						lines[y+1] = l
					}
				}

				// Check left
				if x > 0 {
					if isNumeric(rune(lines[y][x-1])) {
						l, n := findNumber(lines[y], x-1)
						nums = append(nums, n)
						lines[y] = l
					}
				}

				// Check up & left
				if x > 0 && y > 0 {
					if isNumeric(rune(lines[y-1][x-1])) {
						l, n := findNumber(lines[y-1], x-1)
						nums = append(nums, n)
						lines[y-1] = l
					}
				}

				if c == '*' && len(nums) == 2 {
					count += nums[0] * nums[1]
				}
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Count:", count)
}

func isNumeric(c rune) bool {
	n := int(c - '0')
	return n <= 9 && n >= 0
}

func findNumber(s string, x int) (string, int) {
	l := x
	r := x
	for i := x; i >= 0; i-- {
		l = i
		if !isNumeric(rune(s[i])) {
			l++

			break
		}
	}
	for i := x; i < len(s); i++ {
		r = i
		if !isNumeric(rune(s[i])) {
			r--
			break
		}
	}
	//fmt.Println(s[l : r+1])

	n, err := strconv.Atoi(s[l : r+1])
	if err != nil {
		panic(err)
	}

	fmt.Println(n)
	return fmt.Sprint(s[:l], strings.Repeat(".", r-l+1), s[r+1:]), n
}
