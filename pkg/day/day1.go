package day

import (
	"aoc-2024/pkg/utils"
	"fmt"
)

type Day1 struct {
	list1 []int
	list2 []int
}

func NewDay1(debug bool) (*Day1, error) {

	day1 := &Day1{}

	list1, err := utils.ReadColumnFromFile("day1/input.txt", 0)
	if err != nil {
		return nil, err
	}

	list2, err := utils.ReadColumnFromFile("day1/input.txt", 1)
	if err != nil {
		return nil, err
	}

	day1.list1 = utils.SortList(list1)
	day1.list2 = utils.SortList(list2)

	return day1, nil

}

func (d *Day1) Part1() {

	deltaList, err := utils.CalculateDelta(d.list1, d.list2)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Delta List:", deltaList)

	sum := utils.SumList(deltaList)

	fmt.Println("Sum of Delta List:", sum)

}

func (d *Day1) Part2() {

	similarityScore := 0
	for i := 0; i < len(d.list1); i++ {
		count := utils.CountOccurrences(d.list2, d.list1[i])
		//fmt.Printf("Number %d occurs %d times in list2\n", d.list1[i], count)
		similarityScore += d.list1[i] * count
	}
	fmt.Println("Similarity Score:", similarityScore)

}
