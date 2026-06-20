# 2490 — Circular Sentence

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func CircularSentence(sentence string) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2490: Circular Sentence
// https://leetcode.com/problems/circular-sentence/
// Difficulty: Easy
// Time O(n) | Space O(1)

import "fmt"

func main() {
	fmt.Println(CircularSentence("leetcode exercises sound delightful")) // true
	fmt.Println(CircularSentence("eetcode"))                            // true
	fmt.Println(CircularSentence("Leetcode is cool"))                   // false
}

func CircularSentence(sentence string) bool {
	if sentence[0] != sentence[len(sentence)-1] {
		return false
	}
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(sentence); i++ {
		if sentence[i] == ' ' && sentence[i-1] != sentence[i+1] {
			return false
		}
	}
	return true
}
```
