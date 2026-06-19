package main

// LeetCode #1350: Students With Invalid Departments
// https://leetcode.com/problems/students-with-invalid-departments/
// Difficulty: Easy [Paid]
//
// This is a SQL problem. The solution is the SQL query below.
// Tables: Students (id, name, department_id), Departments (id, name)

import "fmt"

func main() {
	fmt.Println(StudentsWithInvalidDepartments())
}

// Time: N/A (SQL query), Space: N/A
func StudentsWithInvalidDepartments() string {
	return `SELECT s.id, s.name
FROM Students s
LEFT JOIN Departments d ON s.department_id = d.id
WHERE d.id IS NULL;`
}
