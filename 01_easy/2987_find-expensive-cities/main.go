package main

// LeetCode #2987: Find Expensive Cities
// https://leetcode.com/problems/find-expensive-cities/
// Difficulty: Easy [Paid - Database]
//
// Note: This is a SQL problem on LeetCode. In Go, we implement the
// equivalent logic: find cities with avg price > overall avg price.

import (
	"fmt"
	"sort"
)

func main() {
	// LeetCode name: findExpensiveCities
	// Input: list of [city, price] pairs
	fmt.Println(FindExpensiveCities([][]string{
		{"New York", "1000"},
		{"New York", "1200"},
		{"Los Angeles", "800"},
		{"Los Angeles", "600"},
		{"Chicago", "500"},
	})) // [New York]

	fmt.Println(FindExpensiveCities([][]string{
		{"A", "200"},
		{"B", "100"},
		{"B", "100"},
	})) // [A B]
}

// Time: O(n log n) | Space: O(n)
// LeetCode submission name: findExpensiveCities
func FindExpensiveCities(listings [][]string) []string {
	// Calculate overall average
	var totalPrice float64
	for _, row := range listings {
		price := parsePrice(row[1])
		totalPrice += price
	}
	avgPrice := totalPrice / float64(len(listings))

	// Calculate per-city average
	citySum := make(map[string]float64)
	cityCount := make(map[string]int)
	for _, row := range listings {
		city := row[0]
		price := parsePrice(row[1])
		citySum[city] += price
		cityCount[city]++
	}

	result := []string{}
	for city, sum := range citySum {
		if sum/float64(cityCount[city]) > avgPrice {
			result = append(result, city)
		}
	}
	sort.Strings(result)
	return result
}

func parsePrice(s string) float64 {
	// Simple string to float conversion
	var result float64
	var neg bool
	for i := 0; i < len(s); i++ {
		if s[i] == '-' {
			neg = true
		} else if s[i] >= '0' && s[i] <= '9' {
			result = result*10 + float64(s[i]-'0')
		} else if s[i] == '.' {
			// Decimal part
			div := 10.0
			for j := i + 1; j < len(s); j++ {
				if s[j] >= '0' && s[j] <= '9' {
					result += float64(s[j]-'0') / div
					div *= 10
				}
			}
			break
		}
	}
	if neg {
		return -result
	}
	return result
}
