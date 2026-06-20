# 2637 — Promise Time Limit

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func timeLimit(fn func(int) int, limit time.Duration) func(int) (int, bool)
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2637: Promise Time Limit
// https://leetcode.com/problems/promise-time-limit/
// Difficulty: Medium
// Time: O(1) | Space: O(1)

import (
	"fmt"
	"sync"
	"time"
)

type TimeLimitedFn struct {
	fn func(int) int
}

type Result struct {
	val int
	err error
}

func timeLimit(fn func(int) int, limit time.Duration) func(int) (int, bool) {
	return func(arg int) (int, bool) {
		var mu sync.Mutex
		done := false
		result := 0

		go func() {
			r := fn(arg)
			mu.Lock()
			if !done {
				result = r
				done = true
			}
			mu.Unlock()
		}()

		time.Sleep(limit)

		mu.Lock()
		defer mu.Unlock()
		if !done {
			done = true
			return 0, false
		}
		return result, true
	}
}

func main() {
	// Test case 1: completes in time
	fastFn := func(x int) int {
		time.Sleep(10 * time.Millisecond)
		return x * 2
	}
	limited := timeLimit(fastFn, 100*time.Millisecond)
	val, ok := limited(5)
	fmt.Println("Test 1:", val, ok)
	// Expected: 10 true

	// Test case 2: times out
	slowFn := func(x int) int {
		time.Sleep(200 * time.Millisecond)
		return x
	}
	limited2 := timeLimit(slowFn, 50*time.Millisecond)
	val2, ok2 := limited2(5)
	fmt.Println("Test 2:", val2, ok2)
	// Expected: 0 false
}
```
