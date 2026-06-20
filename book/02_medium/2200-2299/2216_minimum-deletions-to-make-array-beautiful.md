# 2216 — Minimum Deletions To Make Array Beautiful

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func minDeletion(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2216: Minimum Deletions to Make Array Beautiful
// https://leetcode.com/problems/minimum-deletions-to-make-array-beautiful/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func minDeletion(nums []int) int {
	n := len(nums)
	deletions := 0
	i := 0

	for i < n-1 {
		idx := i - deletions
		if idx%2 == 0 && nums[i] == nums[i+1] {
			deletions++
			i++
		} else {
			i++
		}
	}

	// If after deletions the array length is odd, delete last element
	if (n-deletions)%2 == 1 {
		deletions++
	}
	return deletions
}

func main() {
	// Test case 1
	fmt.Println(minDeletion([]int{1, 1, 2, 3, 5}))
	// Expected: 1

	// Test case 2
	fmt.Println(minDeletion([]int{1, 1, 2, 2, 3, 3}))
	// Expected: 2
}
```
