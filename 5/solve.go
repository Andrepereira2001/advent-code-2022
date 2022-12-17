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

func parseMove(str string) (int, int, int) {
	var total, from, to int
	fmt.Sscanf(str, "move %d from %d to %d", &total, &from, &to)
	return total, from - 1, to - 1
}

func part1(data string) {
	dataArray := strings.Split(data, "\n")
	var stacks [][]string

	for j := 0; j < len(dataArray[0]); j += 4 {
		stacks = append(stacks, []string{})
	}

	for i := 0; i < len(dataArray); i++ {
		if strings.Contains(dataArray[i], "move") {
			total, from, to := parseMove(dataArray[i])
			for j := 0; j < total; j++ {
				stack := stacks[from]
				top := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				stacks[from] = stack
				stack = append(stacks[to], top)
				stacks[to] = stack
			}
		} else {
			elem := dataArray[i]
			for j := 0; j < len(elem); j += 4 {
				val := string(elem[j+1])
				if val == "1" {
					break
				} else if val != " " {
					new := append([]string{val}, stacks[j/4]...)
					stacks[j/4] = new
				}
			}
		}
	}

	for j := 0; j*4 < len(dataArray[0]); j++ {
		fmt.Print(stacks[j][len(stacks[j])-1])
	}
}

func part2(data string) {
	dataArray := strings.Split(data, "\n")
	var stacks [][]string

	for j := 0; j < len(dataArray[0]); j += 4 {
		stacks = append(stacks, []string{})
	}

	for i := 0; i < len(dataArray); i++ {
		if strings.Contains(dataArray[i], "move") {
			total, from, to := parseMove(dataArray[i])
			stack := stacks[from]
			elem := stack[len(stack)-total:]
			stack = stack[:len(stack)-total]
			stacks[from] = stack
			stack = append(stacks[to], elem...)
			stacks[to] = stack

		} else {
			elem := dataArray[i]
			for j := 0; j < len(elem); j += 4 {
				val := string(elem[j+1])
				if val == "1" {
					break
				} else if val != " " {
					new := append([]string{val}, stacks[j/4]...)
					stacks[j/4] = new
				}
			}
		}
	}

	for j := 0; j*4 < len(dataArray[0]); j++ {
		fmt.Print(stacks[j][len(stacks[j])-1])
	}

}

func main() {

	dat, err := os.ReadFile("./input")
	check(err)
	data := string(dat)

	// part1(data)

	part2(data)
}
