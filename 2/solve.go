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

/*
A -> rock
B -> paper
C -> scissor

X -> rock
Y -> paper
Z -> scissor

POINTS
rock -> 1
paper -> 2
scissor -> 3
lost -> 0
draw -> 3
win -> 6
*/

func getPoints1(play string) int {
	shapePoints := map[string]int{
		"X": 1,
		"Y": 2,
		"Z": 3}

	gamePoints := map[string]map[string]int{
		"A": {
			"X": 3,
			"Y": 6,
			"Z": 0},
		"B": {
			"X": 0,
			"Y": 3,
			"Z": 6},
		"C": {
			"X": 6,
			"Y": 0,
			"Z": 3},
	}

	total := 0
	str := strings.Split(play, " ")

	total += shapePoints[str[1]]
	total += gamePoints[str[0]][str[1]]

	return total
}

func part1(data string) {

	dataArray := strings.Split(data, "\n")

	total := 0
	for i := 0; i < len(dataArray); i++ {
		total += getPoints1(dataArray[i])
	}

	fmt.Println(total)
}

func getPoints2(play string) int {
	shapePoints := map[string]int{
		"X": 0,
		"Y": 3,
		"Z": 6}

	gamePoints := map[string]map[string]int{
		"A": {
			"X": 3,
			"Y": 1,
			"Z": 2},
		"B": {
			"X": 1,
			"Y": 2,
			"Z": 3},
		"C": {
			"X": 2,
			"Y": 3,
			"Z": 1},
	}

	total := 0
	str := strings.Split(play, " ")

	total += shapePoints[str[1]]
	total += gamePoints[str[0]][str[1]]

	return total
}

func part2(data string) {
	dataArray := strings.Split(data, "\n")

	total := 0
	for i := 0; i < len(dataArray); i++ {
		total += getPoints2(dataArray[i])
	}

	fmt.Println(total)
}

func main() {

	dat, err := os.ReadFile("./input")
	check(err)
	data := string(dat)

	part1(data)

	part2(data)
}
