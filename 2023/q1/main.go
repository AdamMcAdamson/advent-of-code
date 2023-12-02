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

	r, _ := regexp.Compile("(one)|(two)|(three)|(four)|(five)|(six)|(seven)|(eight)|(nine)|[0-9]")

	sum := 0
	for i := range lines {

		matches := r.FindAllString(lines[i], -1)

		num := stiWithWords(matches[0])*10 + stiWithWords(matches[len(matches)-1])
		// fmt.Printf("%v\n", num)
		sum += num
	}

	fmt.Printf("%v", sum)
}

func stiWithWords(in string) int {

	r, _ := regexp.Compile("(zero)|(one)|(two)|(three)|(four)|(five)|(six)|(seven)|(eight)|(nine)|[0-9]+")
	matches := r.FindAllString(in, -1)

	istr := ""

	for i := range matches {
		switch matches[i] {
		case "zero":
			istr += "0"
		case "one":
			istr += "1"
		case "two":
			istr += "2"
		case "three":
			istr += "3"
		case "four":
			istr += "4"
		case "five":
			istr += "5"
		case "six":
			istr += "6"
		case "seven":
			istr += "7"
		case "eight":
			istr += "8"
		case "nine":
			istr += "9"
		default:
			istr += matches[i]
		}
	}

	out, err := strconv.Atoi(istr)
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
