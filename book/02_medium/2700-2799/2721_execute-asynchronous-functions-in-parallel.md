# 2721 — Execute Asynchronous Functions In Parallel

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func ExecuteAsynchronousFunctionsInParallel(functions []func() int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2721: Execute Asynchronous Functions in Parallel
// https://leetcode.com/problems/execute-asynchronous-functions-in-parallel/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import (
	"fmt"
	"sync"
)

func ExecuteAsynchronousFunctionsInParallel(functions []func() int) []int {
	var wg sync.WaitGroup
  // Alokasi slice integer
	results := make([]int, len(functions))

	for i, fn := range functions {
		wg.Add(1)
		go func(idx int, f func() int) {
			defer wg.Done()
			results[idx] = f()
		}(i, fn)
	}

	wg.Wait()
	return results
}

func main() {
	results := ExecuteAsynchronousFunctionsInParallel([]func() int{
		func() int { return 1 },
		func() int { return 2 },
		func() int { return 3 },
	})
	fmt.Println(results)

	results2 := ExecuteAsynchronousFunctionsInParallel([]func() int{
		func() int { return 42 },
	})
	fmt.Println(results2)
}
```
