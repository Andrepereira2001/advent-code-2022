package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func parseDash(str string) (int, int) {
	div := strings.Split(str, "-")
	res1, _ := strconv.Atoi(div[0])
	res2, _ := strconv.Atoi(div[1])
	return res1, res2
}

func isInside(s1, f1, s2, f2 int) bool {
	return s1 <= s2 && f1 >= f2
}

func part1(data string) {

	dataArray := strings.Split(data, "\n")

	sum := 0
	for _, elem := range dataArray {
		div := strings.Split(elem, ",")
		int1 := div[0]
		int2 := div[1]

		s1, f1 := parseDash(int1)
		s2, f2 := parseDash(int2)

		if isInside(s1, f1, s2, f2) {
			sum += 1
		} else if isInside(s2, f2, s1, f1) {
			sum += 1
		}
	}

	fmt.Println(sum)
}

func overlap(s1, f1, s2, f2 int) bool {
	if s1 <= s2 && f1 >= f2 {
		return true
	}
	if s1 <= s2 && f1 >= s2 {
		return true
	}

	return false
}

func part2(data string) {
	dataArray := strings.Split(data, "\n")

	sum := 0
	for _, elem := range dataArray {
		div := strings.Split(elem, ",")
		int1 := div[0]
		int2 := div[1]

		s1, f1 := parseDash(int1)
		s2, f2 := parseDash(int2)

		if overlap(s1, f1, s2, f2) {
			sum += 1
		} else if overlap(s2, f2, s1, f1) {
			sum += 1
		}
	}

	fmt.Println(sum)
}

func main() {

	dat, err := os.ReadFile("./input")
	check(err)
	data := string(dat)

	// part1(data)

	part2(data)
}
