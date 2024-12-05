package main

import (
	"cmp"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"mlambir.com/adventgo/utils"
)

type location struct {
	val int
	ord int
}

func process(filename string) int {
	fmt.Println("Processing", filename)
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Println("Length:", len(data))
	lines := strings.Split(string(data), "\n")
	fmt.Println("Lines:", len(lines))

	l1 := make([]location, 0, len(lines))
	l2 := make([]location, 0, len(lines))

	for i, line := range lines {
		vals := strings.Fields(line)
		if len(vals) != 2 {
			continue
		}
		v1, err := strconv.Atoi(vals[0])
		if err != nil {
			panic(err)
		}

		v2, err := strconv.Atoi(vals[1])
		if err != nil {
			panic(err)
		}

		l1 = append(l1, location{val: v1, ord: i})
		l2 = append(l2, location{val: v2, ord: i})
	}
	sortf := func(a, b location) int {
		return cmp.Compare(a.val, b.val)
	}
	slices.SortFunc(l1, sortf)
	slices.SortFunc(l2, sortf)
	diff := 0
	for i := 0; i < len(l1); i++ {
		diff += utils.Abs(l1[i].val - l2[i].val)
	}
	return diff
}

func process2(filename string) int {
	fmt.Println("Processing", filename)
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(data), "\n")
	l1 := make([]int, 0, len(lines))
	l2 := make(map[int]int, len(lines))

	for _, line := range lines {
		vals := strings.Fields(line)
		if len(vals) != 2 {
			continue
		}
		v1, err := strconv.Atoi(vals[0])
		if err != nil {
			panic(err)
		}

		v2, err := strconv.Atoi(vals[1])
		if err != nil {
			panic(err)
		}

		l1 = append(l1, v1)
		l2[v2] += 1
	}
	fmt.Println("l1", l1)
	fmt.Println("l2", l2)

	similarity := 0
	for _, n := range l1 {
		similarity += n * l2[n]
	}
	return similarity
}

func main() {
	fmt.Println("example value", process2("d1/d1_ex.txt"))
	fmt.Println("final value", process2("d1/d1_val.txt"))
}
