# 2621 — Sleep

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func sleep(millis int) 
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(millis)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2621: Sleep
// https://leetcode.com/problems/sleep/
// Difficulty: Easy
// Time: O(millis) | Space: O(1)
// Note: JavaScript async problem, adapted to Go.

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	sleep(100)
	fmt.Println("Slept for", time.Since(start).Milliseconds(), "ms")
}

func sleep(millis int) {
	time.Sleep(time.Duration(millis) * time.Millisecond)
}
```
