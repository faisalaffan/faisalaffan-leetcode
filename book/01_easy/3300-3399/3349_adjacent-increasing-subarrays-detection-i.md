# 3349 — Adjacent Increasing Subarrays Detection I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func AdjacentIncreasingSubarraysDetectionI(nums []int, k int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n). Space: O(n).  
**Kompleksitas Ruang:** O(n).

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3349: Adjacent Increasing Subarrays Detection I
// https://leetcode.com/problems/adjacent-increasing-subarrays-detection-i/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(AdjacentIncreasingSubarraysDetectionI([]int{2, 5, 7, 8, 9, 2, 3, 4, 3, 1}, 3))
	fmt.Println(AdjacentIncreasingSubarraysDetectionI([]int{1, 2, 3, 4, 4, 4, 4, 5, 6, 7}, 5))
}

// AdjacentIncreasingSubarraysDetectionI returns true if there exist two adjacent k-length increasing subarrays.
// Time: O(n). Space: O(n).
func AdjacentIncreasingSubarraysDetectionI(nums []int, k int) bool {
	n := len(nums)
	if n < 2*k {
		return false
	}

	// inc[i] = true if subarray starting at i of length k is strictly increasing
	inc := make([]bool, n-k+1)
	for i := 0; i <= n-k; i++ {
		isInc := true
		for j := i; j < i+k-1; j++ {
			if nums[j] >= nums[j+1] {
				isInc = false
				break
			}
		}
		inc[i] = isInc
	}

	for i := 0; i <= n-2*k; i++ {
		if inc[i] && inc[i+k] {
			return true
		}
	}
	return false
}
```
