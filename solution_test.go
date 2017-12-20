package main

import "testing"

func TestSolution(t *testing.T) {

	test_data := [][]int {
		{1, 3, 6, 4, 1, 2},
		{1, 2, 3},
		{-1, -3},
	}
	test_results := []int {
		5,
		4,
		1,
	}

	for i, test_input_array := range test_data {
		if solution(test_input_array) != test_results[i] {
			t.Errorf("Error on test number %d", i)
		}
	}
}
