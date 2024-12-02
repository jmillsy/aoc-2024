package utils

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func SortList(list []int) []int {
	sort.Ints(list)
	return list
}

func CalculateDelta(list1, list2 []int) ([]int, error) {
	if len(list1) != len(list2) {
		return nil, fmt.Errorf("lists must be of the same length")
	}

	delta := make([]int, len(list1))
	for i := 0; i < len(list1); i++ {
		delta[i] = int(math.Abs(float64(list1[i] - list2[i])))
	}
	return delta, nil
}

func SumList(list []int) int {
	sum := 0
	for _, v := range list {
		sum += v
	}
	return sum
}

// ReadColumnFromFile reads a specific column from a file and returns it as a slice of integers.
func ReadColumnFromFile(filePath string, columnIndex int) ([]int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var list []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		columns := strings.Fields(line)
		if len(columns) > columnIndex {
			value, err := strconv.Atoi(columns[columnIndex])
			if err != nil {
				return nil, err
			}
			list = append(list, value)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return list, nil
}

func CountOccurrences(list []int, number int) int {
	count := 0
	for _, v := range list {
		if v == number {
			count++
		}
	}
	return count
}

func ReadMatrixFromFile(filePath string) ([][]int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var matrix [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		columns := strings.Fields(line)
		var row []int
		for _, col := range columns {
			value, err := strconv.Atoi(col)
			if err != nil {
				return nil, err
			}
			row = append(row, value)
		}
		matrix = append(matrix, row)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return matrix, nil
}
