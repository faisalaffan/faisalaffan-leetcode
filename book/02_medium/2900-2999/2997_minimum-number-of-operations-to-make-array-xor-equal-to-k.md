# 2997 — Minimum Number Of Operations To Make Array Xor Equal To K

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func minOperationsXOR(nums []int, k int) (ans int)
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2997: Minimum Number of Operations to Make Array XOR Equal to K
// https://leetcode.com/problems/minimum-number-of-operations-to-make-array-xor-equal-to-k/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Println(minOperationsXOR([]int{2, 1, 3, 4}, 1))
	fmt.Println(minOperationsXOR([]int{2, 0, 2, 0}, 0))
}

func minOperationsXOR(nums []int, k int) (ans int) {
	xor := 0
	for _, x := range nums {
		xor ^= x
	}
	return bits.OnesCount(uint(xor ^ k))
}
```
