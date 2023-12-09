package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type mapEntry struct {
	sStart int
	sEnd   int
	dStart int
	dEnd   int
	length int
}

type rangeEntry struct {
	start int
	end   int
}

type rangeEntries []rangeEntry

func (m rangeEntries) Len() int           { return len(m) }
func (m rangeEntries) Less(i, j int) bool { return m[i].start < m[j].start }
func (m rangeEntries) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }

type mapEntries []mapEntry

func (m mapEntries) Len() int           { return len(m) }
func (m mapEntries) Less(i, j int) bool { return m[i].sStart < m[j].sStart }
func (m mapEntries) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }

func main() {

	var seeds rangeEntries
	var seedSoil mapEntries
	var soilFert mapEntries
	var fertWater mapEntries
	var waterLight mapEntries
	var lightTemp mapEntries
	var tempHumid mapEntries
	var humidLoc mapEntries

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
				seeds = processSeeds(asti(numberReg.FindAllString(subLine[1], -1)))
			}
			// fmt.Printf("%v:%v\n", entryNum, lines[i])
		} else if lines[i] != "" {
			switch entryNum {
			case 1:
				seedSoil = append(seedSoil, processMapEntry(numberReg.FindAllString(lines[i], -1)))
			case 2:
				soilFert = append(soilFert, processMapEntry(numberReg.FindAllString(lines[i], -1)))
			case 3:
				fertWater = append(fertWater, processMapEntry(numberReg.FindAllString(lines[i], -1)))
			case 4:
				waterLight = append(waterLight, processMapEntry(numberReg.FindAllString(lines[i], -1)))
			case 5:
				lightTemp = append(lightTemp, processMapEntry(numberReg.FindAllString(lines[i], -1)))
			case 6:
				tempHumid = append(tempHumid, processMapEntry(numberReg.FindAllString(lines[i], -1)))
			case 7:
				humidLoc = append(humidLoc, processMapEntry(numberReg.FindAllString(lines[i], -1)))
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

	numRanges := mapNumRanges(seeds, seedSoil)
	sort.Sort(numRanges)
	numRanges = mapNumRanges(numRanges, soilFert)
	sort.Sort(numRanges)
	numRanges = mapNumRanges(numRanges, fertWater)
	sort.Sort(numRanges)
	numRanges = mapNumRanges(numRanges, waterLight)
	sort.Sort(numRanges)
	numRanges = mapNumRanges(numRanges, lightTemp)
	sort.Sort(numRanges)
	numRanges = mapNumRanges(numRanges, tempHumid)
	sort.Sort(numRanges)
	numRanges = mapNumRanges(numRanges, humidLoc)
	sort.Sort(numRanges)

	fmt.Printf("%v\n", numRanges[0].start)
}

func processSeeds(seeds []int) rangeEntries {

	var out rangeEntries

	for i := 0; i < len(seeds); i += 2 {
		out = append(out, rangeEntry{start: seeds[i], end: seeds[i] + seeds[i+1] - 1})
	}
	return out
}

func mapNumRanges(rangeEntries rangeEntries, mapEntries mapEntries) rangeEntries {

	var out []rangeEntry

	for _, rEntry := range rangeEntries {
		out = append(out, mapNumRange(rEntry, mapEntries)...)
	}

	return out
}

func mapNumRange(rEntry rangeEntry, mapEntries mapEntries) rangeEntries {

	var out []rangeEntry

	rMin := rEntry.start
	rMax := rEntry.end

	for _, mEntry := range mapEntries {
		mMin := mEntry.sStart
		mMax := mEntry.sEnd

		// Handle resulting ranges
		if rMin > mMax {
			// All higher
			continue
		} else if rMax < mMin {
			// All lower
			out = append(out, rangeEntry{start: rMin, end: rMax})
			break
		} else if rMin < mMin {
			// Lower Part
			out = append(out, rangeEntry{start: rMin, end: mMin - 1})
			rMin = mMin
		}

		if rMax <= mMax {
			// All Inside
			out = append(out, rangeEntry{start: ((rMin - mMin) + mEntry.dStart), end: ((rMax - mMin) + mEntry.dStart)})
			break
		} else if rMax > mMax {
			// Inside and higher part
			out = append(out, rangeEntry{start: ((rMin - mMin) + mEntry.dStart), end: mEntry.dEnd})
			rMin = mMax + 1
			continue
		}
	}
	return out
}

func processMapEntry(in []string) mapEntry {

	var out mapEntry
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

func asti(in []string) []int {
	var out []int

	for _, val := range in {
		out = append(out, sti(val))
	}
	return out
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
