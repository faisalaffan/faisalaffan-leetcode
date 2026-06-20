# 3732 — Maximum Product Of Three Elements After One Replacement

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func maximumProductOfThreeElementsAfterOneReplacement(nums []int) int64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3732: Maximum Product of Three Elements After One Replacement
// https://leetcode.com/problems/maximum-product-of-three-elements-after-one-replacement/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func maximumProductOfThreeElementsAfterOneReplacement(nums []int) int64 {
	var first, second int64 = 0, 0
	for _, v := range nums {
		val := int64(v)
		if val < 0 {
			val = -val
		}
		if val > first {
			second = first
			first = val
		} else if val > second {
			second = val
		}
	}
	return 100000 * first * second
}

func main() {
	fmt.Println(maximumProductOfThreeElementsAfterOneReplacement([]int{1, 2, 3, 4}))
	fmt.Println(maximumProductOfThreeElementsAfterOneReplacement([]int{-4, -2, -1, -3}))
	fmt.Println(maximumProductOfThreeElementsAfterOneReplacement([]int{-5, 7, 0}))
}
```
