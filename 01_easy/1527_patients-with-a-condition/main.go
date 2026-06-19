package main

// LeetCode #1527: Patients With a Condition
// https://leetcode.com/problems/patients-with-a-condition/
// Difficulty: Easy
//
// This is a SQL problem. The solution is the SQL query below.
// Table: Patients (patient_id, patient_name, conditions)

import "fmt"

func main() {
	fmt.Println(PatientsWithACondition())
}

// Time: N/A (SQL query), Space: N/A
func PatientsWithACondition() string {
	return `SELECT patient_id, patient_name, conditions
FROM Patients
WHERE conditions LIKE 'DIAB1%'
   OR conditions LIKE '% DIAB1%';`
}
