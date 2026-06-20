# 2287 — Rearrange Characters To Make Target String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func RearrangeCharactersToMakeTargetString(s string, target string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2287: Rearrange Characters to Make Target String
// https://leetcode.com/problems/rearrange-characters-to-make-target-string/
// Difficulty: Easy
// Time O(n + m) | Space O(1)

import "fmt"

func main() {
	fmt.Println(RearrangeCharactersToMakeTargetString("ilovecodingonleetcode", "code")) // 2
	fmt.Println(RearrangeCharactersToMakeTargetString("abcba", "abc"))                  // 1
}

func RearrangeCharactersToMakeTargetString(s string, target string) int {
	sCount := [26]int{}
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(s); i++ {
		sCount[s[i]-'a']++
	}

	tCount := [26]int{}
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(target); i++ {
		tCount[target[i]-'a']++
	}

	maxCopies := len(s)
	for i := 0; i < 26; i++ {
		if tCount[i] > 0 {
			copies := sCount[i] / tCount[i]
			if copies < maxCopies {
				maxCopies = copies
			}
		}
	}
	return maxCopies
}
```
