package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	digits := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	count := 0
	for scanner.Scan() {
		s := scanner.Text()
		l, r := -1, -1
		for i := 0; i < len(s); i++ {
			c := int(s[i] - '0')
			if l == -1 {
				if c <= 9 && c >= 0 {
					l = c
				} else {
					for w, n := range digits {
						if s[i] == w[0] && i+len(w) <= len(s) {
							if s[i:len(w)+i] == w {
								l = n
								break
							}
						}
					}
				}
			}

			ii := len(s) - 1 - i
			c = int(s[ii] - '0')

			if r == -1 {
				if c <= 9 && c >= 0 {
					r = c
				} else {
					for w, n := range digits {
						if s[ii] == w[len(w)-1] && ii+1-len(w) >= 0 {
							if s[ii-len(w)+1:ii+1] == w {
								r = n
								break
							}
						}
					}
				}
			}

			if l != -1 && r != -1 {
				break
			}
		}
		count += (l * 10) + r
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Count:", count)
}
