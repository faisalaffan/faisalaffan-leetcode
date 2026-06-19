package main

// LeetCode #1333: Filter Restaurants by Vegan-Friendly, Price and Distance
// https://leetcode.com/problems/filter-restaurants-by-vegan-friendly-price-and-distance/
// Difficulty: Medium

import "fmt"
import "sort"

func main() {
	// Test case 1
	fmt.Println(filterRestaurants([][]int{
		{1, 4, 1, 40, 10},
		{2, 8, 0, 50, 5},
		{3, 8, 1, 30, 4},
		{4, 10, 0, 10, 3},
		{5, 1, 1, 15, 1},
	}, 1, 50, 10))
	// [3,1,5]

	// Test case 2
	fmt.Println(filterRestaurants([][]int{
		{1, 4, 1, 40, 10},
		{2, 8, 0, 50, 5},
		{3, 8, 1, 30, 4},
		{4, 10, 0, 10, 3},
		{5, 1, 1, 15, 1},
	}, 0, 50, 10))
	// [4,3,2,1,5]

	// Test case 3
	fmt.Println(filterRestaurants([][]int{
		{1, 4, 1, 40, 10},
	}, 1, 30, 10))
	// []
}

type restaurant struct {
	id, rating, veganFriendly, price, distance int
}

// Time: O(n log n) for sorting
// Space: O(n) for storing filtered results
func filterRestaurants(restaurants [][]int, veganFriendly int, maxPrice int, maxDistance int) []int {
	var filtered []restaurant
	for _, r := range restaurants {
		if veganFriendly == 1 && r[2] != 1 {
			continue
		}
		if r[3] > maxPrice {
			continue
		}
		if r[4] > maxDistance {
			continue
		}
		filtered = append(filtered, restaurant{r[0], r[1], r[2], r[3], r[4]})
	}

	sort.Slice(filtered, func(i, j int) bool {
		if filtered[i].rating != filtered[j].rating {
			return filtered[i].rating > filtered[j].rating
		}
		return filtered[i].id > filtered[j].id
	})

	result := make([]int, len(filtered))
	for i, r := range filtered {
		result[i] = r.id
	}
	return result
}
