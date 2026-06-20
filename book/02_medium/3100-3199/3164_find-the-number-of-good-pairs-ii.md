# 3164 — Find The Number Of Good Pairs Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

**Fungsi yang perlu kamu implementasikan:**
```go
func numberOfPairs(nums1 []int, nums2 []int, k int) int64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n * sqrt(max) + m)  
**Kompleksitas Ruang:** O(max)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3164: Find the Number of Good Pairs II
// https://leetcode.com/problems/find-the-number-of-good-pairs-ii/
// Difficulty: Medium
// Time: O(n * sqrt(max) + m) | Space: O(max)

import "fmt"

func numberOfPairs(nums1 []int, nums2 []int, k int) int64 {
  // Membuat map (HashMap) — pencarian O(1)
	freq := make(map[int]int)
	for _, v := range nums1 {
		if v%k != 0 {
			continue
		}
		v /= k
		for d := 1; d*d <= v; d++ {
			if v%d == 0 {
				freq[d]++
				if d*d != v {
					freq[v/d]++
				}
			}
		}
	}

	var ans int64
	for _, v := range nums2 {
		ans += int64(freq[v])
	}
	return ans
}

func main() {
	fmt.Println(numberOfPairs([]int{1, 3, 4}, []int{1, 3, 4}, 1)) // Expected: 5
	fmt.Println(numberOfPairs([]int{1, 2, 4, 12}, []int{2, 4}, 3)) // Expected: 2
}
```
