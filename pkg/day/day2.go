package day

import (
	"aoc-2024/pkg/utils"
	"fmt"
	"math"
)

type Day2 struct {
	TOLERANCE int
	DEBUG     bool

	reports [][]int
}

func NewDay2(debug bool) (*Day2, error) {

	day2 := &Day2{
		TOLERANCE: 0,
		DEBUG:     debug,
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

	d.TOLERANCE = 0

	countOfSafeReports := 0

	for _, row := range d.reports {
		if d.DetermineLevelSafety(row) {
			countOfSafeReports++
		}

		// for j, value := range row {
		// 	fmt.Printf("reports[%d][%d] = %d\n", i, j, value)
		// }
	}

	println("Count of safe reports: ", countOfSafeReports)

}

func (d *Day2) Part2() {

	d.TOLERANCE = 1

	countOfSafeReports := 0

	for _, row := range d.reports {
		if d.DetermineLevelSafety(row) {
			countOfSafeReports++
		}
	}

	println("Count of safe reports: ", countOfSafeReports)

}

func (d *Day2) DetermineLevelSafety(level []int) bool {

	withinThresholdMagnitude := d.CheckSequentialDifferences(level)

	allIncreasing := d.VerifyAllLevelsAreIncreasing(level)
	allDecreasing := d.VerifyAllLevelsAreDecreasing(level)

	if d.DEBUG {
		fmt.Println()
		fmt.Println(withinThresholdMagnitude, ": ", level)
		fmt.Println(allIncreasing, ": ", level)
		fmt.Println(allDecreasing, ": ", level)
	}

	withinThresholdLinear := allIncreasing || allDecreasing

	return withinThresholdMagnitude && withinThresholdLinear

}

func (d *Day2) CheckSequentialDifferences(list []int) bool {
	tolerance := 0
	for i := 0; i < len(list)-1; i++ {
		if math.Abs(float64(list[i]-list[i+1])) > 3 {
			tolerance++
			if tolerance > d.TOLERANCE {
				return false
			}
		}
	}
	return true
}

func (d *Day2) VerifyAllLevelsAreIncreasing(list []int) bool {

	errorCount := 0
	for i := 0; i < len(list)-1; i++ {

		if list[i] >= list[i+1] {
			errorCount++
			if errorCount > d.TOLERANCE {
				fmt.Println("returning false")
				return false
			}
		}
	}
	return true
}

func (d *Day2) VerifyAllLevelsAreDecreasing(list []int) bool {
	tolerance := 0
	for i := 0; i < len(list)-1; i++ {
		if list[i] <= list[i+1] {
			tolerance++
			if tolerance > d.TOLERANCE {
				return false
			}
		}
	}
	return true
}
