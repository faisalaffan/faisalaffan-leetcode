# 2929 — Distribute Candies Among Children Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func distributeCandies(n int, limit int) int64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2929: Distribute Candies Among Children II
// https://leetcode.com/problems/distribute-candies-among-children-ii/
// Difficulty: Medium
// Time: O(1) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(distributeCandies(5, 2))
	fmt.Println(distributeCandies(3, 3))
	fmt.Println(distributeCandies(10, 3))
}

func distributeCandies(n int, limit int) int64 {
	if n > 3*limit {
		return 0
	}
	ans := int64(n+2) * int64(n+1) / 2
	if n > limit {
		ans -= 3 * int64(n-limit+1) * int64(n-limit) / 2
	}
	if n-2 >= 2*limit {
		ans += 3 * int64(n-2*limit) * int64(n-2*limit+1) / 2
	}
	return ans
}
```
