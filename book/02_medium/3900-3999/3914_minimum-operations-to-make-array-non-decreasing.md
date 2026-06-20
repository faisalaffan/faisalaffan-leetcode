# 3914 — Minimum Operations To Make Array Non Decreasing

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MinimumOperationsToMakeArrayNonDecreasing(nums []int) int64
```

> **💡 Hint:** Greedy. When nums[i] < nums[i-1], need to increase a suffix.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(N)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3914: Minimum Operations to Make Array Non Decreasing
// https://leetcode.com/problems/minimum-operations-to-make-array-non-decreasing/
// Difficulty: Medium
// Time: O(N) | Space: O(1)
// Approach: Greedy. When nums[i] < nums[i-1], need to increase a suffix.
// Track operations with running max.

import "fmt"

func MinimumOperationsToMakeArrayNonDecreasing(nums []int) int64 {
	var ans int64 = 0
	mx := 0
	for _, v := range nums {
		if v >= mx {
			mx = v
		} else {
			ans += int64(mx - v)
		}
	}
	return ans
}

func main() {
	// Example
	fmt.Println(MinimumOperationsToMakeArrayNonDecreasing([]int{1, 2, 3}))        // Expected: 0
	fmt.Println(MinimumOperationsToMakeArrayNonDecreasing([]int{3, 2, 1}))        // Expected: 3
	fmt.Println(MinimumOperationsToMakeArrayNonDecreasing([]int{1, 2, 1, 2, 1})) // Expected: 2
}
```
