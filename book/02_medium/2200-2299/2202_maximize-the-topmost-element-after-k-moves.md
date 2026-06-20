# 2202 — Maximize The Topmost Element After K Moves

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func maximumTop(nums []int, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2202: Maximize the Topmost Element After K Moves
// https://leetcode.com/problems/maximize-the-topmost-element-after-k-moves/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func maximumTop(nums []int, k int) int {
	n := len(nums)
	if n == 1 {
		if k%2 == 1 {
			return -1
		}
		return nums[0]
	}

	if k == 0 {
		return nums[0]
	}
	if k == 1 {
		return nums[1]
	}

	maxVal := -1
	for i := 0; i < n && i < k-1; i++ {
		if nums[i] > maxVal {
			maxVal = nums[i]
		}
	}
	if k < n {
		if nums[k] > maxVal {
			maxVal = nums[k]
		}
	}
	return maxVal
}

func main() {
	// Test case 1
	fmt.Println(maximumTop([]int{5, 2, 4, 3, 1}, 3))
	// Expected: 5

	// Test case 2
	fmt.Println(maximumTop([]int{2}, 1))
	// Expected: -1

	// Test case 3
	fmt.Println(maximumTop([]int{99, 95, 68, 24, 18}, 69))
	// Expected: 99
}
```
