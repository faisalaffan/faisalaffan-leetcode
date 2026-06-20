# 2498 — Frog Jump Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func maxJump(stones []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2498: Frog Jump II
// https://leetcode.com/problems/frog-jump-ii/
// Difficulty: Medium
// Time: O(n) | Space: O(1)
// Frog can jump forward, max distance = max of (stones[i+2] - stones[i]) for i in 0..n-3
// and also stones[1] - stones[0] and stones[n-1] - stones[n-2].

import "fmt"

func main() {
	fmt.Println(maxJump([]int{0, 2, 5, 6, 7})) // 5
	fmt.Println(maxJump([]int{0, 3, 9}))        // 9
}

func maxJump(stones []int) int {
	n := len(stones)
	ans := stones[1] - stones[0]
	if n > 2 {
		ans = stones[n-1] - stones[n-2]
	}
	for i := 2; i < n; i++ {
		diff := stones[i] - stones[i-2]
		if diff > ans {
			ans = diff
		}
	}
	return ans
}
```
