package main

// LeetCode #619: Biggest Single Number
// https://leetcode.com/problems/biggest-single-number/
// Difficulty: Easy

import "fmt"

func BiggestSingleNumber() string {
	return "SELECT MAX(num) AS num FROM (SELECT num FROM MyNumbers GROUP BY num HAVING COUNT(num) = 1) AS single_numbers"
}

func main() {
	fmt.Println(BiggestSingleNumber())
}
