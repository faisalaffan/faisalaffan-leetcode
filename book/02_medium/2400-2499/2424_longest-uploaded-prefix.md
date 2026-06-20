# 2424 — Longest Uploaded Prefix

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func Constructor(n int) LUPrefix
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Prefix Sum

**Kompleksitas Waktu:** O(1) amortized  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Prefix Sum** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2424: Longest Uploaded Prefix
// https://leetcode.com/problems/longest-uploaded-prefix/
// Difficulty: Medium
// Time: O(1) amortized | Space: O(n)
// Track uploaded videos. Longest prefix = longest 1..k where all uploaded.

import "fmt"

type LUPrefix struct {
	uploaded []bool
	longest  int
}

func main() {
	lu := Constructor(4)
	fmt.Println(lu.Longest()) // 0
	lu.Upload(3)
	fmt.Println(lu.Longest()) // 0
	lu.Upload(1)
	fmt.Println(lu.Longest()) // 1
	lu.Upload(2)
	fmt.Println(lu.Longest()) // 3
}

func Constructor(n int) LUPrefix {
	return LUPrefix{uploaded: make([]bool, n+2)}
}

func (l *LUPrefix) Upload(video int) {
	l.uploaded[video] = true
	for l.uploaded[l.longest+1] {
		l.longest++
	}
}

func (l *LUPrefix) Longest() int {
	return l.longest
}
```
