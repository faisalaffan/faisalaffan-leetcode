# 2938 — Separate Black And White Balls

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func minimumSteps(s string) (ans int64)
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2938: Separate Black and White Balls
// https://leetcode.com/problems/separate-black-and-white-balls/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(minimumSteps("101"))
	fmt.Println(minimumSteps("100"))
	fmt.Println(minimumSteps("0111"))
}

func minimumSteps(s string) (ans int64) {
	n := len(s)
	cnt := 0
	for i := n - 1; i >= 0; i-- {
		if s[i] == '1' {
			cnt++
			ans += int64(n - i - cnt)
		}
	}
	return
}
```
