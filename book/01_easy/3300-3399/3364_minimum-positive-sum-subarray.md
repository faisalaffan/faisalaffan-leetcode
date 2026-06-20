# 3364 — Minimum Positive Sum Subarray

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MinimumPositiveSumSubarray(nums []int, l int, r int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n * (r-l+1)). Space: O(1).  
**Kompleksitas Ruang:** O(1).

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3364: Minimum Positive Sum Subarray
// https://leetcode.com/problems/minimum-positive-sum-subarray/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(MinimumPositiveSumSubarray([]int{3, -2, 1, 4}, 2, 3))
	fmt.Println(MinimumPositiveSumSubarray([]int{-2, 2, -2, 2}, 1, 2))
	fmt.Println(MinimumPositiveSumSubarray([]int{1, 2, 3, 4}, 2, 4))
}

// MinimumPositiveSumSubarray returns the minimum positive sum of any subarray with length between l and r.
// Time: O(n * (r-l+1)). Space: O(1).
func MinimumPositiveSumSubarray(nums []int, l int, r int) int {
	n := len(nums)
	minPos := -1
	for length := l; length <= r; length++ {
		for start := 0; start <= n-length; start++ {
			sum := 0
			for i := start; i < start+length; i++ {
				sum += nums[i]
			}
			if sum > 0 && (minPos == -1 || sum < minPos) {
				minPos = sum
			}
		}
	}
	return minPos
}
```
