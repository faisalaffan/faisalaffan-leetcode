# 0276 — Paint Fence

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func numWays(n int, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #276: Paint Fence
// https://leetcode.com/problems/paint-fence/
// Difficulty: Medium [Paid]
// Time: O(n), Space: O(1)

import "fmt"

func numWays(n int, k int) int {
	if n == 0 || k == 0 {
		return 0
	}
	if n == 1 {
		return k
	}

	same := k
	diff := k * (k - 1)

	for i := 3; i <= n; i++ {
		same, diff = diff, (same+diff)*(k-1)
	}

	return same + diff
}

func main() {
	fmt.Println(numWays(3, 2))
	fmt.Println(numWays(1, 1))
	fmt.Println(numWays(7, 2))
}
```
