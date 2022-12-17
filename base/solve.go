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

func part1(data string) {

	dataArray := strings.Split(data, "\n")

	fmt.Println(dataArray)
}

func part2(data string) {
	dataArray := strings.Split(data, "\n")

	fmt.Println(dataArray)
}

func main() {

	dat, err := os.ReadFile("./input")
	check(err)
	data := string(dat)

	// part1(data)

	part2(data)
}
