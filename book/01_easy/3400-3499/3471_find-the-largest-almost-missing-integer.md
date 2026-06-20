# 3471 — Find The Largest Almost Missing Integer

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

**Fungsi yang perlu kamu implementasikan:**
```go
func FindTheLargestAlmostMissingInteger(nums []int, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n). Space: O(n).  
**Kompleksitas Ruang:** O(n).

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3471: Find the Largest Almost Missing Integer
// https://leetcode.com/problems/find-the-largest-almost-missing-integer/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(FindTheLargestAlmostMissingInteger([]int{3, 9, 2, 3, 1, 6, 7, 8, 9}, 2))
	fmt.Println(FindTheLargestAlmostMissingInteger([]int{0, 0}, 1))
}

// FindTheLargestAlmostMissingInteger returns the largest integer that appears fewer than k times in nums.
// Time: O(n). Space: O(n).
func FindTheLargestAlmostMissingInteger(nums []int, k int) int {
  // Membuat map (HashMap) — pencarian O(1)
	freq := make(map[int]int)
	for _, v := range nums {
		freq[v]++
	}
	largest := -1
	for val, count := range freq {
		if count < k && val > largest {
			largest = val
		}
	}
	return largest
}
```
