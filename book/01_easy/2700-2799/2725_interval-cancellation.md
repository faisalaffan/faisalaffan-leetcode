# 2725 — Interval Cancellation

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func IntervalCancellation(fn func() , delay int) func()
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2725: Interval Cancellation
// https://leetcode.com/problems/interval-cancellation/
// Difficulty: Easy
// Time: O(1) | Space: O(1)
// Note: JavaScript setInterval problem, adapted to Go.

import (
	"fmt"
	"time"
)

func main() {
	count := 0
	cancel := IntervalCancellation(func() {
		count++
		fmt.Println("tick", count)
	}, 50)
	time.Sleep(200 * time.Millisecond)
	cancel()
	time.Sleep(100 * time.Millisecond)
	fmt.Println("total ticks:", count)
}

func IntervalCancellation(fn func(), delay int) func() {
	ticker := time.NewTicker(time.Duration(delay) * time.Millisecond)
	done := make(chan struct{})

	go func() {
		for {
			select {
			case <-ticker.C:
				fn()
			case <-done:
				ticker.Stop()
				return
			}
		}
	}()

	return func() {
		close(done)
	}
}
```
