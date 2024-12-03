package day

import (
	"testing"
)

func TestDetermineLevelSafetyTolerance1(t *testing.T) {
	d := &Day2{tolerance: 1}

	testCases := []struct {
		input    []int
		expected bool
	}{
		// {[]int{1, 1, 2, 3, 4, 5}, true},
		// {[]int{1, 2, 3, 4, 5, 5}, true},
		// {[]int{5, 1, 2, 3, 4, 5}, true},
		// {[]int{1, 4, 3, 2, 1}, true},
		// {[]int{1, 6, 7, 8, 9}, true},
		// {[]int{1, 2, 3, 4, 3}, true},
		// {[]int{9, 8, 7, 6, 7}, true},
		// {[]int{7, 10, 8, 10, 11}, true},
		// {[]int{29, 28, 27, 25, 26, 25, 22, 20}, true},
		// {[]int{48, 46, 47, 49, 51, 54, 56}, true},
		// {[]int{8, 9, 10, 11}, true},
		// {[]int{1, 1, 1, 1}, false},
		// {[]int{31, 34, 32, 30, 28, 27, 24, 22}, true},
		// {[]int{2, 6, 6, 8, 6}, false},
		// {[]int{33, 29, 27, 24, 21, 19}, true},
		// //11 8 9 10 12
		// {[]int{11, 8, 9, 10, 12}, true},
		// //53 49 47 44 42 40 38 35
		// {[]int{53, 49, 47, 44, 42, 40, 38, 35}, true},
		// //52 51 48 47 41
		// {[]int{52, 51, 48, 47, 41}, true},
		// //1 3 4 5 8 10 7
		{[]int{73, 75, 77, 80, 81, 78, 81, 82}, false},
	}

	for _, tc := range testCases {
		result := d.SafeLevelWithTolerance(tc.input)
		t.Logf("Input: %v, Result: %v, Expected: %v", tc.input, result, tc.expected)
		if result != tc.expected {
			t.Errorf("DetermineLevelSafety(%v) = %v; expected %v", tc.input, result, tc.expected)
		}
	}
}
