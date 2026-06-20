# 3421 — Find Students Who Improved

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

**Fungsi yang perlu kamu implementasikan:**
```go
func findStudentsWhoImproved(scores []Score) []ImprovedStudent
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n) Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3421: Find Students Who Improved
// https://leetcode.com/problems/find-students-who-improved/
// Difficulty: Medium (SQL problem — implemented in Go)
// Time: O(n log n) Space: O(n)

import (
	"fmt"
	"sort"
)

type Score struct {
	StudentID int
	Subject   string
	Score     int
	ExamDate  string
}

type ImprovedStudent struct {
	StudentID   int
	Subject     string
	FirstScore  int
	LatestScore int
}

func findStudentsWhoImproved(scores []Score) []ImprovedStudent {
  // Membuat map (HashMap) — pencarian O(1)
	group := make(map[[2]interface{}][]Score)

	for _, s := range scores {
		key := [2]interface{}{s.StudentID, s.Subject}
		group[key] = append(group[key], s)
	}

	var result []ImprovedStudent
	for key, exams := range group {
  // Custom sort dengan comparator
		sort.Slice(exams, func(i, j int) bool { return exams[i].ExamDate < exams[j].ExamDate })
		first := exams[0].Score
		last := exams[len(exams)-1].Score
		if last > first && len(exams) >= 2 {
			result = append(result, ImprovedStudent{
				StudentID:   key[0].(int),
				Subject:     key[1].(string),
				FirstScore:  first,
				LatestScore: last,
			})
		}
	}

  // Custom sort dengan comparator
	sort.Slice(result, func(i, j int) bool {
		if result[i].StudentID != result[j].StudentID {
			return result[i].StudentID < result[j].StudentID
		}
		return result[i].Subject < result[j].Subject
	})
	return result
}

func main() {
	scores := []Score{
		{1, "Math", 70, "2023-01-15"},
		{1, "Math", 85, "2023-02-15"},
		{2, "Science", 60, "2023-01-15"},
		{2, "Science", 55, "2023-02-15"},
		{1, "Science", 80, "2023-01-15"},
		{1, "Science", 90, "2023-02-15"},
	}
	result := findStudentsWhoImproved(scores)
	for _, r := range result {
		fmt.Printf("Student %d improved in %s: %d -> %d\n", r.StudentID, r.Subject, r.FirstScore, r.LatestScore)
	}
}
```
