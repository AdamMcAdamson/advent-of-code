package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	instr := read("input.txt")

	lines := strings.Split(instr, "\r\n")

	numberReg, _ := regexp.Compile("[0-9]+")

	cardCounts := make([]int, len(lines))

	for i := range cardCounts {
		cardCounts[i] = 1
	}

	sum := 0

	for i := range lines {
		cardCol := strings.Split(lines[i], ":")
		cardNums := strings.Split(cardCol[1], "|")
		winning := numberReg.FindAllString(cardNums[0], -1)
		haves := numberReg.FindAllString(cardNums[1], -1)

		count := 0
		for _, win := range winning {
			for _, num := range haves {
				if win == num {
					count++
					cardCounts[i+count] += cardCounts[i]
				}
			}
		}
	}

	for i := range cardCounts {
		sum += cardCounts[i]
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
