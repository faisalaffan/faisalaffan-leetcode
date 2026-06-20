# 3229 — Minimum Operations To Make Array Equal To Target

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func minimumOperations(nums []int, target []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3229: Minimum Operations to Make Array Equal to Target
// https://leetcode.com/problems/minimum-operations-to-make-array-equal-to-target/
// Difficulty: Hard
//
// Greedy on diff array. For each position, if diff[i] and diff[i-1] have the
// same sign, we only pay the extra beyond |diff[i-1]|; otherwise we pay |diff[i]|.

import "fmt"

func main() {
	// Example 1: nums=[3,5,1,2], target=[4,6,2,4] => 2
	fmt.Println(minimumOperations([]int{3, 5, 1, 2}, []int{4, 6, 2, 4}))
	// Example 2: all same
	fmt.Println(minimumOperations([]int{1, 2, 3}, []int{1, 2, 3}))
	// Example 3: mixed signs
	fmt.Println(minimumOperations([]int{1, 2, 3}, []int{3, 2, 1}))
	// Example 4: single element
	fmt.Println(minimumOperations([]int{0}, []int{10}))
	// Example 5: alternating
	fmt.Println(minimumOperations([]int{0, 0, 0}, []int{1, -1, 1}))
}

func minimumOperations(nums []int, target []int) int {
	n := len(nums)
  // Edge case: input kosong — langsung return
	if n == 0 {
		return 0
	}

	var ans int
	var prev int

	for i := 0; i < n; i++ {
		curr := target[i] - nums[i]
		if i == 0 {
			ans += abs(curr)
		} else if curr*prev > 0 {
			// Same sign: only pay the increase beyond previous
			absCurr := abs(curr)
			absPrev := abs(prev)
			if absCurr > absPrev {
				ans += absCurr - absPrev
			}
		} else {
			// Different sign or zero: pay full amount
			ans += abs(curr)
		}
		prev = curr
	}

	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```
