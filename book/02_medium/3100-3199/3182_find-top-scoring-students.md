# 3182 — Find Top Scoring Students

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

**Fungsi yang perlu kamu implementasikan:**
```go
func findTopScoringStudents(scores [][]int, threshold int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3182: Find Top Scoring Students
// https://leetcode.com/problems/find-top-scoring-students/
// Difficulty: Medium [Paid]
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func findTopScoringStudents(scores [][]int, threshold int) []int {
	type student struct {
		id    int
		total int
	}

	var list []student
	for _, s := range scores {
		list = append(list, student{s[0], s[1]})
	}

  // Custom sort dengan comparator
	sort.Slice(list, func(i, j int) bool {
		if list[i].total != list[j].total {
			return list[i].total > list[j].total
		}
		return list[i].id < list[j].id
	})

  // Alokasi slice integer
	ans := make([]int, 0)
	for _, s := range list {
		if s.total >= threshold {
			ans = append(ans, s.id)
		}
	}
	return ans
}

func main() {
	fmt.Println(findTopScoringStudents([][]int{{1, 95}, {2, 85}, {3, 90}}, 90)) // Expected: [1 3]
	fmt.Println(findTopScoringStudents([][]int{{1, 70}, {2, 65}}, 80))          // Expected: []
}
```
