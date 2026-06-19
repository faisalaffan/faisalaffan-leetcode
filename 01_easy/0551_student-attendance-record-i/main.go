package main

// LeetCode #551: Student Attendance Record I
// https://leetcode.com/problems/student-attendance-record-i/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(1)
func StudentAttendanceRecordI(s string) bool {
	absences, lateStreak := 0, 0
	for _, c := range s {
		if c == 'A' {
			absences++
			lateStreak = 0
			if absences >= 2 {
				return false
			}
		} else if c == 'L' {
			lateStreak++
			if lateStreak >= 3 {
				return false
			}
		} else {
			lateStreak = 0
		}
	}
	return true
}

func main() {
	fmt.Println(StudentAttendanceRecordI("PPALLP"))
	fmt.Println(StudentAttendanceRecordI("PPALLL"))
}
