# 2111 — Minimum Operations To Make The Array K Increasing

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func minOperations(arr []int, k int) int
```

> **💡 Hint:** LIS per mod-k subsequence.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2111: Minimum Operations to Make the Array K-Increasing
// https://leetcode.com/problems/minimum-operations-to-make-the-array-k-increasing/
// Difficulty: Hard
//
// Approach: LIS per mod-k subsequence.
// Split array into k subsequences: arr[i], arr[i+k], arr[i+2k], ...
// For each subsequence, compute length of longest non-decreasing subsequence (LIS).
// Min operations = len(seq) - LIS(seq). Sum over all k subsequences.

import (
	"fmt"
	"sort"
)

func main() {
	// Example from problem statement
	arr1 := []int{5, 4, 3, 2, 1}
	k1 := 1
	fmt.Printf("minOperations(%v, %d) = %d (expected 4)\n", arr1, k1, minOperations(arr1, k1))

	// Additional tests
	arr2 := []int{4, 1, 5, 2, 6, 2}
	k2 := 2
	fmt.Printf("minOperations(%v, %d) = %d (expected 0)\n", arr2, k2, minOperations(arr2, k2))

	arr3 := []int{1, 2, 3, 4, 5, 6}
	k3 := 2
	fmt.Printf("minOperations(%v, %d) = %d\n", arr3, k3, minOperations(arr3, k3))

	arr4 := []int{5, 3, 1, 4, 2}
	k4 := 2
	fmt.Printf("minOperations(%v, %d) = %d\n", arr4, k4, minOperations(arr4, k4))
}

func minOperations(arr []int, k int) int {
	n := len(arr)
	total := 0

	for i := 0; i < k; i++ {
		seq := []int{}
		for j := i; j < n; j += k {
			seq = append(seq, arr[j])
		}
		total += len(seq) - lengthOfLIS(seq)
	}

	return total
}

// lengthOfLIS returns the length of the longest non-decreasing subsequence.
func lengthOfLIS(nums []int) int {
	tails := []int{}
	for _, x := range nums {
		// Find first element > x (since we want non-decreasing, equal values can extend)
		idx := sort.SearchInts(tails, x+1)
		if idx == len(tails) {
			tails = append(tails, x)
		} else {
			tails[idx] = x
		}
	}
	return len(tails)
}
```
