# 2623 — Memoize

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func memoize(fn func(...int) int) MemoizedFn
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(1) amortized per call  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2623: Memoize
// https://leetcode.com/problems/memoize/
// Difficulty: Medium
// Time: O(1) amortized per call | Space: O(n)

import (
	"fmt"
	"sync"
)

type MemoizedFn func(args ...int) int

func memoize(fn func(...int) int) MemoizedFn {
	var mu sync.Mutex
  // Membuat map (HashMap) — pencarian O(1)
	cache := make(map[string]int)

	return func(args ...int) int {
		// Create a key from args
		key := ""
		for i, arg := range args {
			if i > 0 {
				key += ","
			}
			key += fmt.Sprintf("%d", arg)
		}

		mu.Lock()
		defer mu.Unlock()

		if val, ok := cache[key]; ok {
			return val
		}
		result := fn(args...)
		cache[key] = result
		return result
	}
}

var callCount int

func add(a, b int) int {
	callCount++
	return a + b
}

func main() {
	callCount = 0
	memoizedAdd := memoize(func(args ...int) int {
		return add(args[0], args[1])
	})

	// Test case 1
	fmt.Println("Test 1:", memoizedAdd(1, 2))
	// Expected: 3

	// Test case 2: same args, should use cache
	fmt.Println("Test 2:", memoizedAdd(1, 2))
	// Expected: 3

	// Test case 3: different args
	fmt.Println("Test 3:", memoizedAdd(2, 3))
	// Expected: 5

	fmt.Println("Calls:", callCount)
	// Expected: 2 (not 3)
}
```
