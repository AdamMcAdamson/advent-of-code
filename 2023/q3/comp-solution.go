// Advent of Code 2023

// Day       Time  Rank  Score       Time  Rank  Score
//   3   00:46:55  5392      0   00:50:22  3447      0

// Functions have been appended with _PX for given part X

// Postmortem:
//
// Sat down with < 8 mins to problem. Started stream < 5 min seconds to problem.
// Started problem faster. Was too short sighted in my initial steps. Lost 3 minutes for premature searching.
// I stepped on my own foot, putting code in loops it shouldn't be, making false assumptions, commenting out necessary lines, etc..  5+ mins lost.. Generally this will get better with exposure/practice. Would not prescribe any corrective measure yet, as it is a balancing act between speed vs accuracy.
// Regex FindStringIndex gives start and end+1 in array of 2 elements... FindAllStringIndex is an array of those arrays.. This cost 15+ minutes.. holy shit.. it was so bad.. so much extra work, so much debugging..
// On debugging, working back from answer was great. So much better than it could have been. Still could do it more.
// Part 2 was so much better than Part 1. Solution was better structered earlier questions. Super easy to add 'added' field to numloc. Super easy to remove when not needed. Typing accuracy was main time loss ~ a minute.
// It would be good to:
// - Have code snipets on the side (I will need to add to/adjust the collection over time) (So far this is small enough that going over previous questions works better)
// -- strings.Split()
// -- regex.FindString,FindAllString,FindAllStringIndex, etc
// - Extend boilerplate to support command line argument for reading test vs real input file, and keep both throughout the question.
// I should have gotten: 30 mins, 34 mins
// I could have gotten: 20 mins, 23 mins
// Postmortems help a ton

package main

import (
	"fmt"
	"regexp"
	"strings"
)

// type loc struct {
// 	x int
// 	y int
// }

type numloc_P1 struct {
	xval  []int
	y     int
	value int
	added bool
}

func main_P1() {
	instr := read("input.txt")

	lines := strings.Split(instr, "\r\n")

	grid := make([][]string, len(lines))

	// colorReg, _ := regexp.Compile("(blue)|(red)|(green)")
	// colorReg, _ := regexp.Compile("(blue)|(red)|(green)")
	numberReg, _ := regexp.Compile("[0-9]+")
	symbolReg, _ := regexp.Compile(`[^0-9\.]`)

	sum := 0

	var numbers []numloc_P1

	for i := range lines {
		grid[i] = strings.Split(lines[i], "")
		// fmt.Printf("%v: %v\n", i, symbolReg.MatchString(lines[i]))
		numStr := numberReg.FindAllString(lines[i], -1)
		numStrXmin := numberReg.FindAllStringIndex(lines[i], -1)
		workingOff := 0

		fmt.Printf("%v\n", numStr)
		fmt.Printf("%v\n", numStrXmin)

		for j := range numStr {
			number := sti(numStr[j])
			xval := numStrXmin[j]
			length := -1
			temp := number

			for temp > 0 {
				temp /= 10
				length++
			}

			numbers = append(numbers, numloc_P1{xval: xval, y: i, value: number, added: false})
			workingOff += length

		}
	}

	fmt.Printf("%v\n", grid)
	var symbols []loc

	for i := range grid {
		for j := range grid[i] {
			// fmt.Printf("%v\n", grid[i][j])
			if symbolReg.MatchString(grid[i][j]) {
				symbols = append(symbols, loc{x: j, y: i})
			}
		}
	}

	// 467..114..
	// ...*......

	for i := range symbols {
		x := symbols[i].x
		y := symbols[i].y

		for j, numL := range numbers {
			if numL.added {
				continue
			}
			if y == numL.y+1 || y == numL.y || y == numL.y-1 {

				if numL.xval[1] >= x && numL.xval[0] <= x+1 {
					if numL.value == 114 {
						fmt.Printf("%v\n", numL)
					}
					sum += numL.value
					numbers[j].added = true
				}
			}
		}

	}
	fmt.Printf("%v\n", symbols)

	fmt.Printf("%v\n", numbers)
	fmt.Printf("%v\n", sum)
}

type numloc_P2 struct {
	xval  []int
	y     int
	value int
}

func main_P2() {
	instr := read("input.txt")

	lines := strings.Split(instr, "\r\n")

	grid := make([][]string, len(lines))

	// colorReg, _ := regexp.Compile("(blue)|(red)|(green)")
	// colorReg, _ := regexp.Compile("(blue)|(red)|(green)")
	numberReg, _ := regexp.Compile("[0-9]+")
	symbolReg, _ := regexp.Compile(`\*`)

	sum := 0

	var numbers []numloc_P2

	for i := range lines {
		grid[i] = strings.Split(lines[i], "")
		// fmt.Printf("%v: %v\n", i, symbolReg.MatchString(lines[i]))
		numStr := numberReg.FindAllString(lines[i], -1)
		numStrXmin := numberReg.FindAllStringIndex(lines[i], -1)
		workingOff := 0

		fmt.Printf("%v\n", numStr)
		fmt.Printf("%v\n", numStrXmin)

		for j := range numStr {
			number := sti(numStr[j])
			xval := numStrXmin[j]
			length := -1
			temp := number

			for temp > 0 {
				temp /= 10
				length++
			}

			numbers = append(numbers, numloc_P2{xval: xval, y: i, value: number})
			workingOff += length

		}
	}

	fmt.Printf("%v\n", grid)
	var symbols []loc

	for i := range grid {
		for j := range grid[i] {
			// fmt.Printf("%v\n", grid[i][j])
			if symbolReg.MatchString(grid[i][j]) {
				symbols = append(symbols, loc{x: j, y: i})
			}
		}
	}

	// 467..114..
	// ...*......

	for i := range symbols {
		x := symbols[i].x
		y := symbols[i].y

		var gearNums []int

		for _, numL := range numbers {
			if y == numL.y+1 || y == numL.y || y == numL.y-1 {

				if numL.xval[1] >= x && numL.xval[0] <= x+1 {
					if numL.value == 114 {
						fmt.Printf("%v\n", numL)
					}
					gearNums = append(gearNums, numL.value)
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
	fmt.Printf("%v\n", symbols)

	fmt.Printf("%v\n", numbers)
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
// 	inbyte, err := ioutil.ReadFile(file)
// 	if err != nil {
// 		fmt.Printf("Error reading input file %v\n", file)
// 	}
// 	return string(inbyte)
// }
