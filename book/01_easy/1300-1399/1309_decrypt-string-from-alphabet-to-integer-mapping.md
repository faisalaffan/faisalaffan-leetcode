# 1309 — Decrypt String From Alphabet To Integer Mapping

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func freqAlphabets(s string) string

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

// LeetCode #1309: Decrypt String from Alphabet to Integer Mapping
// https://leetcode.com/problems/decrypt-string-from-alphabet-to-integer-mapping/
// Difficulty: Easy
//
// LeetCode submission: func freqAlphabets(s string) string

import "fmt"

func main() {
	fmt.Println(DecryptStringFromAlphabetToIntegerMapping("10#11#12"))       // "jkab"
	fmt.Println(DecryptStringFromAlphabetToIntegerMapping("1326#"))           // "acz"
	fmt.Println(DecryptStringFromAlphabetToIntegerMapping("12345678910#11#12#13#14#15#16#17#18#19#20#21#22#23#24#25#26#")) // "abcdefghijklmnopqrstuvwxyz"
}

// Time: O(n), Space: O(n)
func DecryptStringFromAlphabetToIntegerMapping(s string) string {
	res := make([]byte, 0, len(s))
	i := 0
	for i < len(s) {
		if i+2 < len(s) && s[i+2] == '#' {
			num := (s[i]-'0')*10 + (s[i+1] - '0')
			res = append(res, byte('a'+num-1))
			i += 3
		} else {
			num := s[i] - '0'
			res = append(res, byte('a'+num-1))
			i++
		}
	}
	return string(res)
}
```
