# 2715 — Timeout Cancellation

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func TimeoutCancellation(fn func() , delay int) func()
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2715: Timeout Cancellation
// https://leetcode.com/problems/timeout-cancellation/
// Difficulty: Easy
// Time: O(1) | Space: O(1)
// Note: JavaScript problem, adapted to Go. Returns a cancel function.

import (
	"fmt"
	"time"
)

func main() {
	cancel := TimeoutCancellation(func() { fmt.Println("executed") }, 100)
	time.Sleep(50 * time.Millisecond)
	cancel()
	time.Sleep(100 * time.Millisecond)
	fmt.Println("done")
}

func TimeoutCancellation(fn func(), delay int) func() {
	timer := time.AfterFunc(time.Duration(delay)*time.Millisecond, fn)
	return func() {
		timer.Stop()
	}
}
```
