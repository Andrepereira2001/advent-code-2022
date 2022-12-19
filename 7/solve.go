package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type file struct {
	name string
	size int
}

type dir struct {
	name    string
	dirs    map[string]*dir
	files   []file
	sizeVal int
}

func (d *dir) size() int {
	if d.sizeVal != 0 {
		return d.sizeVal
	}

	total := 0
	for _, elem := range d.dirs {
		total += elem.size()
	}
	for _, elem := range d.files {
		total += elem.size
	}
	d.sizeVal = total
	return total
}

func (d *dir) printDirs() {
	for _, elem := range d.dirs {
		fmt.Print(elem.name)
		d.printDirs()
	}
}

func (d *dir) sumSize(maxSize int) int {
	sum := 0
	for _, v := range d.dirs {
		size := v.size()
		if size < maxSize {
			sum += size
		}
		sum += v.sumSize(maxSize)
	}
	return sum
}

func (d *dir) minDiff(space int) int {
	res := d.size()
	if res < space {
		res = 70000000
	}
	for _, v := range d.dirs {
		size := v.minDiff(space)
		if size < res {
			res = size
		}

	}
	return res
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func parseLine(str string) (string, string) {
	val, cmd := "", ""
	_, err := fmt.Sscanf(str, "$ cd %s", &val)
	if err == nil {
		return "cd", val
	}

	_, err = fmt.Sscanf(str, "$ ls")
	if err == nil {
		return "ls", ""
	}

	_, err = fmt.Sscanf(str, "dir %s", &val)
	if err == nil {
		return "dir", val
	}

	_, err = fmt.Sscanf(str, "%s %s", &cmd, &val)
	if err == nil {
		return cmd, val
	}

	return "", ""
}

func part1(data string) {

	dataArray := strings.Split(data, "\n")
	cmd, val := "", ""
	stack := []*dir{}
	var curr *dir

	for _, elem := range dataArray {
		cmd, val = parseLine(elem)
		switch cmd {
		case "cd":
			if val == ".." {
				stack = stack[:len(stack)-1]
				curr = stack[len(stack)-1]
				// fmt.Println(curr)
				break
			}
			var final *dir
			if curr != nil {
				entry, prs := curr.dirs[val]
				if !prs {
					entry = &dir{name: val, dirs: make(map[string]*dir)}
				}
				final = entry
			} else {
				final = &dir{name: val, dirs: make(map[string]*dir)}
			}
			stack = append(stack, final)
			curr = stack[len(stack)-1]
			break
		case "ls":
			break
		case "dir":
			entry, prs := curr.dirs[val]
			if !prs {
				entry = &dir{name: val, dirs: make(map[string]*dir)}
				curr.dirs[val] = entry
			}
			break
		default:
			size, _ := strconv.Atoi(cmd)
			fileObj := file{name: val, size: size}
			curr.files = append(curr.files, fileObj)
			break
		}
	}

	sum := 0
	maxSize := 100000
	sum = stack[0].sumSize(maxSize)

	fmt.Println(sum)
}

func part2(data string) {
	dataArray := strings.Split(data, "\n")
	cmd, val := "", ""
	stack := []*dir{}
	var curr *dir

	for _, elem := range dataArray {
		cmd, val = parseLine(elem)
		switch cmd {
		case "cd":
			if val == ".." {
				stack = stack[:len(stack)-1]
				curr = stack[len(stack)-1]
				// fmt.Println(curr)
				break
			}
			var final *dir
			if curr != nil {
				entry, prs := curr.dirs[val]
				if !prs {
					entry = &dir{name: val, dirs: make(map[string]*dir)}
				}
				final = entry
			} else {
				final = &dir{name: val, dirs: make(map[string]*dir)}
			}
			stack = append(stack, final)
			curr = stack[len(stack)-1]
			break
		case "ls":
			break
		case "dir":
			entry, prs := curr.dirs[val]
			if !prs {
				entry = &dir{name: val, dirs: make(map[string]*dir)}
				curr.dirs[val] = entry
			}
			break
		default:
			size, _ := strconv.Atoi(cmd)
			fileObj := file{name: val, size: size}
			curr.files = append(curr.files, fileObj)
			break
		}
	}

	diff := 70000000 - stack[0].size()
	diff = 30000000 - diff

	fmt.Println(diff, stack[0].size())
	fmt.Println(stack[0].minDiff(diff))
}

func main() {

	dat, err := os.ReadFile("./input")
	check(err)
	data := string(dat)

	// part1(data)

	part2(data)
}
