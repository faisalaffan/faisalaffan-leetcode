# 1832 — Check If The Sentence Is Pangram

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func CheckIfPangram(sentence string) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Bitmask

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Bitmask** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1832: Check if the Sentence Is Pangram
// https://leetcode.com/problems/check-if-the-sentence-is-pangram/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(1)
func CheckIfPangram(sentence string) bool {
	seen := 0
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(sentence); i++ {
		seen |= 1 << (sentence[i] - 'a')
	}
	return seen == (1<<26)-1
}

func main() {
	fmt.Println(CheckIfPangram("thequickbrownfoxjumpsoverthelazydog"))
	fmt.Println(CheckIfPangram("leetcode"))
}
```
