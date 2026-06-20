# 3156 — Employee Task Duration And Concurrent Tasks

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func employeeTaskDurationAndConcurrentTasks(tasks []Task) []EmployeeResult
```

> **💡 Hint:** group tasks by employee, then sweep-line for concurrency.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3156: Employee Task Duration and Concurrent Tasks
// https://leetcode.com/problems/employee-task-duration-and-concurrent-tasks/
// Difficulty: Hard [Paid]
//
// For each employee, compute:
//   1. Total duration of all assigned tasks (sum of end - start for each task)
//   2. Maximum number of tasks running concurrently at any point in time
//
// Approach: group tasks by employee, then sweep-line for concurrency.

import (
	"fmt"
	"sort"
)

type Task struct {
	Start int
	End   int
	EmpID int
}

type EmployeeResult struct {
	EmpID             int
	TotalDuration     int
	MaxConcurrent     int
}

func employeeTaskDurationAndConcurrentTasks(tasks []Task) []EmployeeResult {
	if len(tasks) == 0 {
		return nil
	}

	// Group tasks by employee.
	type interval struct{ start, end int }
  // Membuat map (HashMap) — pencarian O(1)
	empMap := make(map[int][]interval)
	for _, t := range tasks {
		empMap[t.EmpID] = append(empMap[t.EmpID], interval{t.Start, t.End})
	}

	// Collect and sort employee IDs.
  // Alokasi slice integer
	empIDs := make([]int, 0, len(empMap))
	for id := range empMap {
		empIDs = append(empIDs, id)
	}
  // Urutkan secara ascending — O(n log n)
	sort.Ints(empIDs)

	res := make([]EmployeeResult, len(empIDs))
	for idx, empID := range empIDs {
		intervals := empMap[empID]

		// Total duration: sum of all interval lengths.
		total := 0
		// Events for sweep-line.
		type event struct{ pos, delta int }
  // Alokasi slice integer
		events := make([]event, 0, len(intervals)*2)
		for _, iv := range intervals {
			total += iv.end - iv.start
			events = append(events, event{iv.start, 1})
			events = append(events, event{iv.end, -1})
		}

		// Sort events: by position, then end (-1) before start (+1).
  // Custom sort dengan comparator
		sort.Slice(events, func(i, j int) bool {
			if events[i].pos != events[j].pos {
				return events[i].pos < events[j].pos
			}
			return events[i].delta < events[j].delta
		})

		cur := 0
		maxCur := 0
		for _, e := range events {
			cur += e.delta
			if cur > maxCur {
				maxCur = cur
			}
		}

		res[idx] = EmployeeResult{
			EmpID:         empID,
			TotalDuration: total,
			MaxConcurrent: maxCur,
		}
	}
	return res
}

func main() {
	// Example
	tasks := []Task{
		{0, 5, 1},
		{2, 7, 1},
		{1, 3, 2},
	}
	fmt.Println(employeeTaskDurationAndConcurrentTasks(tasks))
	// Expect: [{1 10 2} {2 2 1}]
}
```
