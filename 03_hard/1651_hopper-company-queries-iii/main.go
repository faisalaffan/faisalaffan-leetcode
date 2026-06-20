package main

// LeetCode #1651: Hopper Company Queries III
// https://leetcode.com/problems/hopper-company-queries-iii/
// Difficulty: Hard [Paid] (SQL)
//
// For each 3-month window (Jan-Mar, Feb-Apr, ..., Oct-Dec) of 2020,
// compute the average ride distance and average ride duration.
//
// Approach: Aggregate ride stats by month, then compute rolling
// 3-month averages.

import (
	"fmt"
	"time"
)

func main() {
	// Example data
	rides := []struct {
		rideID   int
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
		rideID  int
		dist    int
		dur     int
	}{
		{2, 63, 38},
		{13, 58, 0},
		{7, 51, 23},
		{3, 50, 36},
		{4, 58, 19},
		{11, 53, 26},
		{1, 50, 42},
		{5, 55, 42},
		{12, 51, 23},
		{6, 56, 41},
	}

	fmt.Println(hopperQueriesIII(rides, accepted))
}

func hopperQueriesIII(rides []struct {
	rideID  int
	reqDate string
}, accepted []struct {
	rideID int
	dist   int
	dur    int
}) []struct {
	month                 int
	avgRideDistance       float64
	avgRideDuration       float64
} {
	// Map ride ID -> month (only 2020)
	rideMonth := make(map[int]int)
	rideData := make(map[int]struct{ dist, dur int })
	for _, r := range rides {
		t, _ := time.Parse("2006-1-2", r.reqDate)
		if t.Year() == 2020 {
			rideMonth[r.rideID] = int(t.Month())
		}
	}
	for _, a := range accepted {
		rideData[a.rideID] = struct{ dist, dur int }{a.dist, a.dur}
	}

	// Monthly totals
	monthDist := make([]int, 13)
	monthDur := make([]int, 13)
	monthCount := make([]int, 13)

	for rideID, m := range rideMonth {
		if d, ok := rideData[rideID]; ok {
			monthDist[m] += d.dist
			monthDur[m] += d.dur
			monthCount[m]++
		}
	}

	result := make([]struct {
		month                 int
		avgRideDistance       float64
		avgRideDuration       float64
	}, 10)

	for start := 1; start <= 10; start++ {
		totalDist := 0
		totalDur := 0
		for m := start; m <= start+2; m++ {
			totalDist += monthDist[m]
			totalDur += monthDur[m]
		}
		result[start-1] = struct {
			month                 int
			avgRideDistance       float64
			avgRideDuration       float64
		}{start, float64(totalDist) / 3.0, float64(totalDur) / 3.0}
	}

	return result
}
