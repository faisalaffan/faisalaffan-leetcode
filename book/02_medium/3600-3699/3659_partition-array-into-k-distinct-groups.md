# 3659 — Partition Array Into K Distinct Groups

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func partitionArrayIntoKDistinctGroups(nums []int, k int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(max(nums))

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3659: Partition Array Into K-Distinct Groups
// https://leetcode.com/problems/partition-array-into-k-distinct-groups/
// Difficulty: Medium
// Time: O(n) | Space: O(max(nums))

import (
	"fmt"
	"slices"
)

func partitionArrayIntoKDistinctGroups(nums []int, k int) bool {
	n := len(nums)
	if n%k != 0 {
		return false
	}

	maxVal := slices.Max(nums)
  // Alokasi slice integer
	cnt := make([]int, maxVal+1)
	for _, x := range nums {
		cnt[x]++
		if cnt[x] > n/k {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(partitionArrayIntoKDistinctGroups([]int{1, 2, 3, 4}, 2))
	fmt.Println(partitionArrayIntoKDistinctGroups([]int{1, 1, 1, 1}, 2))
	fmt.Println(partitionArrayIntoKDistinctGroups([]int{1, 2, 2, 3, 3, 4}, 3))
}
```
