# 3405 — Count The Number Of Arrays With K Matching Adjacent Elements

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func powMod(a, b int64) int64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3405: Count the Number of Arrays with K Matching Adjacent Elements
// https://leetcode.com/problems/count-the-number-of-arrays-with-k-matching-adjacent-elements/
// Difficulty: Hard
//
// Combinatorics: C(n-1, k) * m * (m-1)^(n-1-k) mod MOD.

import "fmt"

func main() {
	fmt.Println(CountTheNumberOfArraysWithKMatchingAdjacentElements(4, 2, 2))
}

const MOD3405 = 1000000007

func powMod(a, b int64) int64 {
	res := int64(1)
	for b > 0 {
		if b&1 == 1 {
			res = res * a % MOD3405
		}
		a = a * a % MOD3405
		b >>= 1
	}
	return res
}

func CountTheNumberOfArraysWithKMatchingAdjacentElements(n, m, k int) int {
	if k >= n {
		return 0
	}

	// Precompute factorials
	size := n
  // Alokasi slice integer
	fact := make([]int64, size+1)
	fact[0] = 1
	for i := 1; i <= size; i++ {
		fact[i] = fact[i-1] * int64(i) % MOD3405
	}
  // Alokasi slice integer
	invFact := make([]int64, size+1)
	invFact[size] = powMod(fact[size], MOD3405-2)
	for i := size - 1; i >= 0; i-- {
		invFact[i] = invFact[i+1] * int64(i+1) % MOD3405
	}

	nCr := fact[n-1] * invFact[k] % MOD3405 * invFact[n-1-k] % MOD3405
	ans := nCr * int64(m) % MOD3405
	if n-1-k > 0 {
		ans = ans * powMod(int64(m-1), int64(n-1-k)) % MOD3405
	}
	return int(ans)
}
```
