package main

import (
	"aoc-2024/pkg/day"
)

func main() {

	var d day.Day

	day1, err := day.NewDay1(false)
	if err != nil {
		panic(err)
	}
	d = day1

	d.Part1()
	d.Part2()
}
