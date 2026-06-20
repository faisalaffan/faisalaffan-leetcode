# 1115 — Print Foobar Alternately

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #1115: Print FooBar Alternately
// https://leetcode.com/problems/print-foobar-alternately/
// Difficulty: Medium
//
// Approach: Two goroutines synchronized with channels
// Time: O(n)
// Space: O(1)

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	fooCh := make(chan bool, 1)
	barCh := make(chan bool, 1)
	fooCh <- true

	go func() {
		defer wg.Done()
		for i := 0; i < 3; i++ {
			<-fooCh
			fmt.Print("foo")
			barCh <- true
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 3; i++ {
			<-barCh
			fmt.Print("bar")
			fooCh <- true
		}
	}()

	wg.Wait()
	fmt.Println()
}
```
