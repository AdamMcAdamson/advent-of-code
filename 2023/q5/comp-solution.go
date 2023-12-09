// Advent of Code 2023

// Day       Time  Rank  Score       Time  Rank  Score
//   5   20:38:26  64127      0   21:33:24  39651      0

// Started at Time: 19:59

// pt 1: ~40 mins
// pt 2: ~95 mins

// Functions have been appended with _PX for given part X

// Postmortem:
//
// Bad headache night of release, went to bed right after work.
// Took this slower, this problem was much harder, and the math in part 1 required a lot a rework as not handling ranges would blow out the stack.
// Using sort.Sort and sort.Interface was a helpful find.
// I wasn't really in a rush to finish this.
// In general, I was very slow on figuring out how I wanted to represent the data and getting the types set up.
// Pen and paper for ranges was SUPER helpful. Mapped out all 6 possibilities, and handled each. Was slow, but 100% accurate. Didn't immediately, but knew that it was better to do it right.
// At this point doing problems right might have to be the way to go. They are getting too complex to be sloppy with.
// I assume sorting is basically required if you don't want a nightmare in processessing variable ranges.

// pt 1: 39:59
// pt 2: 94:58

package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

type mapEntry_P1 struct {
	source      int
	destination int
	length      int
}

func main_P1() {

	var seeds []int
	var seedSoil []mapEntry_P1
	var soilFert []mapEntry_P1
	var fertWater []mapEntry_P1
	var waterLight []mapEntry_P1
	var lightTemp []mapEntry_P1
	var tempHumid []mapEntry_P1
	var humidLoc []mapEntry_P1

	entryNum := -1

	minLoc := -1

	numberReg, _ := regexp.Compile("[0-9]+")
	instr := read("input.txt")

	lines := strings.Split(instr, "\r\n")

	for i := range lines {
		subLine := strings.Split(lines[i], ":")
		// fmt.Printf("subline: %v\n", subLine)
		if len(subLine) > 1 {
			entryNum++
			if entryNum == 0 {
				seeds = asti(numberReg.FindAllString(subLine[1], -1))
			}
			// fmt.Printf("%v:%v\n", entryNum, lines[i])
		} else if lines[i] != "" {
			switch entryNum {
			case 1:
				seedSoil = append(seedSoil, processMapEntry_P1(numberReg.FindAllString(lines[i], -1)))
			case 2:
				soilFert = append(soilFert, processMapEntry_P1(numberReg.FindAllString(lines[i], -1)))
			case 3:
				fertWater = append(fertWater, processMapEntry_P1(numberReg.FindAllString(lines[i], -1)))
			case 4:
				waterLight = append(waterLight, processMapEntry_P1(numberReg.FindAllString(lines[i], -1)))
			case 5:
				lightTemp = append(lightTemp, processMapEntry_P1(numberReg.FindAllString(lines[i], -1)))
			case 6:
				tempHumid = append(tempHumid, processMapEntry_P1(numberReg.FindAllString(lines[i], -1)))
			case 7:
				humidLoc = append(humidLoc, processMapEntry_P1(numberReg.FindAllString(lines[i], -1)))
			default:
				// fmt.Printf("line: %v\n", lines[i])
			}
		}
	}

	// fmt.Printf("%v: %v\n", "seedSoil", seedSoil)
	// fmt.Printf("%v: %v\n", "soilFert", soilFert)
	// fmt.Printf("%v: %v\n", "fertWater", fertWater)
	// fmt.Printf("%v: %v\n", "waterLight", waterLight)
	// fmt.Printf("%v: %v\n", "lightTemp", lightTemp)
	// fmt.Printf("%v: %v\n", "tempHumid", tempHumid)
	// fmt.Printf("%v: %v\n", "humidLoc", humidLoc)

	for _, seed := range seeds {
		num := mapNum_P1(seed, seedSoil)
		num = mapNum_P1(num, soilFert)
		num = mapNum_P1(num, fertWater)
		num = mapNum_P1(num, waterLight)
		num = mapNum_P1(num, lightTemp)
		num = mapNum_P1(num, tempHumid)
		num = mapNum_P1(num, humidLoc)

		if minLoc == -1 || num < minLoc {
			minLoc = num
		}
	}

	fmt.Printf("%v\n", minLoc)
}

func mapNum_P1(num int, mapEntries []mapEntry_P1) int {

	for _, entry := range mapEntries {
		max := entry.source + entry.length - 1
		if num <= max && num >= entry.source {
			return ((num - entry.source) + entry.destination)
		}
	}
	return num
}

func processMapEntry_P1(in []string) mapEntry_P1 {

	var out mapEntry_P1
	rawEntry := asti(in)

	if len(rawEntry) < 3 {
		fmt.Printf("in: %v\n", in)
	} else {
		out.destination = rawEntry[0]
		out.source = rawEntry[1]
		out.length = rawEntry[2]
	}

	return out
}

type mapEntry_P2 struct {
	sStart int
	sEnd   int
	dStart int
	dEnd   int
	length int
}

type rangeEntry_P2 struct {
	start int
	end   int
}

type rangeEntries_P2 []rangeEntry_P2

func (m rangeEntries_P2) Len() int           { return len(m) }
func (m rangeEntries_P2) Less(i, j int) bool { return m[i].start < m[j].start }
func (m rangeEntries_P2) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }

type mapEntries_P2 []mapEntry_P2

func (m mapEntries_P2) Len() int           { return len(m) }
func (m mapEntries_P2) Less(i, j int) bool { return m[i].sStart < m[j].sStart }
func (m mapEntries_P2) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }

func main_P2() {

	var seeds rangeEntries_P2
	var seedSoil mapEntries_P2
	var soilFert mapEntries_P2
	var fertWater mapEntries_P2
	var waterLight mapEntries_P2
	var lightTemp mapEntries_P2
	var tempHumid mapEntries_P2
	var humidLoc mapEntries_P2

	entryNum := -1

	// minLoc := -1

	numberReg, _ := regexp.Compile("[0-9]+")
	instr := read("input.txt")

	lines := strings.Split(instr, "\r\n")

	for i := range lines {
		subLine := strings.Split(lines[i], ":")
		// fmt.Printf("subline: %v\n", subLine)
		if len(subLine) > 1 {
			entryNum++
			if entryNum == 0 {
				seeds = processSeeds_P2(asti(numberReg.FindAllString(subLine[1], -1)))
			}
			// fmt.Printf("%v:%v\n", entryNum, lines[i])
		} else if lines[i] != "" {
			switch entryNum {
			case 1:
				seedSoil = append(seedSoil, processMapEntry_P2(numberReg.FindAllString(lines[i], -1)))
			case 2:
				soilFert = append(soilFert, processMapEntry_P2(numberReg.FindAllString(lines[i], -1)))
			case 3:
				fertWater = append(fertWater, processMapEntry_P2(numberReg.FindAllString(lines[i], -1)))
			case 4:
				waterLight = append(waterLight, processMapEntry_P2(numberReg.FindAllString(lines[i], -1)))
			case 5:
				lightTemp = append(lightTemp, processMapEntry_P2(numberReg.FindAllString(lines[i], -1)))
			case 6:
				tempHumid = append(tempHumid, processMapEntry_P2(numberReg.FindAllString(lines[i], -1)))
			case 7:
				humidLoc = append(humidLoc, processMapEntry_P2(numberReg.FindAllString(lines[i], -1)))
			default:
				// fmt.Printf("line: %v\n", lines[i])
			}
		}
	}

	sort.Sort(seeds)
	sort.Sort(seedSoil)
	sort.Sort(soilFert)
	sort.Sort(fertWater)
	sort.Sort(waterLight)
	sort.Sort(lightTemp)
	sort.Sort(tempHumid)
	sort.Sort(humidLoc)

	numRanges := mapNumRanges_P2(seeds, seedSoil)
	sort.Sort(numRanges)
	numRanges = mapNumRanges_P2(numRanges, soilFert)
	sort.Sort(numRanges)
	numRanges = mapNumRanges_P2(numRanges, fertWater)
	sort.Sort(numRanges)
	numRanges = mapNumRanges_P2(numRanges, waterLight)
	sort.Sort(numRanges)
	numRanges = mapNumRanges_P2(numRanges, lightTemp)
	sort.Sort(numRanges)
	numRanges = mapNumRanges_P2(numRanges, tempHumid)
	sort.Sort(numRanges)
	numRanges = mapNumRanges_P2(numRanges, humidLoc)
	sort.Sort(numRanges)

	fmt.Printf("%v\n", numRanges[0].start)
}

func processSeeds_P2(seeds []int) rangeEntries_P2 {

	var out rangeEntries_P2

	for i := 0; i < len(seeds); i += 2 {
		out = append(out, rangeEntry_P2{start: seeds[i], end: seeds[i] + seeds[i+1] - 1})
	}
	return out
}

func mapNumRanges_P2(rangeEntries rangeEntries_P2, mapEntries mapEntries_P2) rangeEntries_P2 {

	var out []rangeEntry_P2

	for _, rEntry := range rangeEntries {
		out = append(out, mapNumRange_P2(rEntry, mapEntries)...)
	}

	return out
}

func mapNumRange_P2(rEntry rangeEntry_P2, mapEntries mapEntries_P2) rangeEntries_P2 {

	var out []rangeEntry_P2

	rMin := rEntry.start
	rMax := rEntry.end

	for _, mEntry := range mapEntries {
		mMin := mEntry.sStart
		mMax := mEntry.sEnd

		if rMin > mMax {
			continue
		} else if rMax < mMin {
			out = append(out, rangeEntry_P2{start: rMin, end: rMax})
			break
		} else if rMin < mMin {
			out = append(out, rangeEntry_P2{start: rMin, end: mMin - 1})
			rMin = mMin
		}

		if rMax <= mMax && rMin >= mMin {
			out = append(out, rangeEntry_P2{start: ((rMin - mMin) + mEntry.dStart), end: ((rMax - mMin) + mEntry.dStart)})
			break
		} else if rMax > mMax {
			out = append(out, rangeEntry_P2{start: ((rMin - mMin) + mEntry.dStart), end: mEntry.dEnd})
			rMin = mMax + 1
			continue
		}
	}
	return out
}

// func mapNum(num int, mapEntries []mapEntry) int {

// 	for _, entry := range mapEntries {
// 		max := entry.source + entry.length - 1
// 		if num <= max && num >= entry.source {
// 			return ((num - entry.source) + entry.destination)
// 		}
// 	}
// 	return num
// }

func processMapEntry_P2(in []string) mapEntry_P2 {

	var out mapEntry_P2
	rawEntry := asti(in)

	if len(rawEntry) < 3 {
		fmt.Printf("in: %v\n", in)
	} else {
		out.dStart = rawEntry[0]
		out.sStart = rawEntry[1]
		out.length = rawEntry[2]
		out.sEnd = out.sStart + out.length - 1
		out.dEnd = out.dStart + out.length - 1
	}

	return out
}

// func asti(in []string) []int {
// 	var out []int

// 	for _, val := range in {
// 		out = append(out, sti(val))
// 	}
// 	return out
// }

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
