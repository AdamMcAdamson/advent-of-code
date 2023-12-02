// Advent of Code 2023

// Day       Time  Rank  Score       Time  Rank  Score
//  1    00:10:28  5140      0   00:18:46  1562      0

// Functions have been appended with _PX for given part X

// Postmortem:
//
// Not knowing how to set up and use regex in golang by heart was a real time-killer, everything else was fairly trivial.
// Handling single digit lines was completed by default on accident with the select last element of array behavior.
// AoC challenges seem to tell you how each line should resolve.. May be worth to check the line-by-line test output first, as the 1 minute wait could be brutal.
// Submitted wrong answer once for each part:
// - First was due to using the incorrect regex method (FindStringSubmatch instead of FindAllString).
// - Second due to overlapping numbers (ie. eightwo -> 82), hence commented out strings.Replace lines.
// Both could have avoided by checking output against test.
// Input file workflow is solid.
// It may be good to:
// - Have code snipets on the side (I will need to add to/adjust the collection over time)
// - Extend boilerplate to support command line argument for reading test vs real input file, and keep both throughout the question.

// Part 1
package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func main_P1() {
	instr := read_P1("input.txt")

	lines := strings.Split(instr, "\r\n")

	r, _ := regexp.Compile("[0-9]")

	sum := 0
	for i := range lines {
		matches := r.FindAllString(lines[i], -1)
		num := sti_P1(matches[0])*10 + sti_P1(matches[len(matches)-1])
		fmt.Printf("%v\n", num)
		sum += num
	}

	fmt.Printf("%v", sum)
}

func sti_P1(in string) int {
	out, err := strconv.Atoi(in)
	if err != nil {
		fmt.Printf("ERROR: %s\n", in)
		panic(err)
	}
	return out
}

func read_P1(file string) string {
	inbyte, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("Err")
	}
	return string(inbyte)
}

//-----------------

// Part 2

func main_P2() {
	instr := read_P1("input.txt")

	// istr := strings.Replace(instr, "one", "1", -1)
	// istr = strings.Replace(istr, "two", "2", -1)
	// istr = strings.Replace(istr, "three", "3", -1)
	// istr = strings.Replace(istr, "four", "4", -1)
	// istr = strings.Replace(istr, "five", "5", -1)
	// istr = strings.Replace(istr, "six", "6", -1)
	// istr = strings.Replace(istr, "seven", "7", -1)
	// istr = strings.Replace(istr, "eight", "8", -1)
	// istr = strings.Replace(istr, "nine", "9", -1)

	lines := strings.Split(instr, "\r\n")

	r, _ := regexp.Compile("(one)|(two)|(three)|(four)|(five)|(six)|(seven)|(eight)|(nine)|[0-9]")

	sum := 0
	for i := range lines {

		matches := r.FindAllString(lines[i], -1)

		num := sti_P2(matches[0])*10 + sti_P2(matches[len(matches)-1])
		// fmt.Printf("%v\n", num)
		sum += num
	}

	fmt.Printf("%v", sum)
}

func sti_P2(in string) int {

	istr := strings.Replace(in, "one", "1", -1)
	istr = strings.Replace(istr, "two", "2", -1)
	istr = strings.Replace(istr, "three", "3", -1)
	istr = strings.Replace(istr, "four", "4", -1)
	istr = strings.Replace(istr, "five", "5", -1)
	istr = strings.Replace(istr, "six", "6", -1)
	istr = strings.Replace(istr, "seven", "7", -1)
	istr = strings.Replace(istr, "eight", "8", -1)
	istr = strings.Replace(istr, "nine", "9", -1)

	out, err := strconv.Atoi(istr)
	if err != nil {
		fmt.Printf("ERROR: %s\n", in)
		panic(err)
	}
	return out
}
