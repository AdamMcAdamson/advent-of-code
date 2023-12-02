package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	instr := read("input.txt")

	lines := strings.Split(instr, "\r\n")

	colorReg, _ := regexp.Compile("(blue)|(red)|(green)")
	numberReg, _ := regexp.Compile("[0-9]+")

	max := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}

	sum := 0

	for i := range lines {

		game := strings.Split(lines[i], ":")
		hands := strings.Split(game[1], ";")

		for j := range hands {

			shows := strings.Split(hands[j], ",")

			for k := range shows {

				color := colorReg.FindString(shows[k])
				num := sti(numberReg.FindString(shows[k]))

				if num > max[color] {
					max[color] = num
				}
			}
		}

		temp := 1
		for _, val := range max {
			temp *= val
		}

		sum += temp
		max = map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}

	}

	fmt.Printf("%v", sum)
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
	inbyte, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Printf("Error reading input file %v\n", file)
	}
	return string(inbyte)
}
