package main

// LeetCode #605: Can Place Flowers
// https://leetcode.com/problems/can-place-flowers/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(1)
func CanPlaceFlowers(flowerbed []int, n int) bool {
	count := 0
	for i := 0; i < len(flowerbed) && count < n; i++ {
		if flowerbed[i] == 0 &&
			(i == 0 || flowerbed[i-1] == 0) &&
			(i == len(flowerbed)-1 || flowerbed[i+1] == 0) {
			flowerbed[i] = 1
			count++
		}
	}
	return count >= n
}

func main() {
	fmt.Println(CanPlaceFlowers([]int{1, 0, 0, 0, 1}, 1))
	fmt.Println(CanPlaceFlowers([]int{1, 0, 0, 0, 1}, 2))
	fmt.Println(CanPlaceFlowers([]int{0, 0, 1, 0, 0}, 1))
}
