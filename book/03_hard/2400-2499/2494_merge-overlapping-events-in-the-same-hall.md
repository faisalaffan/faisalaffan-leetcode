# 2494 — Merge Overlapping Events In The Same Hall

## Deskripsi

**Soal:** [2494. Merge Overlapping Events In The Same Hall](https://leetcode.com/problems/merge-overlapping-events-in-the-same-hall/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

> **Ide Kunci:** Sort events by hall and start_day, then merge overlapping

## Solusi Go

```go
package main

// LeetCode #2494: Merge Overlapping Events in the Same Hall
// https://leetcode.com/problems/merge-overlapping-events-in-the-same-hall/
// Difficulty: Hard [Paid] (SQL)
//
// Merge overlapping events held in the same hall. Two events overlap
// if they share at least one day.
//
// Approach: Sort events by hall and start_day, then merge overlapping
// ranges per hall.

import (
	"fmt"
	"sort"
)

// HallEvent represents an event in a hall
type HallEvent struct {
	HallID   int
	StartDay string // "YYYY-MM-DD"
	EndDay   string // "YYYY-MM-DD"
}

func main() {
	// Example
	events := []HallEvent{
		{1, "2023-01-13", "2023-01-14"},
		{1, "2023-01-14", "2023-01-17"},
		{1, "2023-01-18", "2023-01-25"},
		{2, "2023-01-01", "2023-01-02"},
		{2, "2023-01-02", "2023-01-03"},
		{2, "2023-02-01", "2023-02-05"},
		{3, "2023-03-01", "2023-03-10"},
	}

	fmt.Println(mergeOverlappingEvents(events))
}

func mergeOverlappingEvents(events []HallEvent) []HallEvent {
	// Sort by hall_id, then start_day
	sort.Slice(events, func(i, j int) bool {
		if events[i].HallID != events[j].HallID {
			return events[i].HallID < events[j].HallID
		}
		return events[i].StartDay < events[j].StartDay
	})

	var result []HallEvent
	i := 0
	for i < len(events) {
		hallID := events[i].HallID
		start := events[i].StartDay
		end := events[i].EndDay
		i++

		for i < len(events) && events[i].HallID == hallID {
			// Check if overlapping (events[i].StartDay <= end means overlap)
			if events[i].StartDay <= end {
				if events[i].EndDay > end {
					end = events[i].EndDay
				}
				i++
			} else {
				break
			}
		}

		result = append(result, HallEvent{hallID, start, end})
	}

	return result
}
```
