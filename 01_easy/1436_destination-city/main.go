package main

// LeetCode #1436: Destination City
// https://leetcode.com/problems/destination-city/
// Difficulty: Easy
//
// LeetCode submission: func destCity(paths [][]string) string

import "fmt"

func main() {
	paths1 := [][]string{{"London", "New York"}, {"New York", "Lima"}, {"Lima", "Sao Paulo"}}
	fmt.Println(DestinationCity(paths1)) // "Sao Paulo"

	paths2 := [][]string{{"B", "C"}, {"D", "B"}, {"C", "A"}}
	fmt.Println(DestinationCity(paths2)) // "A"
}

// Time: O(n), Space: O(n)
func DestinationCity(paths [][]string) string {
	outgoing := make(map[string]bool, len(paths))
	for _, p := range paths {
		outgoing[p[0]] = true
	}
	for _, p := range paths {
		if !outgoing[p[1]] {
			return p[1]
		}
	}
	return ""
}
