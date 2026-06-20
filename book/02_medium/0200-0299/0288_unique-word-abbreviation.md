# 0288 — Unique Word Abbreviation

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func Constructor(dictionary []string) ValidWordAbbr
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n) for init, O(1) for isUnique, Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #288: Unique Word Abbreviation
// https://leetcode.com/problems/unique-word-abbreviation/
// Difficulty: Medium [Paid]
// Time: O(n) for init, O(1) for isUnique, Space: O(n)

import (
	"fmt"
	"strconv"
)

type ValidWordAbbr struct {
	abbrMap map[string]string
}

func Constructor(dictionary []string) ValidWordAbbr {
  // Membuat map (HashMap) — pencarian O(1)
	abbrMap := make(map[string]string)
	for _, word := range dictionary {
		abbr := getAbbr(word)
		if existing, ok := abbrMap[abbr]; ok {
			if existing != word {
				abbrMap[abbr] = ""
			}
		} else {
			abbrMap[abbr] = word
		}
	}
	return ValidWordAbbr{abbrMap}
}

func (this *ValidWordAbbr) IsUnique(word string) bool {
	abbr := getAbbr(word)
	val, ok := this.abbrMap[abbr]
	return !ok || val == word
}

func getAbbr(s string) string {
	if len(s) <= 2 {
		return s
	}
	return string(s[0]) + strconv.Itoa(len(s)-2) + string(s[len(s)-1])
}

func main() {
	vwa := Constructor([]string{"deer", "door", "cake", "card"})
	fmt.Println(vwa.IsUnique("dear"))
	fmt.Println(vwa.IsUnique("cart"))
	fmt.Println(vwa.IsUnique("cane"))
	fmt.Println(vwa.IsUnique("make"))

	vwa2 := Constructor([]string{"a", "a"})
	fmt.Println(vwa2.IsUnique("a"))
}
```
