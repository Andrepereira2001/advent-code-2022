package main

import (
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func allDiff(str string) bool {
	for i, elem := range str {
		for j, comp := range str {
			if j == i {
				continue
			} else if elem == comp {
				return false
			}
		}
	}
	return true
}

func part1(data string) {

	dataArray := strings.Split(data, "\n")
	res := 0
	for _, elem := range dataArray {
		fmt.Println(elem)
		for i := 4; i < len(elem)+1; i++ {
			if allDiff(elem[i-4 : i]) {
				res = i
				break
			}
		}
		fmt.Println(res)
	}

}

func part2(data string) {
	dataArray := strings.Split(data, "\n")

	res := 0
	for _, elem := range dataArray {
		fmt.Println(elem)
		for i := 14; i < len(elem)+1; i++ {
			if allDiff(elem[i-14 : i]) {
				res = i
				break
			}
		}
		fmt.Println(res)
	}

}

func main() {

	dat, err := os.ReadFile("./input")
	check(err)
	data := string(dat)

	// part1(data)

	part2(data)
}
