package main

// LeetCode #2292: Products With Three or More Orders in Two Consecutive Years
// https://leetcode.com/problems/products-with-three-or-more-orders-in-two-consecutive-years/
// Difficulty: Medium [Paid]
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

type Order struct {
	ProductID int
	Year      int
}

func findProducts(orders [][]int) []int {
	// Group orders by product and year
	productYears := make(map[int]map[int]int)
	for _, o := range orders {
		productID, year := o[0], o[1]
		if productYears[productID] == nil {
			productYears[productID] = make(map[int]int)
		}
		productYears[productID][year]++
	}

	result := []int{}
	for pid, years := range productYears {
		yearList := []int{}
		for y := range years {
			yearList = append(yearList, y)
		}
		sort.Ints(yearList)

		for i := 1; i < len(yearList); i++ {
			if yearList[i]-yearList[i-1] == 1 &&
				years[yearList[i]] >= 3 &&
				years[yearList[i-1]] >= 3 {
				result = append(result, pid)
				break
			}
		}
	}
	sort.Ints(result)
	return result
}

func main() {
	// Test case 1
	fmt.Println(findProducts([][]int{{1, 2020}, {1, 2020}, {1, 2020}, {1, 2021}, {1, 2021}, {1, 2021}, {2, 2020}}))
	// Expected: [1]

	// Test case 2
	fmt.Println(findProducts([][]int{{1, 2020}, {2, 2020}}))
	// Expected: []
}
