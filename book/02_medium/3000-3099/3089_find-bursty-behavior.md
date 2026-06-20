# 3089 — Find Bursty Behavior

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

**Fungsi yang perlu kamu implementasikan:**
```go
func findBurstyBehavior(posts [][]int, k int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3089: Find Bursty Behavior
// https://leetcode.com/problems/find-bursty-behavior/
// Difficulty: Medium [Paid]
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func findBurstyBehavior(posts [][]int, k int) []int {
	n := len(posts)
  // Edge case: input kosong — langsung return
	if n == 0 {
		return nil
	}

  // Membuat map (HashMap) — pencarian O(1)
	userPosts := make(map[int][]int)
	for _, p := range posts {
		userID, timestamp := p[0], p[1]
		userPosts[userID] = append(userPosts[userID], timestamp)
	}

	var bursty []int
	for uid, timestamps := range userPosts {
  // Urutkan secara ascending — O(n log n)
		sort.Ints(timestamps)
		for i := k - 1; i < len(timestamps); i++ {
			if timestamps[i]-timestamps[i-k+1] <= 100 {
				bursty = append(bursty, uid)
				break
			}
		}
	}

  // Urutkan secara ascending — O(n log n)
	sort.Ints(bursty)
	return bursty
}

func main() {
	fmt.Println(findBurstyBehavior([][]int{{1, 10}, {1, 20}, {1, 30}, {2, 5}, {2, 200}}, 3)) // Expected: [1]
	fmt.Println(findBurstyBehavior([][]int{{1, 1}, {2, 2}, {3, 3}}, 2))                       // Expected: []
}
```
