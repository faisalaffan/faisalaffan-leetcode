package main

// LeetCode #175: Combine Two Tables
// https://leetcode.com/problems/combine-two-tables/
// Difficulty: Easy

import "fmt"

// Time: O(n) | Space: O(1)
// Write your MySQL query statement below
func CombineTwoTables() string {
	return "SELECT Person.firstName, Person.lastName, Address.city, Address.state FROM Person LEFT JOIN Address ON Person.personId = Address.personId"
}

func main() {
	fmt.Println(CombineTwoTables())
}
