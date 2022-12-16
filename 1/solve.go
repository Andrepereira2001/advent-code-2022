package main

import (
    "fmt"
    "os"
    "strings"
    "strconv"
    "sort"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func maxArray(x int , array []int, size int ) []int {
    if (len(array) < size){
        array = append(array, x)
    } else if (array[0] < x){
        array[0] = x
    }
    sort.Ints(array)
    return array
}

func sumArray(array[] int) int {
    result := 0
    for _, v := range array {
        result += v
    }
    return result
}

func part1(data string){
    dataArray := strings.Split(data,"\n")
    sum := 0
    best := 0
    for i := 0; i < len(dataArray); i++ {
        if (dataArray[i] == ""){
            best = max(sum, best)
            sum = 0
        } else {
            intVal, _ := strconv.Atoi(dataArray[i])
            sum += intVal
        }
    }
    fmt.Println(max(sum, best))
}

func part2(data string){
    dataArray := strings.Split(data,"\n")
    sum := 0
    best := []int{}
    for i := 0; i < len(dataArray); i++ {
        if (dataArray[i] == ""){
            best = maxArray(sum, best, 3)
            sum = 0
        } else {
            intVal, _ := strconv.Atoi(dataArray[i])
            sum += intVal
        }
    }

    fmt.Println(sumArray(maxArray(sum, best, 3)))
}

func main() {

    dat, err := os.ReadFile("./input")
    check(err)
    data := string(dat)
    part1(data)
    
    part2(data)
}