package main

/*
Notes - This algorithm performs in O(N) time: there's one O(N) iteration of the
array and a hash table addition and lookup, which should usually be O(1) but may
be O(N) in the worst case scenario, according to the implementation. Hash table
space complexity has a worst case complexity of O(N). It seems as though the Go
spec does not provide a guarantee on the run time complexity of the map type, so
those are typical assumptions about performance.

The algorithm works by setting every positive integer in the array as a key in a
map, and then checking to see how many consecutive positive numbers, starting at
1, are found in the map. The first number not found is the returned value.
*/

func solution(A []int) int {

	positive_numbers := make(map[int]int)
	smallest_number := 1

	for _, number := range A {
		if number > 0 {
			positive_numbers[number] = 0
		}
	}

	for {
		_, exists := positive_numbers[smallest_number]

		if !exists {
			return smallest_number
		}
		smallest_number++
	}
}
