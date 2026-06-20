# 1116 — Print Zero Even Odd

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

> **💡 Hint:** Three goroutines with channels for ordering

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1116: Print Zero Even Odd
// https://leetcode.com/problems/print-zero-even-odd/
// Difficulty: Medium
//
// Approach: Three goroutines with channels for ordering
// Time: O(n)
// Space: O(1)

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(3)

	zeroCh := make(chan bool, 1)
	evenCh := make(chan bool, 1)
	oddCh := make(chan bool, 1)
	zeroCh <- true

	n := 5
	go func() {
		defer wg.Done()
		for i := 1; i <= n; i++ {
			<-zeroCh
			fmt.Print(0)
			if i%2 == 0 {
				evenCh <- true
			} else {
				oddCh <- true
			}
		}
	}()

	go func() {
		defer wg.Done()
		for i := 2; i <= n; i += 2 {
			<-evenCh
			fmt.Print(i)
			zeroCh <- true
		}
	}()

	go func() {
		defer wg.Done()
		for i := 1; i <= n; i += 2 {
			<-oddCh
			fmt.Print(i)
			zeroCh <- true
		}
	}()

	wg.Wait()
	fmt.Println()
}
```
