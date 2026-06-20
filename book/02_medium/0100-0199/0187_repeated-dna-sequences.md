# 0187 — Repeated Dna Sequences

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func findRepeatedDnaSequences(s string) []string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #187: Repeated DNA Sequences
// https://leetcode.com/problems/repeated-dna-sequences/
// Difficulty: Medium
// Time: O(n), Space: O(n)

import "fmt"

func findRepeatedDnaSequences(s string) []string {
	if len(s) < 10 {
		return nil
	}

  // Membuat map (HashMap) — pencarian O(1)
	seen := make(map[string]int)
	result := []string{}

	for i := 0; i <= len(s)-10; i++ {
		sub := s[i : i+10]
		seen[sub]++
		if seen[sub] == 2 {
			result = append(result, sub)
		}
	}

	return result
}

func main() {
	fmt.Println(findRepeatedDnaSequences("AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"))
	fmt.Println(findRepeatedDnaSequences("AAAAAAAAAAAAA"))
	fmt.Println(findRepeatedDnaSequences("AAAAAAAAAAA"))
}
```
