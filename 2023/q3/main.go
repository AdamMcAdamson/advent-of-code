package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type loc struct {
	x int
	y int
}

type part struct {
	xval []int
	y    int
	num  int
}

func main() {
	instr := read("input.txt")

	lines := strings.Split(instr, "\r\n")

	grid := make([][]string, len(lines))

	numberReg, _ := regexp.Compile("[0-9]+")
	symbolReg, _ := regexp.Compile(`\*`)

	sum := 0

	var parts []part
	var symbols []loc

	for i := range lines {

		numStr := numberReg.FindAllString(lines[i], -1)
		numStrXval := numberReg.FindAllStringIndex(lines[i], -1)

		for j := range numStr {
			number := sti(numStr[j])
			xval := numStrXval[j]
			parts = append(parts, part{xval: xval, y: i, num: number})
		}

		grid[i] = strings.Split(lines[i], "")
		for j := range grid[i] {
			if symbolReg.MatchString(grid[i][j]) {
				symbols = append(symbols, loc{x: j, y: i})
			}
		}

	}

	for i := range symbols {
		x := symbols[i].x
		y := symbols[i].y

		var gearNums []int

		for _, part := range parts {
			if y <= part.y+1 && y >= part.y-1 {
				if part.xval[1] >= x && part.xval[0] <= x+1 {
					gearNums = append(gearNums, part.num)
				}
			}
		}
		if len(gearNums) == 2 {
			temp := 1
			for _, num := range gearNums {
				temp *= num
			}
			sum += temp
		}

	}
	fmt.Printf("%v\n", sum)
}

func sti(in string) int {
	out, err := strconv.Atoi(in)
	if err != nil {
		fmt.Printf("ERROR: %s\n", in)
		panic(err)
	}
	return out
}

func stf(in string) float64 {
	out, err := strconv.ParseFloat(in, 64)
	if err != nil {
		fmt.Printf("ERROR: %s\n", in)
		panic(err)
	}
	return out
}

func read(file string) string {
	inbyte, err := os.ReadFile(file)
	if err != nil {
		fmt.Printf("Error reading input file %v\n", file)
	}
	return string(inbyte)
}
