# 3721 — Longest Balanced Subarray Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func longestBalanced(nums []int) int
```

> **💡 Hint:** Sliding window with frequency maps. For each window,

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, Sliding Window

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3721: Longest Balanced Subarray II
// https://leetcode.com/problems/longest-balanced-subarray-ii/
// Difficulty: Hard
//
// Find longest subarray where number of distinct even numbers equals
// number of distinct odd numbers.
//
// Approach: Sliding window with frequency maps. For each window,
// track distinct even and odd counts using hash maps.

import "fmt"

func main() {
	// Example 1
	fmt.Println(longestBalanced([]int{2, 5, 4, 3}))
	// Example 2
	fmt.Println(longestBalanced([]int{1, 3, 5, 2, 4}))
	// Edge: single element
	fmt.Println(longestBalanced([]int{1}))
	// Edge: all even
	fmt.Println(longestBalanced([]int{2, 4, 6}))
}

func longestBalanced(nums []int) int {
	n := len(nums)
	result := 0

	// For each starting position, expand window
	for i := 0; i < n; i++ {
  // Membuat map (HashMap) — pencarian O(1)
		evenSet := make(map[int]bool)
  // Membuat map (HashMap) — pencarian O(1)
		oddSet := make(map[int]bool)
		evenCount := 0
		oddCount := 0

		for j := i; j < n; j++ {
			if nums[j]%2 == 0 {
				if !evenSet[nums[j]] {
					evenSet[nums[j]] = true
					evenCount++
				}
			} else {
				if !oddSet[nums[j]] {
					oddSet[nums[j]] = true
					oddCount++
				}
			}
			if evenCount == oddCount {
				length := j - i + 1
				if length > result {
					result = length
				}
			}
		}
	}

	return result
}
```
