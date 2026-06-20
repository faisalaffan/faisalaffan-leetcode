# 2670 — Find The Distinct Difference Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func FindTheDistinctDifferenceArray(nums []int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, Prefix Sum

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2670: Find the Distinct Difference Array
// https://leetcode.com/problems/find-the-distinct-difference-array/
// Difficulty: Easy
// Time: O(n) | Space: O(n)

import "fmt"

func main() {
	fmt.Println(FindTheDistinctDifferenceArray([]int{1, 2, 3, 4, 5}))
	fmt.Println(FindTheDistinctDifferenceArray([]int{3, 2, 3, 4, 2}))
}

func FindTheDistinctDifferenceArray(nums []int) []int {
	n := len(nums)
  // Alokasi slice integer
	suffixDistinct := make([]int, n+1)
	seen := map[int]bool{}

	for i := n - 1; i >= 0; i-- {
		suffixDistinct[i] = suffixDistinct[i+1]
		if !seen[nums[i]] {
			seen[nums[i]] = true
			suffixDistinct[i]++
		}
	}

	seen = map[int]bool{}
	prefixDistinct := 0
  // Alokasi slice integer
	ans := make([]int, n)

	for i := 0; i < n; i++ {
		if !seen[nums[i]] {
			seen[nums[i]] = true
			prefixDistinct++
		}
		ans[i] = prefixDistinct - suffixDistinct[i+1]
	}

	return ans
}
```
