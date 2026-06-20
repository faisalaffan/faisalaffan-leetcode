# 1370 — Increasing Decreasing String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func sortString(s string) string

import "fmt"

func main()
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1370: Increasing Decreasing String
// https://leetcode.com/problems/increasing-decreasing-string/
// Difficulty: Easy
//
// LeetCode submission: func sortString(s string) string

import "fmt"

func main() {
	fmt.Println(IncreasingDecreasingString("aaaabbbbcccc")) // "abccbaabccba"
	fmt.Println(IncreasingDecreasingString("rat"))          // "art"
	fmt.Println(IncreasingDecreasingString("leetcode"))     // "cdelotee"
}

// Time: O(n), Space: O(n)
func IncreasingDecreasingString(s string) string {
  // Alokasi slice integer
	freq := make([]int, 26)
	for _, ch := range s {
		freq[ch-'a']++
	}

	res := make([]byte, 0, len(s))
	for len(res) < len(s) {
		for i := 0; i < 26; i++ {
			if freq[i] > 0 {
				res = append(res, byte('a'+i))
				freq[i]--
			}
		}
		for i := 25; i >= 0; i-- {
			if freq[i] > 0 {
				res = append(res, byte('a'+i))
				freq[i]--
			}
		}
	}
	return string(res)
}
```
