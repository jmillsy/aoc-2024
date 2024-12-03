package day

import (
	"aoc-2024/pkg/utils"
)

type Day2 struct {
	tolerance int
	debug     bool

	reports [][]int
}

func NewDay2(debug bool) (*Day2, error) {

	day2 := &Day2{
		tolerance: 0,
		debug:     debug,
	}

	inputFileName := "inputs/day2.txt"
	if debug {
		inputFileName = "inputs/day2-debug.txt"
	}

	reports, err := utils.ReadMatrixFromFile(inputFileName)
	if err != nil {
		return nil, err
	}

	day2.reports = reports
	return day2, nil

}

func (d *Day2) Part1() {

	d.tolerance = 0
	countOfSafeReports := 0

	for _, row := range d.reports {
		if d.SafeLevel(row) {
			countOfSafeReports++
		}
	}

	println("Count of safe reports: ", countOfSafeReports)

}

func (d *Day2) Part2() {

	countOfSafeReports := 0

	for _, row := range d.reports {
		if d.SafeLevelWithTolerance(row) {
			countOfSafeReports++
		}
	}

	println("Count of safe reports: ", countOfSafeReports)

}

func (d *Day2) SafeLevelWithTolerance(level []int) bool {

	isSafe := false
	isSafe = d.SafeLevel(level)

	if isSafe {
		return true
	} else {
		// Remove the baddie and check again
		for i := 0; i < len(level); i++ {
			modifiedLevels := []int{}
			modifiedLevels = append(modifiedLevels, level[:i]...)
			modifiedLevels = append(modifiedLevels, level[i+1:]...)
			isSafe = d.SafeLevel(modifiedLevels)

			if isSafe {
				return true
			}
		}
	}

	return false

}

func (d *Day2) SafeLevel(levels []int) bool {
	isIncreasing := true
	isSafe := true
	badLevel := 0

	i := 0
	for i = 1; i < len(levels); i++ {
		diff := levels[i] - levels[i-1]

		if diff == 0 {
			badLevel++
			isSafe = false
			break
		}

		if diff < -3 || diff > 3 {
			isSafe = false
			break
		}

		if i == 1 && diff < 0 {
			isIncreasing = false
		} else if (isIncreasing && diff < 0) || (!isIncreasing && diff > 0) {
			isSafe = false
			break
		}
	}

	return isSafe
}
