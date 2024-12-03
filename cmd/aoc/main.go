package main

import (
	"aoc-2024/pkg/day"
)

func main() {

	var d day.Day

	day, err := day.NewDay3(false)
	if err != nil {
		panic(err)
	}
	d = day

	d.Part1()
	d.Part2()
}
