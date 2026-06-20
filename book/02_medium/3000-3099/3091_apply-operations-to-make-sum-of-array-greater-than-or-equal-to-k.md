# 3091 — Apply Operations To Make Sum Of Array Greater Than Or Equal To K

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func minOperations(k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(sqrt(k))  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3091: Apply Operations to Make Sum of Array Greater Than or Equal to k
// https://leetcode.com/problems/apply-operations-to-make-sum-of-array-greater-than-or-equal-to-k/
// Difficulty: Medium
// Time: O(sqrt(k)) | Space: O(1)

import "fmt"

func minOperations(k int) int {
	ans := k - 1
	for a := 1; a <= k; a++ {
		b := (k + a - 1) / a
		ops := (a - 1) + (b - 1)
		if ops < ans {
			ans = ops
		}
	}
	return ans
}

func main() {
	fmt.Println(minOperations(11))  // Expected: 5
	fmt.Println(minOperations(1))   // Expected: 0
	fmt.Println(minOperations(5))   // Expected: 3
}
```
