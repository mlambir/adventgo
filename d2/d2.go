package main

import (
	"fmt"
	"strings"

	"mlambir.com/adventgo/utils"
)

func removeCopy(slice []string, i int) []string {
	if i < 0 || i >= len(slice) {
		return slice
	}
	s := make([]string, len(slice))
	copy(s, slice)
	copy(s[i:], s[i+1:])
	return s[:len(s)-1]
}

func checkSafe(line []string) int {
	prev := utils.AtoiP(line[0])
	dir := 0
	for i, n := range line[1:] {
		nn := utils.AtoiP(n)
		diff := nn - prev
		if dir == 0 {
			if diff == 0 {
				return i
			}
			dir = diff / utils.Abs(diff)
		}
		diff *= dir

		if diff > 3 || diff <= 0 {
			return i
		}
		prev = nn
	}
	return -1
}

func process(filename string) int {
	lines := utils.ReadLines(filename)
	safe := 0
	for _, line := range lines {
		splitLine := strings.Fields(line)
		unsafe := checkSafe(splitLine)
		if unsafe == -1 {
			safe += 1
		}
	}
	return safe
}

func process2(filename string) int {
	lines := utils.ReadLines(filename)
	safe := 0
	for _, line := range lines {
		splitLine := strings.Fields(line)
		unsafe := checkSafe(splitLine)
		if unsafe == -1 {
			fmt.Println("safe ", splitLine)
			safe += 1
		} else if checkSafe(removeCopy(splitLine, unsafe)) == -1 {
			fmt.Println("safe removing", unsafe, splitLine)
			safe += 1
		} else if checkSafe(removeCopy(splitLine, unsafe+1)) == -1 {
			fmt.Println("safe removing", unsafe+1, splitLine)
			safe += 1
		} else if checkSafe(removeCopy(splitLine, unsafe-1)) == -1 {
			fmt.Println("safe removing", unsafe-1, splitLine)
			safe += 1
		} else {
			fmt.Println("unsafe ", splitLine, unsafe)
		}
	}
	return safe
}
func main() {
	fmt.Println("example output: ", process2("d2/d2_ex.txt"))
	fmt.Println("real output: ", process2("d2/d2_val.txt"))
}
