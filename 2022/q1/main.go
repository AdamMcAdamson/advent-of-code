package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Hello")

	inbyte, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		fmt.Println("Err")
	}
	instr := string(inbyte)

	lines := strings.Split(instr, "\r\n")

	curSum := 0
	curElf := 1
	maxSum := 0

	for i := range lines {
		if lines[i] == "" || i == len(lines)-1 {
			if curSum > maxSum {
				maxSum = curSum
			}
			curElf++
			curSum = 0
		} else {
			num, err := strconv.Atoi(lines[i])
			if err != nil {
				fmt.Printf("ERROR: %s\n", lines[i])
				panic(err)
			}
			curSum += num
		}
	}
	fmt.Printf("%d", maxSum)
}
