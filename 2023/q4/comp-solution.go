// Advent of Code 2023

// Day       Time  Rank  Score       Time  Rank  Score
//   4   00:09:04  2656      0   00:13:38   895      0

// Functions have been appended with _PX for given part X

// Postmortem:
//
// Started stream < 5 min seconds to problem, slight to moderate headache. Was generally slow as a result.
// Started problem fast, but could have been better. Took over a minute to start working, 2 to really start..
// Was quite slow typing the split commands. Should take me  <30 seconds, not over a minute.
// Problem was much easier than previous ones.
// Typos cost me a couple minutes. Still need to work on typing accuracy and speed in go.
// On debugging, working back from answer was again great. Still could do it faster and earlier, should copy fmt.Printf("%v\n")..
// Part 2 was again much better than Part 1. Took me longer to get up and running, but immediately knew how to solve and didn't have to debug.
// Need to be sure to open qX folder in VS code. linters,etc so opinionated things tend to break otherwise. go.work file is not working..
// It would be good to:
// - Have code snipets on the side (I will need to add to/adjust the collection over time) (So far this is small enough that going over previous questions works better)
// -- strings.Split(), regex.Split()
// -- regex.FindString,FindAllString,FindAllStringIndex,Split etc
// - Extend boilerplate to support command line argument for reading test vs real input file, and keep both throughout the question.
// I should have gotten: 7 mins, 11 mins
// I could have gotten: 5 mins, 8 mins

package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main_P1() {
	instr := read("input.txt")

	lines := strings.Split(instr, "\r\n")

	numberReg, _ := regexp.Compile("[0-9]+")
	// symbolReg, _ := regexp.Compile(`\*`)

	sum := 0

	for i := range lines {

		cardCol := strings.Split(lines[i], ":")
		cardNums := strings.Split(cardCol[1], "|")
		winning := strings.Split(cardNums[0], " ")
		haves := strings.Split(cardNums[1], " ")

		count := -1
		for _, win := range winning {
			for _, num := range haves {

				if numberReg.MatchString(win) && win == num {

					// fmt.Printf("%v:%v\n", win, num)
					count++
				}
			}
		}

		val := 1
		for j := 0; j < count; j++ {
			val *= 2
		}

		if count >= 0 {
			fmt.Printf("%v\n", val)
			sum += val
		}
	}

	fmt.Printf("%v\n", sum)
}

func main_P2() {
	instr := read("input.txt")

	lines := strings.Split(instr, "\r\n")

	cardCounts := make([]int, len(lines))

	for i := range cardCounts {
		cardCounts[i] = 1
	}

	numberReg, _ := regexp.Compile("[0-9]+")
	// symbolReg, _ := regexp.Compile(`\*`)

	sum := 0

	for i := range lines {

		cardCol := strings.Split(lines[i], ":")
		cardNums := strings.Split(cardCol[1], "|")
		winning := strings.Split(cardNums[0], " ")
		haves := strings.Split(cardNums[1], " ")

		count := 0
		for _, win := range winning {
			for _, num := range haves {
				if numberReg.MatchString(win) && win == num {
					count++
					cardCounts[i+count] += cardCounts[i]
					// fmt.Printf("%v:%v\n", win, num)
				}
			}
		}
	}
	for i := range cardCounts {
		sum += cardCounts[i]
	}
	fmt.Printf("%v\n", sum)
}

// func sti(in string) int {
// 	out, err := strconv.Atoi(in)
// 	if err != nil {
// 		fmt.Printf("ERROR: %s\n", in)
// 		panic(err)
// 	}
// 	return out
// }

// func stf(in string) float64 {
// 	out, err := strconv.ParseFloat(in, 64)
// 	if err != nil {
// 		fmt.Printf("ERROR: %s\n", in)
// 		panic(err)
// 	}
// 	return out
// }

// func read(file string) string {
// 	inbyte, err := os.ReadFile(file)
// 	if err != nil {
// 		fmt.Printf("Error reading input file %v\n", file)
// 	}
// 	return string(inbyte)
// }
