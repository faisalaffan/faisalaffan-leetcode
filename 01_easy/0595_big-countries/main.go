package main

// LeetCode #595: Big Countries
// https://leetcode.com/problems/big-countries/
// Difficulty: Easy

import "fmt"

func BigCountries() string {
	return "SELECT name, population, area FROM World WHERE area >= 3000000 OR population >= 25000000"
}

func main() {
	fmt.Println(BigCountries())
}
