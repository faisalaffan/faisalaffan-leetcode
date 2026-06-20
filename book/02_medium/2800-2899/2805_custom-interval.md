# 2805 — Custom Interval

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func CustomInterval(fn func() , interval time.Duration, times int)
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2805: Custom Interval
// https://leetcode.com/problems/custom-interval/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(1)

import (
	"fmt"
	"time"
)

func CustomInterval(fn func(), interval time.Duration, times int) {
	for i := 0; i < times; i++ {
		fn()
		time.Sleep(interval)
	}
}

func main() {
	count := 0
	fn := func() {
		count++
		fmt.Println("Executed:", count)
	}
	CustomInterval(fn, 10*time.Millisecond, 3)
	fmt.Println("Done")
}
```
