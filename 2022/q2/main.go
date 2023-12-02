package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Hello")

	instr := read("input.txt")

	lines := strings.Split(instr, "\r\n")

	curSum := 0
	curElf := 1

	elf1 := 0
	elf2 := 0
	elf3 := 0

	for i := range lines {
		if lines[i] == "" || i == len(lines)-1 {
			if curSum > elf1 {
				elf3 = elf2
				elf2 = elf1
				elf1 = curSum
			} else if curSum > elf2{
				elf3 = elf2
				elf2 = curSum
			} else if curSum > elf3{
				elf3 = curSum
			}
			curElf++
			curSum = 0
		} else {
			num := sti(lines[i])
			curSum += num
		}
	}

	sum := elf1+elf2+elf3
	fmt.Printf("%d", sum)
}



func sti(in string) int {
	out, err := strconv.Atoi(in)
	if err != nil {
		fmt.Printf("ERROR: %s\n", in)
		panic(err)
	}
	return out
}

func read(file string) string {
	inbyte, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("Err")
	}
	return string(inbyte)
}