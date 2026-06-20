# 1562 — Find Latest Group Of Size M

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

**Fungsi yang perlu kamu implementasikan:**
```go
func FindLatestStep(arr []int, m int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer, Union-Find (DSU)

**Kompleksitas Waktu:** O(N), Space: O(N)  
**Kompleksitas Ruang:** O(N)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1562: Find Latest Group of Size M
// https://leetcode.com/problems/find-latest-group-of-size-m/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(FindLatestStep([]int{3, 5, 1, 2, 4}, 1))
	fmt.Println(FindLatestStep([]int{3, 1, 5, 4, 2}, 2))
	fmt.Println(FindLatestStep([]int{1}, 1))
}

func FindLatestStep(arr []int, m int) int {
	// Time: O(N), Space: O(N)
	// Use union-find-like approach tracking group lengths
	n := len(arr)
	if n == m {
		return n
	}

  // Alokasi slice integer
	length := make([]int, n+2) // length of group at position i
  // Alokasi slice integer
	count := make([]int, n+2)  // count of groups of length i
	result := -1

	for step := 0; step < n; step++ {
		pos := arr[step]
		leftLen := length[pos-1]
		rightLen := length[pos+1]
		total := leftLen + rightLen + 1

		// Decrement counts for the merging groups
		count[leftLen]--
		count[rightLen]--
		// Increment count for the new merged group
		count[total]++

		// Update lengths at boundaries
		length[pos-leftLen] = total
		length[pos+rightLen] = total

		// Check if we have exactly m groups of size m
		if count[m] > 0 {
			// This step is valid (but since we want the latest step before m disappears,
			// we'll track the result after the step)
		}

		// The result is the latest step where count[m] > 0
		// But we want the latest step where a group of size m *exists*
		// We need to check AFTER the current operation
	}

	// Re-simulate to find latest step with count[m] > 0
	length = make([]int, n+2)
	count = make([]int, n+2)

	for step := 0; step < n; step++ {
		pos := arr[step]
		leftLen := length[pos-1]
		rightLen := length[pos+1]
		total := leftLen + rightLen + 1

		count[leftLen]--
		count[rightLen]--
		count[total]++

		length[pos-leftLen] = total
		length[pos+rightLen] = total

		if count[m] > 0 {
			result = step + 1 // 1-indexed step
		}
	}

	return result
}
```
