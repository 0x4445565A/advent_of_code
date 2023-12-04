package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	count := 0
	scanner := bufio.NewScanner(file)

	scores := map[int]int{}
	winningNumbers := map[int][]int{}
	order := []int{}
	for scanner.Scan() {
		s := scanner.Text()

		//fmt.Println(s)
		cardNumber, nums := findNumbers(s)
		score := 0
		if len(nums) > 0 {
			score = int(math.Pow(2, float64(len(nums)-1)))
		}
		winningNumbers[cardNumber] = nums
		scores[cardNumber] = score
		order = append(order, cardNumber)

		count += score
	}

	fmt.Println(order)

	orderLen := len(order)
	for i := 0; i < orderLen; i++ {
		cardNumber := order[i]
		if _, ok := winningNumbers[cardNumber]; !ok {
			continue
		}
		for ii := 1; ii <= len(winningNumbers[cardNumber]); ii++ {
			order = append(order, cardNumber+ii)

		}
		orderLen = len(order)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(len(order))
}

func isNumeric(c rune) bool {
	n := int(c - '0')
	return n <= 9 && n >= 0
}

func findNumbers(s string) (int, []int) {
	s = s + " "
	cardNumber := 0
	winning := map[int]struct{}{}
	myNumbers := []int{}

	afterDash := false

	l := 0
	num := false
	for i, c := range s {
		if num && !isNumeric(c) {
			num = false
			n, err := strconv.Atoi(s[l:i])
			if err != nil {
				panic(err)
			}
			if cardNumber == 0 {
				cardNumber = n
			} else if !afterDash {
				winning[n] = struct{}{}
			} else {
				if _, ok := winning[n]; ok {
					myNumbers = append(myNumbers, n)
				}
			}

		}

		if !num && isNumeric(c) {
			l = i
			num = true
		}

		if c == '|' {
			afterDash = true
		}
	}
	return cardNumber, myNumbers
}
