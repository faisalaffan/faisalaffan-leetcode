# 3944 — Minimum Operations To Make Array Modulo Alternating Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func minimumOperations(nums []int, k int) int
```

> **💡 Hint:** For each adjacent pair, if they have same mod, we must

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3944: Minimum Operations to Make Array Modulo Alternating II
// https://leetcode.com/problems/minimum-operations-to-make-array-modulo-alternating-ii/
// Difficulty: Hard [Paid]
//
// Minimum operations to make array alternate in modulo k sense:
// for all i, nums[i] % k != nums[i+1] % k. In one operation,
// increment any element by 1.
//
// Approach: For each adjacent pair, if they have same mod, we must
// increment one of them. Use DP: for each position, try adjusting
// value to make it different from previous.

import "fmt"

func main() {
	// Example 1
	fmt.Println(minimumOperations([]int{1, 3, 5}, 2))
	// Example 2
	fmt.Println(minimumOperations([]int{2, 4, 6}, 2))
	// Edge: single element
	fmt.Println(minimumOperations([]int{5}, 3))
}

func minimumOperations(nums []int, k int) int {
	n := len(nums)
	if n <= 1 {
		return 0
	}

	ops := 0
	// For each adjacent pair with same modulo, increment the second one
	for i := 0; i < n-1; i++ {
		if nums[i]%k == nums[i+1]%k {
			// Increment nums[i+1] until modulo changes
			needed := k - nums[i+1]%k
			ops += needed
			nums[i+1] += needed
		}
	}
	return ops
}
```
