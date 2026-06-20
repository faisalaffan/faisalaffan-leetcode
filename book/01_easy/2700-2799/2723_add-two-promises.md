# 2723 — Add Two Promises

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func AddTwoPromises(promise1 func() int, promise2 func() int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2723: Add Two Promises
// https://leetcode.com/problems/add-two-promises/
// Difficulty: Easy
// Time: O(1) | Space: O(1)
// Note: JavaScript Promise problem, adapted to Go. Adds two integer results asynchronously.

import (
	"fmt"
	"time"
)

func main() {
	p1 := func() int { time.Sleep(50 * time.Millisecond); return 10 }
	p2 := func() int { time.Sleep(100 * time.Millisecond); return 20 }
	result := AddTwoPromises(p1, p2)
	fmt.Println(result)
}

func AddTwoPromises(promise1 func() int, promise2 func() int) int {
	result1 := make(chan int)
	result2 := make(chan int)

	go func() { result1 <- promise1() }()
	go func() { result2 <- promise2() }()

	return <-result1 + <-result2
}
```
