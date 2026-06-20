# 2970 — Count The Number Of Incremovable Subarrays I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func CountTheNumberOfIncremovableSubarraysI(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n^3)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2970: Count the Number of Incremovable Subarrays I
// https://leetcode.com/problems/count-the-number-of-incremovable-subarrays-i/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: incremovableSubarrayCount
	fmt.Println(CountTheNumberOfIncremovableSubarraysI([]int{1, 2, 3, 4})) // 10
	fmt.Println(CountTheNumberOfIncremovableSubarraysI([]int{6, 5, 7, 8})) // 7
	fmt.Println(CountTheNumberOfIncremovableSubarraysI([]int{8, 7, 6, 6})) // 3
}

// Time: O(n^3) | Space: O(n)
// LeetCode submission name: incremovableSubarrayCount
func CountTheNumberOfIncremovableSubarraysI(nums []int) int {
	n := len(nums)
	count := 0

	for l := 0; l < n; l++ {
		for r := l; r < n; r++ {
			// Check if array without nums[l..r] is strictly increasing
			if isStrictlyIncreasing(nums, l, r) {
				count++
			}
		}
	}
	return count
}

func isStrictlyIncreasing(nums []int, l, r int) bool {
	prev := -1
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(nums); i++ {
		if i >= l && i <= r {
			continue
		}
		if nums[i] <= prev {
			return false
		}
		prev = nums[i]
	}
	return true
}
```
