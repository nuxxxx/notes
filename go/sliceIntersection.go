package main

import "fmt"

/**
Func that creates the intersection of 2 slices
*/

func intersection[T comparable](a, b []T) []T {
	res := []T{}
	m := map[T]int{}

	for _, v := range a {
		m[v]++
	}
	
	for _, v := range b {
		for i := m[v]; i > 0; i-- {
			res = append(res, v)
		}
		m[v] = -1;
	}

	return res
}

func testIntersection() {
	a := []int{23, 3, 1, 2}
	b := []int{6, 2, 4, 23}
	// [2, 23]
	fmt.Printf("%v\n", intersection(a, b))
	a = []int{1, 1, 1}
	b = []int{1, 1, 1, 1}
	// [1, 1, 1]
	fmt.Printf("%v\n", intersection(a, b))
} 