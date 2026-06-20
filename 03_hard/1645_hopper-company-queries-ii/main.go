package main

// LeetCode #1645: Hopper Company Queries II
// https://leetcode.com/problems/hopper-company-queries-ii/
// Difficulty: Hard [Paid] (SQL)
//
// For each month of 2020, report the percentage of working drivers
// (drivers who accepted at least one ride that month) relative to
// the number of available drivers by month-end.
//
// Approach: Process driver and ride data in Go.

import (
	"fmt"
	"time"
)

func main() {
	// Example data
	drivers := []struct {
		id       int
		joinDate string
	}{
		{10, "2019-12-10"},
		{8, "2020-1-13"},
		{5, "2020-2-16"},
		{7, "2020-3-8"},
		{4, "2020-5-17"},
		{1, "2020-10-24"},
		{6, "2021-1-5"},
	}
	rides := []struct {
		id       int
		reqDate  string
	}{
		{6, "2019-12-9"},
		{1, "2020-1-2"},
		{10, "2020-1-11"},
		{11, "2020-1-19"},
		{12, "2020-1-29"},
		{3, "2020-1-17"},
		{9, "2020-2-13"},
		{2, "2020-2-22"},
	}
	accepted := []struct {
		rideID   int
		driverID int
	}{
		{2, 10},
		{13, 10},
		{7, 8},
		{3, 5},
		{4, 5},
		{11, 7},
		{1, 10},
		{5, 7},
		{12, 8},
		{6, 7},
	}

	fmt.Println(hopperQueriesII(drivers, rides, accepted))
}

func hopperQueriesII(drivers []struct {
	id       int
	joinDate string
}, rides []struct {
	id      int
	reqDate string
}, accepted []struct {
	rideID   int
	driverID int
}) []struct {
	month             int
	workingPercentage float64
} {
	// Build a set of accepted ride IDs
	accSet := make(map[int]bool)
	for _, a := range accepted {
		accSet[a.rideID] = true
	}

	// Map ride ID -> month and month -> set of working drivers
	rideMonth := make(map[int]int)
	for _, r := range rides {
		t, _ := time.Parse("2006-1-2", r.reqDate)
		if t.Year() == 2020 {
			rideMonth[r.id] = int(t.Month())
		}
	}

	workingDrivers := make([]map[int]bool, 13)
	for i := range workingDrivers {
		workingDrivers[i] = make(map[int]bool)
	}
	for rideID, month := range rideMonth {
		if accSet[rideID] {
			for _, a := range accepted {
				if a.rideID == rideID {
					workingDrivers[month][a.driverID] = true
				}
			}
		}
	}

	// Count available drivers by end of each month
	available := make([]int, 13)
	for _, d := range drivers {
		t, _ := time.Parse("2006-1-2", d.joinDate)
		joinYear, joinMonth := t.Year(), t.Month()
		if joinYear < 2020 {
			for m := 1; m <= 12; m++ {
				available[m]++
			}
		} else if joinYear == 2020 {
			for m := int(joinMonth); m <= 12; m++ {
				available[m]++
			}
		}
	}

	result := make([]struct {
		month             int
		workingPercentage float64
	}, 12)
	for m := 1; m <= 12; m++ {
		pct := 0.0
		if available[m] > 0 {
			pct = float64(len(workingDrivers[m])) / float64(available[m]) * 100.0
		}
		result[m-1] = struct {
			month             int
			workingPercentage float64
		}{m, pct}
	}
	return result
}
