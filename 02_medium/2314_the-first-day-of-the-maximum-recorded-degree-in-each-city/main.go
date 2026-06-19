package main

// LeetCode #2314: The First Day of the Maximum Recorded Degree in Each City
// https://leetcode.com/problems/the-first-day-of-the-maximum-recorded-degree-in-each-city/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(n)

import (
	"fmt"
	"sort"
)

type DayDegree struct {
	CityID int
	Day    int
	Degree int
}

func firstDayOfMaxDegree(data [][]int) []int {
	// data[i] = [city_id, day, degree]
	cityMax := make(map[int]int)
	cityFirstDay := make(map[int]int)

	for _, d := range data {
		cityID, day, degree := d[0], d[1], d[2]
		if prev, ok := cityMax[cityID]; !ok || degree > prev {
			cityMax[cityID] = degree
			cityFirstDay[cityID] = day
		} else if degree == prev && day < cityFirstDay[cityID] {
			cityFirstDay[cityID] = day
		}
	}

	cities := []int{}
	for c := range cityMax {
		cities = append(cities, c)
	}
	sort.Ints(cities)

	result := make([]int, len(cities))
	for i, c := range cities {
		result[i] = cityFirstDay[c]
	}
	return result
}

func main() {
	// Test case 1
	fmt.Println(firstDayOfMaxDegree([][]int{{1, 1, 30}, {1, 2, 35}, {1, 3, 35}, {2, 1, 25}, {2, 2, 30}}))
	// Expected: [2, 2]

	// Test case 2
	fmt.Println(firstDayOfMaxDegree([][]int{{1, 1, 30}, {2, 2, 25}}))
	// Expected: [1, 2]
}
