# 3162 — Find The Number Of Good Pairs I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

**Fungsi yang perlu kamu implementasikan:**
```go
func FindTheNumberOfGoodPairsI(nums1 []int, nums2 []int, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n * m)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3162: Find the Number of Good Pairs I
// https://leetcode.com/problems/find-the-number-of-good-pairs-i/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: numberOfPairs
	fmt.Println(FindTheNumberOfGoodPairsI([]int{1, 2, 4, 12}, []int{2, 4}, 3)) // 2
	fmt.Println(FindTheNumberOfGoodPairsI([]int{1, 2, 3}, []int{1, 2, 3}, 1))  // 5
}

// Time: O(n * m) | Space: O(1)
// LeetCode submission name: numberOfPairs
func FindTheNumberOfGoodPairsI(nums1 []int, nums2 []int, k int) int {
	count := 0
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			if nums1[i]%(nums2[j]*k) == 0 {
				count++
			}
		}
	}
	return count
}
```
