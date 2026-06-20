# 3709 — Design Exam Scores Tracker

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diminta untuk mendesain (merancang) sebuah struktur data kustom dengan operasi tertentu (insert, delete, search, update). Tugasmu adalah memilih representasi data yang tepat agar setiap operasi berjalan efisien — biasanya O(1) atau O(log n).

Ini adalah soal yang paling sering muncul di interview sistem desain. Kamu perlu memilih kombinasi struktur data yang tepat (HashMap + Heap + LinkedList) untuk mencapai kompleksitas yang diminta.

**Konsep kunci:** HashMap (O(1) lookup), Heap (priority), Doubly Linked List (O(1) remove), TreeMap (ordered keys).

**Fungsi yang perlu kamu implementasikan:**
```go
func NewExamScoresTracker() *ExamScoresTracker
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(log n) per query  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3709: Design Exam Scores Tracker
// https://leetcode.com/problems/design-exam-scores-tracker/
// Difficulty: Medium
// Time: O(log n) per query | Space: O(n)

import (
	"fmt"
	"sort"
)

type ExamScoresTracker struct {
	times []int
	pref  []int64
}

func NewExamScoresTracker() *ExamScoresTracker {
	return &ExamScoresTracker{
		times: []int{0},
		pref:  []int64{0},
	}
}

func (t *ExamScoresTracker) Record(time int, score int) {
	t.times = append(t.times, time)
	t.pref = append(t.pref, t.pref[len(t.pref)-1]+int64(score))
}

func (t *ExamScoresTracker) TotalScore(startTime int, endTime int) int64 {
	l := sort.SearchInts(t.times, startTime)
	if l < len(t.times) && t.times[l] < startTime {
		l++
	}
	l--

	r := sort.SearchInts(t.times, endTime+1) - 1

	if l < 0 {
		l = 0
	}
	if r < 0 {
		return 0
	}
	if r >= len(t.pref) {
		r = len(t.pref) - 1
	}
	return t.pref[r] - t.pref[l]
}

func main() {
	tracker := NewExamScoresTracker()
	tracker.Record(10, 5)
	tracker.Record(20, 3)
	tracker.Record(30, 7)
	fmt.Println(tracker.TotalScore(10, 20))
	fmt.Println(tracker.TotalScore(15, 25))
	fmt.Println(tracker.TotalScore(5, 35))

	tracker2 := NewExamScoresTracker()
	fmt.Println(tracker2.TotalScore(0, 100))
}
```
