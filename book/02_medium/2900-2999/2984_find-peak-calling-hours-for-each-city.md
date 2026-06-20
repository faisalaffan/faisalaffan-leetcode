# 2984 — Find Peak Calling Hours For Each City

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan data tabel database. Tugasmu adalah menganalisis data tersebut. Karena repo ini Go, query SQL disimulasikan dengan map, slice, dan struct.

**Cara berpikir:** Tentukan SELECT, FROM, JOIN, GROUP BY, ORDER BY. Lalu terjemahkan ke Go.

**Fungsi Solusi:** `func extractHour(datetime string) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Sorting

**Waktu:** O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2984: Find Peak Calling Hours for Each City
// https://leetcode.com/problems/find-peak-calling-hours-for-each-city/
// Difficulty: Medium (SQL problem — implemented in Go)
// Time: O(n) | Space: O(n)

import (
	"fmt"
	"sort"
)

type Call struct {
	CallerID    int
	RecipientID int
	CallTime    string // datetime string, e.g. "2024-01-01 14:30:00"
	City        string
}

type PeakHour struct {
	City      string
	PeakHour  int
	CallCount int
}

func extractHour(datetime string) int {
	// Format: "2024-01-01 14:30:00"
	// Extract the hour (two digits after the first space and before the first colon after space)
	var year, month, day, hour, min, sec int
	_, err := fmt.Sscanf(datetime, "%d-%d-%d %d:%d:%d", &year, &month, &day, &hour, &min, &sec)
	if err != nil {
		return 0
	}
	return hour
}

func findPeakCallingHours(calls []Call) []PeakHour {
	// Count calls per city per hour: map[city]map[hour]count
  // HashMap: O(1) lookup
	cityHourCounts := make(map[string]map[int]int)
	for _, c := range calls {
		hour := extractHour(c.CallTime)
		if cityHourCounts[c.City] == nil {
			cityHourCounts[c.City] = make(map[int]int)
		}
		cityHourCounts[c.City][hour]++
	}

	var results []PeakHour

	for city, hourCounts := range cityHourCounts {
		// Find max count for this city
		maxCount := 0
		for _, count := range hourCounts {
			if count > maxCount {
				maxCount = count
			}
		}

		// Collect all hours with max count
		for hour, count := range hourCounts {
			if count == maxCount {
				results = append(results, PeakHour{
					City:      city,
					PeakHour:  hour,
					CallCount: count,
				})
			}
		}
	}

	// Order by peak_hour DESC, city DESC
  // Custom sort
	sort.Slice(results, func(i, j int) bool {
		if results[i].PeakHour != results[j].PeakHour {
			return results[i].PeakHour > results[j].PeakHour // DESC
		}
		return results[i].City > results[j].City // DESC
	})

	return results
}

func main() {
	calls := []Call{
		{CallerID: 1, RecipientID: 2, CallTime: "2024-01-01 09:15:00", City: "New York"},
		{CallerID: 3, RecipientID: 4, CallTime: "2024-01-01 09:30:00", City: "New York"},
		{CallerID: 5, RecipientID: 6, CallTime: "2024-01-01 14:00:00", City: "New York"},
		{CallerID: 7, RecipientID: 8, CallTime: "2024-01-01 09:45:00", City: "New York"},
		{CallerID: 9, RecipientID: 10, CallTime: "2024-01-01 10:00:00", City: "Los Angeles"},
		{CallerID: 11, RecipientID: 12, CallTime: "2024-01-01 10:15:00", City: "Los Angeles"},
		{CallerID: 13, RecipientID: 14, CallTime: "2024-01-01 14:00:00", City: "Los Angeles"},
		{CallerID: 15, RecipientID: 16, CallTime: "2024-01-01 10:30:00", City: "Los Angeles"},
		{CallerID: 17, RecipientID: 18, CallTime: "2024-01-01 10:00:00", City: "Chicago"},
		{CallerID: 19, RecipientID: 20, CallTime: "2024-01-01 11:00:00", City: "Chicago"},
		{CallerID: 21, RecipientID: 22, CallTime: "2024-01-01 10:00:00", City: "Chicago"},
	}

	fmt.Println("Find Peak Calling Hours for Each City")
	fmt.Println("====================================")
	fmt.Printf("%-15s %-12s %s\n", "City", "Peak Hour", "Call Count")
	fmt.Println("--------------------------------------")

	results := findPeakCallingHours(calls)
	for _, r := range results {
		fmt.Printf("%-15s %-12d %d\n", r.City, r.PeakHour, r.CallCount)
	}
}
```
