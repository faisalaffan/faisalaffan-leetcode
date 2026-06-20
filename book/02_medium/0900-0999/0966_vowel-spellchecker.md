# 0966 — Vowel Spellchecker

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func spellchecker(wordlist []string, queries []string) []string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n * L)  
**Kompleksitas Ruang:** O(n * L)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #966: Vowel Spellchecker
// https://leetcode.com/problems/vowel-spellchecker/
// Difficulty: Medium

import (
	"fmt"
	"strings"
)

// Time: O(n * L) | Space: O(n * L)
func spellchecker(wordlist []string, queries []string) []string {
  // Membuat map (HashMap) — pencarian O(1)
	exact := make(map[string]bool)
  // Membuat map (HashMap) — pencarian O(1)
	lower := make(map[string]string)
  // Membuat map (HashMap) — pencarian O(1)
	vowel := make(map[string]string)

	for _, w := range wordlist {
		exact[w] = true
		lo := strings.ToLower(w)
		if _, ok := lower[lo]; !ok {
			lower[lo] = w
		}
		vw := devowel(lo)
		if _, ok := vowel[vw]; !ok {
			vowel[vw] = w
		}
	}

	ans := make([]string, len(queries))
	for i, q := range queries {
		if exact[q] {
			ans[i] = q
		} else if w, ok := lower[strings.ToLower(q)]; ok {
			ans[i] = w
		} else if w, ok := vowel[devowel(strings.ToLower(q))]; ok {
			ans[i] = w
		} else {
			ans[i] = ""
		}
	}
	return ans
}

func devowel(s string) string {
	var sb strings.Builder
	for _, ch := range s {
		if ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u' {
			sb.WriteByte('*')
		} else {
			sb.WriteRune(ch)
		}
	}
	return sb.String()
}

func main() {
	fmt.Println(spellchecker([]string{"KiTe", "kite", "hare", "Hare"}, []string{"kite", "Kite", "KiTe", "Hare", "HARE", "Hear", "hear", "keti", "keet", "keto"}))
	fmt.Println(spellchecker([]string{"yellow", "wood"}, []string{"Yello", "wood", "yellow"}))
}
```
