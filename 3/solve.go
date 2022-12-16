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

func getSameStr(str1, str2 string) string {
	chars := strings.Split(str1, "")

	for _, val := range chars {
		if strings.Contains(str2, val) {
			return val
		}
	}

	return ""
}

func getValue(str string) int {
	if str >= "a" && str <= "z" {
		return int(str[0]) - 96
	} else {
		return int(str[0]) - 38
	}
}

func part1(data string) {

	dataArray := strings.Split(data, "\n")
	res := 0
	for _, val := range dataArray {
		val1 := val[:len(val)/2]
		val2 := val[len(val)/2:]

		str := getSameStr(val1, val2)
		res += getValue(str)
	}

	fmt.Println(res)
}

func getSameStrArr(str1, str2 string) string {
	chars := strings.Split(str1, "")
	res := ""
	for _, val := range chars {
		if strings.Contains(str2, val) {
			res += val
		}
	}

	return res
}

func part2(data string) {
	dataArray := strings.Split(data, "\n")

	res := 0
	for i := 0; i < len(dataArray); i += 3 {
		arr := getSameStrArr(dataArray[i], dataArray[i+1])
		str := getSameStr(arr, dataArray[i+2])
		res += getValue(str)
	}

	fmt.Println(res)
}

func main() {

	dat, err := os.ReadFile("./input")
	check(err)
	data := string(dat)

	// part1(data)

	part2(data)
}
