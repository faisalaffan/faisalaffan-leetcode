# 1647 — Minimum Deletions To Make Character Frequencies Unique

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MinDeletions(s string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(N), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1647: Minimum Deletions to Make Character Frequencies Unique
// https://leetcode.com/problems/minimum-deletions-to-make-character-frequencies-unique/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(MinDeletions("aab"))
	fmt.Println(MinDeletions("aaabbbcc"))
	fmt.Println(MinDeletions("ceabaacb"))
}

func MinDeletions(s string) int {
	// Time: O(N), Space: O(1)
  // Alokasi slice integer
	freq := make([]int, 26)
	for _, ch := range s {
		freq[ch-'a']++
	}

  // Membuat map (HashMap) — pencarian O(1)
	used := make(map[int]bool)
	deletions := 0

	for _, f := range freq {
		for f > 0 && used[f] {
			f--
			deletions++
		}
		if f > 0 {
			used[f] = true
		}
	}

	return deletions
}
```
