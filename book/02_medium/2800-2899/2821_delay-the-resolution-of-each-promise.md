# 2821 — Delay The Resolution Of Each Promise

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func DelayTheResolutionOfEachPromise(functions []func() int, delay time.Duration) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2821: Delay the Resolution of Each Promise
// https://leetcode.com/problems/delay-the-resolution-of-each-promise/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(1)

import (
	"fmt"
	"sync"
	"time"
)

func DelayTheResolutionOfEachPromise(functions []func() int, delay time.Duration) []int {
	n := len(functions)
  // Alokasi slice integer
	results := make([]int, n)
	var wg sync.WaitGroup

	for i, fn := range functions {
		wg.Add(1)
		go func(idx int, f func() int) {
			defer wg.Done()
			time.Sleep(delay)
			results[idx] = f()
		}(i, fn)
	}

	wg.Wait()
	return results
}

func main() {
	start := time.Now()
	results := DelayTheResolutionOfEachPromise([]func() int{
		func() int { return 10 },
		func() int { return 20 },
		func() int { return 30 },
	}, 5*time.Millisecond)
	fmt.Println(results)
	fmt.Println("Took:", time.Since(start).Round(time.Millisecond))
}
```
