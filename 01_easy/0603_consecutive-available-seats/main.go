package main

// LeetCode #603: Consecutive Available Seats
// https://leetcode.com/problems/consecutive-available-seats/
// Difficulty: Easy [Paid]

import "fmt"

func ConsecutiveAvailableSeats() string {
	return "SELECT DISTINCT c1.seat_id FROM Cinema c1 JOIN Cinema c2 ON ABS(c1.seat_id - c2.seat_id) = 1 AND c1.free = 1 AND c2.free = 1 ORDER BY c1.seat_id"
}

func main() {
	fmt.Println(ConsecutiveAvailableSeats())
}
