# 1760 — Minimum Limit Of Balls In A Bag

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func minimumSize(nums []int, maxOperations int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log M) where M = max(nums), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1760: Minimum Limit of Balls in a Bag
// https://leetcode.com/problems/minimum-limit-of-balls-in-a-bag/
// Difficulty: Medium
// Time: O(n log M) where M = max(nums), Space: O(1)

import "fmt"

func minimumSize(nums []int, maxOperations int) int {
	left, right := 1, 0
	for _, v := range nums {
		if v > right {
			right = v
		}
	}

  // Two-pointer: gerakkan kiri atau kanan
	for left < right {
		mid := left + (right-left)/2
		if canDivide(nums, maxOperations, mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func canDivide(nums []int, maxOps, limit int) bool {
	ops := 0
	for _, v := range nums {
		if v > limit {
			ops += (v - 1) / limit
			if ops > maxOps {
				return false
			}
		}
	}
	return true
}

func main() {
	fmt.Println(minimumSize([]int{9}, 2))               // Expected: 3
	fmt.Println(minimumSize([]int{2, 4, 8, 2}, 4))      // Expected: 2
	fmt.Println(minimumSize([]int{7, 17}, 2))            // Expected: 7
}
```
