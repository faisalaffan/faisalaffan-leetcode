# 3124 — Find Longest Calls

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

**Fungsi yang perlu kamu implementasikan:**
```go
func findLongestCalls(calls [][]int, k int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3124: Find Longest Calls
// https://leetcode.com/problems/find-longest-calls/
// Difficulty: Medium [Paid]
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func findLongestCalls(calls [][]int, k int) []int {
	type call struct {
		id     int
		dur    int
		caller int
	}

	var list []call
	for _, c := range calls {
		list = append(list, call{c[0], c[2] - c[1], c[3]})
	}

  // Custom sort dengan comparator
	sort.Slice(list, func(i, j int) bool {
		if list[i].dur != list[j].dur {
			return list[i].dur > list[j].dur
		}
		return list[i].id < list[j].id
	})

  // Alokasi slice integer
	ans := make([]int, 0, k)
	for i := 0; i < k && i < len(list); i++ {
		ans = append(ans, list[i].id)
	}
	return ans
}

func main() {
	fmt.Println(findLongestCalls([][]int{{1, 0, 30, 1}, {2, 5, 25, 2}, {3, 10, 20, 1}}, 2)) // Expected: [1 2]
	fmt.Println(findLongestCalls([][]int{{1, 0, 10, 1}, {2, 0, 5, 2}}, 3))                    // Expected: [1 2]
}
```
