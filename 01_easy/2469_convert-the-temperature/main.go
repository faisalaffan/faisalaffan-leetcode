package main

// LeetCode #2469: Convert the Temperature
// https://leetcode.com/problems/convert-the-temperature/
// Difficulty: Easy
// Time O(1) | Space O(1)

import "fmt"

func main() {
	fmt.Println(ConvertTheTemperature(36.50)) // [309.65, 97.7]
	fmt.Println(ConvertTheTemperature(122.11)) // [395.26, 251.798]
}

func ConvertTheTemperature(celsius float64) []float64 {
	return []float64{celsius + 273.15, celsius*1.8 + 32.0}
}
