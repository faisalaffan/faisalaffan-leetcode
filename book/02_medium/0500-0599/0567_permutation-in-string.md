# 0567 — Permutation In String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func CheckInclusion(s1 string, s2 string) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n + m) where n = len(s1), m = len(s2)  
**Kompleksitas Ruang:** O(1) (fixed 26 chars)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #567: Permutation in String
// https://leetcode.com/problems/permutation-in-string/
// Difficulty: Medium
// Time: O(n + m) where n = len(s1), m = len(s2)
// Space: O(1) (fixed 26 chars)

import "fmt"

func main() {
	fmt.Println(CheckInclusion("ab", "eidbaooo"))
	fmt.Println(CheckInclusion("ab", "eidboaoo"))
}

func CheckInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	count1 := [26]int{}
	count2 := [26]int{}

  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(s1); i++ {
		count1[s1[i]-'a']++
		count2[s2[i]-'a']++
	}

	if count1 == count2 {
		return true
	}

	for i := len(s1); i < len(s2); i++ {
		count2[s2[i]-'a']++
		count2[s2[i-len(s1)]-'a']--
		if count1 == count2 {
			return true
		}
	}

	return false
}
```
