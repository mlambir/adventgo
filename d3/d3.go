package main

import (
	"fmt"
	"os"
	"regexp"

	"mlambir.com/adventgo/utils"
)

func process(filename string) int {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	r, _ := regexp.Compile("mul\\((\\d+),(\\d+)\\)")
	fmt.Println(string(data))
	found := r.FindAllStringSubmatch(string(data), -1)
	fmt.Println(found)
	total := 0
	for _, v := range found {
		a := utils.AtoiP(v[1])
		b := utils.AtoiP(v[2])
		total += a * b
	}

	return total
}

func process2(filename string) int {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	r, _ := regexp.Compile("(?:do\\(\\))|(?:don't\\(\\))|(?:mul\\((\\d+),(\\d+)\\))")
	fmt.Println(string(data))
	found := r.FindAllStringSubmatch(string(data), -1)
	fmt.Println(found)
	total := 0
	enabled := true
	for _, v := range found {
		if v[0] == "do()" {
			enabled = true
		} else if v[0] == "don't()" {
			enabled = false
		} else if enabled {
			a := utils.AtoiP(v[1])
			b := utils.AtoiP(v[2])
			total += a * b
		}
	}

	return total
}
func main() {
	fmt.Println(process2("d3/d3_ex.txt"))
	fmt.Println(process2("d3/d3_val.txt"))
}
