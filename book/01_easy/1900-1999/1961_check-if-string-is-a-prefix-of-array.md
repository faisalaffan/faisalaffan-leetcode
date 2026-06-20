# 1961 — Check If String Is A Prefix Of Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func CheckIfStringIsAPrefixOfArray(s string, words []string) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Prefix Sum

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Prefix Sum** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1961: Check If String Is a Prefix of Array
// https://leetcode.com/problems/check-if-string-is-a-prefix-of-array/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(CheckIfStringIsAPrefixOfArray("iloveleetcode", []string{"i", "love", "leetcode", "apples"})) // true
	fmt.Println(CheckIfStringIsAPrefixOfArray("iloveleetcode", []string{"apples", "i", "love", "leetcode"})) // false
}

// Time: O(n), Space: O(1)
func CheckIfStringIsAPrefixOfArray(s string, words []string) bool {
	i := 0
	for _, w := range words {
		if i >= len(s) {
			break
		}
		for j := 0; j < len(w); j++ {
			if i >= len(s) || s[i] != w[j] {
				return false
			}
			i++
		}
	}
	return i == len(s)
}
```
