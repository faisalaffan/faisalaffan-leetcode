# 1817 — Finding The Users Active Minutes

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

**Fungsi yang perlu kamu implementasikan:**
```go
func findingUsersActiveMinutes(logs [][]int, k int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1817: Finding the Users Active Minutes
// https://leetcode.com/problems/finding-the-users-active-minutes/
// Difficulty: Medium
// Time: O(n), Space: O(n)

import "fmt"

func findingUsersActiveMinutes(logs [][]int, k int) []int {
  // Membuat map (HashMap) — pencarian O(1)
	userMinutes := make(map[int]map[int]bool)
	for _, log := range logs {
		id, min := log[0], log[1]
		if userMinutes[id] == nil {
			userMinutes[id] = make(map[int]bool)
		}
		userMinutes[id][min] = true
	}

  // Alokasi slice integer
	result := make([]int, k)
	for _, minutes := range userMinutes {
		uam := len(minutes)
		if uam <= k {
			result[uam-1]++
		}
	}
	return result
}

func main() {
	fmt.Println(findingUsersActiveMinutes([][]int{{0, 5}, {1, 2}, {0, 2}, {0, 5}, {1, 3}}, 5)) // Expected: [0, 2, 0, 0, 0]
	fmt.Println(findingUsersActiveMinutes([][]int{{1, 1}, {2, 2}, {2, 3}}, 4)) // Expected: [1, 1, 0, 0]
}
```
