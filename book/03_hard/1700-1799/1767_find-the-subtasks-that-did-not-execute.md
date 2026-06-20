# 1767 — Find The Subtasks That Did Not Execute

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

**Fungsi yang perlu kamu implementasikan:**
```go
func findSubtasksThatDidNotExecute(tasks []Task, executed []ExecutedSubtask) []ExecutedSubtask
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1767: Find the Subtasks That Did Not Execute
// https://leetcode.com/problems/find-the-subtasks-that-did-not-execute/
// Difficulty: Hard [Premium]

import (
	"fmt"
	"sort"
)

type Task struct {
	TaskId         int
	SubtasksCount  int
}

type ExecutedSubtask struct {
	TaskId    int
	SubtaskId int
}

func main() {
	tasks := []Task{
		{1, 3},
		{2, 2},
		{3, 4},
	}
	executed := []ExecutedSubtask{
		{1, 1},
		{1, 2},
		{2, 1},
		{3, 2},
		{3, 4},
	}

	result := findSubtasksThatDidNotExecute(tasks, executed)
	fmt.Println("Test 1 - Subtasks that did not execute:")
	fmt.Print("Expected: [(1,3) (2,2) (3,1) (3,3)]\nGot:      ")
	for _, r := range result {
		fmt.Printf("(%d,%d) ", r.TaskId, r.SubtaskId)
	}
	fmt.Println()

	// Test 2: empty (all executed)
	tasks2 := []Task{{1, 2}}
	executed2 := []ExecutedSubtask{{1, 1}, {1, 2}}
	result2 := findSubtasksThatDidNotExecute(tasks2, executed2)
	fmt.Println("\nTest 2 (all executed):")
	fmt.Print("Expected: []\nGot:      ")
	for _, r := range result2 {
		fmt.Printf("(%d,%d) ", r.TaskId, r.SubtaskId)
	}
	fmt.Println()
}

func findSubtasksThatDidNotExecute(tasks []Task, executed []ExecutedSubtask) []ExecutedSubtask {
  // Membuat map (HashMap) — pencarian O(1)
	executedSet := make(map[[2]int]bool)
	for _, e := range executed {
		executedSet[[2]int{e.TaskId, e.SubtaskId}] = true
	}

	notExecuted := make([]ExecutedSubtask, 0)
	for _, t := range tasks {
		for sub := 1; sub <= t.SubtasksCount; sub++ {
			if !executedSet[[2]int{t.TaskId, sub}] {
				notExecuted = append(notExecuted, ExecutedSubtask{TaskId: t.TaskId, SubtaskId: sub})
			}
		}
	}

  // Custom sort dengan comparator
	sort.Slice(notExecuted, func(i, j int) bool {
		if notExecuted[i].TaskId != notExecuted[j].TaskId {
			return notExecuted[i].TaskId < notExecuted[j].TaskId
		}
		return notExecuted[i].SubtaskId < notExecuted[j].SubtaskId
	})

	return notExecuted
}
```
