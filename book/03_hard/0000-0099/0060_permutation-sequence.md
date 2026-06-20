# 0060 — Permutation Sequence

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func getPermutation(n int, k int) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #60: Permutation Sequence
// https://leetcode.com/problems/permutation-sequence/
// Difficulty: Hard

import "fmt"

func main() {
	fmt.Println("60. Permutation Sequence")
	fmt.Println("n=3, k=3:", getPermutation(3, 3), "(expected 213)")
	fmt.Println("n=4, k=9:", getPermutation(4, 9), "(expected 2314)")
	fmt.Println("n=3, k=1:", getPermutation(3, 1), "(expected 123)")
}

func getPermutation(n int, k int) string {
	fact := 1
	nums := make([]byte, 0, n)
	for i := 1; i <= n; i++ {
		fact *= i
		nums = append(nums, byte('0'+i))
	}

	k-- // convert to 0-indexed
	result := make([]byte, 0, n)

	for i := 0; i < n; i++ {
		fact /= (n - i)
		idx := k / fact
		result = append(result, nums[idx])
		nums = append(nums[:idx], nums[idx+1:]...)
		k %= fact
	}

	return string(result)
}
```
