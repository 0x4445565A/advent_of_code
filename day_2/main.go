package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	colors := map[string]*regexp.Regexp{
		"red":   regexp.MustCompile(`([0-9]+) (red)`),
		"green": regexp.MustCompile(`([0-9]+) (green)`),
		"blue":  regexp.MustCompile(`([0-9]+) (blue)`),
	}

	l := 1
	count := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		fmt.Println(s)
		strs := strings.Split(s, ";")
		fmt.Println(strs)

		maxGreen := 0
		maxRed := 0
		maxBlue := 0
		for _, ss := range strs {
			for color, r := range colors {
				m := r.FindStringSubmatch(ss)
				if m == nil {
					continue
				}
				fmt.Printf("%#v\n", m)
				n, err := strconv.Atoi(m[1])
				if err != nil {
					panic(err)
				}

				switch color {
				case "green":
					if n > maxGreen {
						maxGreen = n
					}
				case "red":
					if n > maxRed {
						maxRed = n
					}
				case "blue":
					if n > maxBlue {
						maxBlue = n
					}
				}
				//fmt.Println(ss, color, count, m[1])
			}
		}
		count += (maxRed * maxGreen * maxBlue)
		l++
	}
	fmt.Println(count)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
