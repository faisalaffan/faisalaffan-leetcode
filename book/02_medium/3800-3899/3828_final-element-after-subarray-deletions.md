# 3828 — Final Element After Subarray Deletions

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func FinalElementAfterSubarrayDeletions(nums []int) int
```

> **💡 Hint:** Alice (first) can always keep first or last. Optimal play yields max(first, last).

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3828: Final Element After Subarray Deletions
// https://leetcode.com/problems/final-element-after-subarray-deletions/
// Difficulty: Medium
// Time: O(1) | Space: O(1)
// Approach: Alice (first) can always keep first or last. Optimal play yields max(first, last).

import "fmt"

func FinalElementAfterSubarrayDeletions(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if nums[0] > nums[len(nums)-1] {
		return nums[0]
	}
	return nums[len(nums)-1]
}

func main() {
	// Example 1
	fmt.Println(FinalElementAfterSubarrayDeletions([]int{1, 5, 2})) // Expected: 2

	// Example 2
	fmt.Println(FinalElementAfterSubarrayDeletions([]int{3, 7})) // Expected: 7

	// Example 3
	fmt.Println(FinalElementAfterSubarrayDeletions([]int{5})) // Expected: 5
}
```
