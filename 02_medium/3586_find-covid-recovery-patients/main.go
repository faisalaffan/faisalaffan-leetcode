package main

// LeetCode #3586: Find COVID Recovery Patients
// https://leetcode.com/problems/find-covid-recovery-patients/
// Difficulty: Medium
// Complexity: O(n) time, O(n) space

import "fmt"

func main() {
	// Test case 1
	records := []int{1, 0, 1, 0, 0, 1}
	fmt.Println("Test 1:", FindCovidRecoveryPatients(records))
	// Test case 2
	records2 := []int{1, 1, 1}
	fmt.Println("Test 2:", FindCovidRecoveryPatients(records2))
	// Test case 3
	records3 := []int{0, 0, 0}
	fmt.Println("Test 3:", FindCovidRecoveryPatients(records3))
}

func FindCovidRecoveryPatients(records []int) int {
	// Count patients who have recovered (positive followed by negative)
	recovered := 0
	for i := 1; i < len(records); i++ {
		if records[i-1] == 1 && records[i] == 0 {
			recovered++
		}
	}
	return recovered
}
