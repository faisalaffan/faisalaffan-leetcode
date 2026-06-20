# 1673 — Find The Most Competitive Subsequence

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

**Fungsi yang perlu kamu implementasikan:**
```go
func mostCompetitive(nums []int, k int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Stack

**Kompleksitas Waktu:** O(n), Space: O(k)  
**Kompleksitas Ruang:** O(k)

> **Untuk fresh graduate:** Kuasai dulu teknik **Stack** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1673: Find the Most Competitive Subsequence
// https://leetcode.com/problems/find-the-most-competitive-subsequence/
// Difficulty: Medium
// Time: O(n), Space: O(k)

import "fmt"

func mostCompetitive(nums []int, k int) []int {
	n := len(nums)
  // Alokasi slice integer
	stack := make([]int, 0, k)
	toRemove := n - k

	for _, num := range nums {
		for len(stack) > 0 && toRemove > 0 && stack[len(stack)-1] > num {
			stack = stack[:len(stack)-1]
			toRemove--
		}
		stack = append(stack, num)
	}

	// If we haven't removed enough, trim from end
	return stack[:k]
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", mostCompetitive([]int{3, 5, 2, 6}, 2)) // Expected: [2, 6]

	// Test case 2
	fmt.Println("Test 2:", mostCompetitive([]int{2, 4, 3, 3, 5, 4, 9, 6}, 4)) // Expected: [2, 3, 3, 4]

	// Test case 3
	fmt.Println("Test 3:", mostCompetitive([]int{1, 2, 3, 4, 5}, 3)) // Expected: [1, 2, 3]

	// Test case 4
	fmt.Println("Test 4:", mostCompetitive([]int{5, 4, 3, 2, 1}, 2)) // Expected: [1, 2]
}
```
