package main

// LeetCode #2153: The Number of Passengers in Each Bus II
// https://leetcode.com/problems/the-number-of-passengers-in-each-bus-ii/
// Difficulty: Hard [Paid]
//
// Sort buses and passengers by arrival time. For each bus in order, board
// waiting passengers up to capacity. Each bus departs at its arrival time;
// passengers arriving at the exact same time can board.

import (
	"fmt"
	"sort"
)

type bus struct {
	id    int
	time  int
	cap   int
}

type passenger struct {
	time int
	id   int
}

func main() {
	// Example 1
	buses1 := []bus{{1, 2, 2}, {2, 5, 3}}
	passengers1 := []passenger{{1, 101}, {2, 102}, {3, 103}, {4, 104}, {5, 105}}
	fmt.Println(calculateBusPassengers(buses1, passengers1))

	// Example 2
	buses2 := []bus{{1, 3, 2}, {2, 5, 2}}
	passengers2 := []passenger{{1, 201}, {2, 202}, {5, 203}}
	fmt.Println(calculateBusPassengers(buses2, passengers2))
}

func calculateBusPassengers(buses []bus, passengers []passenger) map[int]int {
	// Sort buses by arrival time
	sort.Slice(buses, func(i, j int) bool { return buses[i].time < buses[j].time })

	// Sort passengers by arrival time
	sort.Slice(passengers, func(i, j int) bool { return passengers[i].time < passengers[j].time })

	result := make(map[int]int)
	passIdx := 0
	n := len(passengers)

	for _, b := range buses {
		boarded := 0
		for boarded < b.cap && passIdx < n && passengers[passIdx].time <= b.time {
			result[b.id]++
			boarded++
			passIdx++
		}
		if result[b.id] == 0 {
			result[b.id] = 0
		}
	}

	return result
}
