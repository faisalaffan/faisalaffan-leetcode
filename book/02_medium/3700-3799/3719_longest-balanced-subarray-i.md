# 3719 — Longest Balanced Subarray I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func longestBalancedSubarrayI(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n^2)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3719: Longest Balanced Subarray I
// https://leetcode.com/problems/longest-balanced-subarray-i/
// Difficulty: Medium
// Time: O(n^2) | Space: O(n)

import "fmt"

func longestBalancedSubarrayI(nums []int) int {
	n := len(nums)
	ans := 0

	for i := 0; i < n; i++ {
		if n-i <= ans {
			break
		}
  // Membuat map (HashMap) — pencarian O(1)
		evenVisited := make(map[int]bool)
  // Membuat map (HashMap) — pencarian O(1)
		oddVisited := make(map[int]bool)
		evenCount := 0
		oddCount := 0

		for j := i; j < n; j++ {
			if nums[j]%2 == 0 {
				if !evenVisited[nums[j]] {
					evenVisited[nums[j]] = true
					evenCount++
				}
			} else {
				if !oddVisited[nums[j]] {
					oddVisited[nums[j]] = true
					oddCount++
				}
			}
			if evenCount == oddCount {
				if j-i+1 > ans {
					ans = j - i + 1
				}
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(longestBalancedSubarrayI([]int{1, 2, 3, 4}))
	fmt.Println(longestBalancedSubarrayI([]int{2, 4, 6, 8}))
	fmt.Println(longestBalancedSubarrayI([]int{1, 1, 2, 2, 3, 3}))
}
```
