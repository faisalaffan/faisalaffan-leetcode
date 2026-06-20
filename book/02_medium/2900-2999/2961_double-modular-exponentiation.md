# 2961 — Double Modular Exponentiation

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func getGoodIndices(variables [][]int, target int) (ans []int)
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log m)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2961: Double Modular Exponentiation
// https://leetcode.com/problems/double-modular-exponentiation/
// Difficulty: Medium
// Time: O(n log m) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(getGoodIndices([][]int{{2, 3, 3, 10}, {3, 3, 3, 1}, {6, 1, 1, 4}}, 2))
	fmt.Println(getGoodIndices([][]int{{39, 3, 1000, 1000}}, 17))
}

func getGoodIndices(variables [][]int, target int) (ans []int) {
	qpow := func(a, n, mod int) int {
		ans := 1
		for ; n > 0; n >>= 1 {
			if n&1 == 1 {
				ans = ans * a % mod
			}
			a = a * a % mod
		}
		return ans
	}
	for i, e := range variables {
		a, b, c, m := e[0], e[1], e[2], e[3]
		if qpow(qpow(a, b, 10), c, m) == target {
			ans = append(ans, i)
		}
	}
	return
}
```
