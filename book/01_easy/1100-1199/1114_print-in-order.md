# 1114 — Print In Order

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func NewFoo() *Foo
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1114: Print in Order
// https://leetcode.com/problems/print-in-order/
// Difficulty: Easy (Concurrency)
// Time: O(1) | Space: O(1)

import (
	"fmt"
	"sync"
)

type Foo struct {
	wg1 sync.WaitGroup
	wg2 sync.WaitGroup
}

func NewFoo() *Foo {
	f := &Foo{}
	f.wg1.Add(1)
	f.wg2.Add(1)
	return f
}

func (f *Foo) first() {
	fmt.Print("first")
	f.wg1.Done()
}

func (f *Foo) second() {
	f.wg1.Wait()
	fmt.Print("second")
	f.wg2.Done()
}

func (f *Foo) third() {
	f.wg2.Wait()
	fmt.Print("third")
}

func main() {
	// Test: run in order 1,2,3
	f := NewFoo()
	var wg sync.WaitGroup
	wg.Add(3)
	go func() { f.first(); wg.Done() }()
	go func() { f.second(); wg.Done() }()
	go func() { f.third(); wg.Done() }()
	wg.Wait()
	fmt.Println()
}
```
