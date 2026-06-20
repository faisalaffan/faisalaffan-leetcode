# 3682 — Minimum Index Sum Of Common Elements

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func minimumIndexSumOfCommonElements(nums1 []int, nums2 []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n + m)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3682: Minimum Index Sum of Common Elements
// https://leetcode.com/problems/minimum-index-sum-of-common-elements/
// Difficulty: Medium [Paid]
// Time: O(n + m) | Space: O(n)

import "fmt"

func minimumIndexSumOfCommonElements(nums1 []int, nums2 []int) int {
  // Membuat map (HashMap) — pencarian O(1)
	idxMap := make(map[int]int)
	for i, v := range nums2 {
		idxMap[v] = i
	}

	minSum := int(^uint(0) >> 1)
	found := false

	for i, v := range nums1 {
		if j, ok := idxMap[v]; ok {
			sum := i + j
			if !found || sum < minSum {
				minSum = sum
				found = true
			}
		}
	}

	if !found {
		return -1
	}
	return minSum
}

func main() {
	fmt.Println(minimumIndexSumOfCommonElements([]int{3, 2, 1}, []int{1, 3, 1}))
	fmt.Println(minimumIndexSumOfCommonElements([]int{5, 1, 2}, []int{2, 1, 3}))
	fmt.Println(minimumIndexSumOfCommonElements([]int{6, 4}, []int{7, 8}))
}
```
