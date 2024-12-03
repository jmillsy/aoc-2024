package day

import (
	"aoc-2024/pkg/utils"
	"fmt"
	"regexp"
	"strconv"
)

type Day3 struct {
	corruptedMemory string
}

type Token struct {
	Type  string
	Value string
}

func NewDay3(debug bool) (*Day3, error) {

	day := &Day3{}

	inputFileName := "inputs/day3.txt"
	if debug {
		inputFileName = "inputs/day3-debug.txt"
	}

	corruptedMemory, err := utils.ReadFileToString(inputFileName)
	if err != nil {
		return nil, err
	}

	day.corruptedMemory = corruptedMemory

	return day, nil

}

func (d *Day3) Part1() {

	fmt.Println("====================== DAY 3 PART 1 ======================")

	mulSections := parseMulSections(d.corruptedMemory)
	//fmt.Println("Parsed mul sections:", mulSections)
	result := 0
	for _, section := range mulSections {
		product := parseMulAndReturnProduct(section)
		//fmt.Println("Product of section:", product)
		result += product
	}

	fmt.Println("Result:", result)

}

func (d *Day3) Part2() {

	fmt.Println("====================== DAY 3 PART 2 ======================")

	tokens := tokenizeInput(d.corruptedMemory)
	enabled := true
	result := 0
	for _, token := range tokens {

		if token.Type == "do" {
			enabled = true
		} else if token.Type == "don't" {
			enabled = false
		} else if enabled && token.Type == "mul" {
			product := parseMulAndReturnProduct(token.Value)
			fmt.Println("Parsing mul section:", token.Value, "Product:", product)
			result += product
		} else if !enabled && token.Type == "mul" {
			fmt.Println("Skipping mul section:", token.Value)
		}

	}

	fmt.Println("Result:", result)

}

func parseMulSections(input string) []string {
	re := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	return re.FindAllString(input, -1)
}

func parseMulAndReturnProduct(input string) int {

	//input will be mul(X,Y), i want to retur nX*Y
	re := regexp.MustCompile(`\d{1,3}`)
	numbers := re.FindAllString(input, -1)
	x, _ := strconv.Atoi(numbers[0])
	y, _ := strconv.Atoi(numbers[1])
	return x * y

}

func tokenizeInput(input string) []Token {
	var tokens []Token

	reMul := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	reDo := regexp.MustCompile(`do\(\)`)
	reDont := regexp.MustCompile(`don't\(\)`)

	for len(input) > 0 {
		if loc := reMul.FindStringIndex(input); loc != nil && loc[0] == 0 {
			tokens = append(tokens, Token{Type: "mul", Value: input[loc[0]:loc[1]]})
			input = input[loc[1]:]
		} else if loc := reDo.FindStringIndex(input); loc != nil && loc[0] == 0 {
			tokens = append(tokens, Token{Type: "do", Value: input[loc[0]:loc[1]]})
			input = input[loc[1]:]
		} else if loc := reDont.FindStringIndex(input); loc != nil && loc[0] == 0 {
			tokens = append(tokens, Token{Type: "don't", Value: input[loc[0]:loc[1]]})
			input = input[loc[1]:]
		} else {
			tokens = append(tokens, Token{Type: "garbage", Value: string(input[0])})
			input = input[1:]
		}
	}

	return tokens
}
