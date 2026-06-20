# 3796 — Find Maximum Value In A Constrained Sequence

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

**Fungsi yang perlu kamu implementasikan:**
```go
func FindMaximumValueInAConstrainedSequence(n int, restrictions [][]int, diff []int) int
```

> **💡 Hint:** Two-pass greedy constraint propagation. Forward pass applies

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3796: Find Maximum Value in a Constrained Sequence
// https://leetcode.com/problems/find-maximum-value-in-a-constrained-sequence/
// Difficulty: Medium
// Time: O(n) | Space: O(n)
// Approach: Two-pass greedy constraint propagation. Forward pass applies
// constraints from left to right, backward pass propagates from right to left.

import "fmt"

func FindMaximumValueInAConstrainedSequence(n int, restrictions [][]int, diff []int) int {
	const INF = 1 << 60
  // Alokasi slice integer
	a := make([]int, n)
  // Range loop: iterasi dengan indeks + nilai
	for i := range a {
		a[i] = INF
	}
	a[0] = 0

	// Apply restrictions
	for _, r := range restrictions {
		idx, maxVal := r[0], r[1]
		if a[idx] > maxVal {
			a[idx] = maxVal
		}
	}

	// Forward pass: propagate constraints left to right
	for i := 1; i < n; i++ {
		limit := a[i-1] + diff[i-1]
		if a[i] > limit {
			a[i] = limit
		}
	}

	// Backward pass: propagate constraints right to left
	for i := n - 2; i >= 0; i-- {
		limit := a[i+1] + diff[i]
		if a[i] > limit {
			a[i] = limit
		}
	}

	// Find maximum value
	ans := 0
	for _, v := range a {
		if v > ans {
			ans = v
		}
	}
	return ans
}

func main() {
	// Example 1
	n1 := 10
	restrictions1 := [][]int{{3, 1}, {8, 1}}
	diff1 := []int{2, 2, 3, 1, 4, 5, 1, 1, 2}
	fmt.Println(FindMaximumValueInAConstrainedSequence(n1, restrictions1, diff1)) // Expected: 6

	// Example 2
	n2 := 8
	restrictions2 := [][]int{{3, 2}}
	diff2 := []int{3, 5, 2, 4, 2, 3, 1}
	fmt.Println(FindMaximumValueInAConstrainedSequence(n2, restrictions2, diff2)) // Expected: 12
}
```
