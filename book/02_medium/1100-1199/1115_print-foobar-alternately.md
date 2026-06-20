# 1115 — Print Foobar Alternately

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

> **💡 Hint:** Two goroutines synchronized with channels

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

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
