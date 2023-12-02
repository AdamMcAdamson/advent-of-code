// Advent of Code 2023

// Day       Time  Rank  Score       Time  Rank  Score
//   2   00:20:11  4455      0   00:24:56  4042      0

// Functions have been appended with _PX for given part X

// Postmortem:
//
// Sat down with < 2 mins to problem. Started stream < 30 seconds to problem.
// Took 5 minutes to really start making progress. Definitely need to warm up before question appears, and read faster.
// It is strings.Split() Adam... You should know that..
// It is break Adam... You should know that..
// input file didn't update fast enough after putting in test data, resulting > 2 mins lost. Again, lack of regex knowledge hurt here (came to same code at the end, didn't even realize at the time).
// - Seperate test and real input files is probably the way to go.
// Regex FindString is greedy, very good to know. Grabbing numbers is easy.
// Checking line-by-line test output was good, but slow. Arrays delimited by spaces with %v... yucky
// To debug, should work back from answer. Here in part 1 I was returning the sum of gameIDs. I should have checked gameIDs that made up the sum first. Lost >90 seconds here checking everything else first..
// Golang "declared and not used" error is the worst for this..... Can't disable it... Opinionated shit is so annoying...
// Part 2 was again better than Part 1. Solution was structured worse than q1, but still decent. The way I adjusted 'sti(numberReg.FindString(shows[k]))' for each case in the switch statement was slow.
// It still would be good to:
// - Have code snipets on the side (I will need to add to/adjust the collection over time)
// -- strings.Split()
// -- regex.FindString,FindAllString,etc..
// - Extend boilerplate to support command line argument for reading test vs real input file, and keep both throughout the question.
// I should have gotten: 15 mins, 18 mins
// I could have gotten: 12 mins, 15 mins

package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func main_P1() {
	instr := read("input.txt")

	lines := strings.Split(instr, "\r\n")

	colorReg, _ := regexp.Compile("(blue)|(red)|(green)")
	numberReg, _ := regexp.Compile("[0-9]+")
	// colorReg, _ := regexp.Compile("(blue)|(red)|(green)")

	skipHand := false

	maxRed := 12
	maxGreen := 13
	maxBlue := 14

	sum := 0
	for i := range lines {
		game := strings.Split(lines[i], ":")

		gameID := sti(numberReg.FindString(game[0]))
		hands := strings.Split(game[1], ";")

		// fmt.Printf("%v\n", hands)

		for j := range hands {
			if skipHand {
				break
			}
			shows := strings.Split(hands[j], ",")
			// fmt.Printf("%v\n", shows)

			for k := range shows {
				color := colorReg.FindString(shows[k])
				switch color {
				case "blue":
					if sti(numberReg.FindString(shows[k])) > maxBlue {
						skipHand = true
						break
					}
				case "red":
					if sti(numberReg.FindString(shows[k])) > maxRed {
						skipHand = true
						break
					}
				case "green":
					if sti(numberReg.FindString(shows[k])) > maxGreen {
						skipHand = true
						break
					}
				}
			}
		}

		if !skipHand {
			sum += gameID
		}

		skipHand = false
		// fmt.Printf("%v\n", gameID)

		// fmt.Printf("%v\n", gameID)

		// matches := r.FindAllString(lines[i], -1)

	}

	fmt.Printf("%v", sum)
}

func main_P2() {
	instr := read("input.txt")

	lines := strings.Split(instr, "\r\n")

	colorReg, _ := regexp.Compile("(blue)|(red)|(green)")
	numberReg, _ := regexp.Compile("[0-9]+")
	// colorReg, _ := regexp.Compile("(blue)|(red)|(green)")

	// skipHand := false

	maxRed := 0
	maxGreen := 0
	maxBlue := 0

	red := 0
	green := 0
	blue := 0

	sum := 0
	for i := range lines {
		game := strings.Split(lines[i], ":")

		// gameID := sti(numberReg.FindString(game[0]))
		hands := strings.Split(game[1], ";")

		// fmt.Printf("%v\n", hands)

		for j := range hands {

			shows := strings.Split(hands[j], ",")
			// fmt.Printf("%v\n", shows)

			for k := range shows {
				color := colorReg.FindString(shows[k])
				switch color {
				case "blue":
					blue = sti(numberReg.FindString(shows[k]))
					if blue > maxBlue {
						maxBlue = blue
					}
				case "red":
					red = sti(numberReg.FindString(shows[k]))
					if red > maxRed {
						maxRed = red
					}
				case "green":
					green = sti(numberReg.FindString(shows[k]))
					if green > maxGreen {
						maxGreen = green
					}
				}
			}
		}

		sum += maxBlue * maxGreen * maxRed
		maxRed = 0
		maxGreen = 0
		maxBlue = 0

		// fmt.Printf("%v\n", gameID)

		// fmt.Printf("%v\n", gameID)

		// matches := r.FindAllString(lines[i], -1)

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
