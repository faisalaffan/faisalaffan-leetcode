package main

// LeetCode #577: Employee Bonus
// https://leetcode.com/problems/employee-bonus/
// Difficulty: Easy

import "fmt"

func EmployeeBonus() string {
	return "SELECT e.name, b.bonus FROM Employee e LEFT JOIN Bonus b ON e.empId = b.empId WHERE b.bonus < 1000 OR b.bonus IS NULL"
}

func main() {
	fmt.Println(EmployeeBonus())
}
