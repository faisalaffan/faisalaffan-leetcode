# 0014 — Longest Common Prefix

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func LongestCommonPrefix(strs []string) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Prefix Sum

**Kompleksitas Waktu:** O(n*m)  
**Kompleksitas Ruang:** O(1) where n=len(strs), m=len(shortest string)

> **Untuk fresh graduate:** Kuasai dulu teknik **Prefix Sum** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #14: Longest Common Prefix
// https://leetcode.com/problems/longest-common-prefix/
// Difficulty: Easy

import "fmt"

// Time: O(n*m) | Space: O(1) where n=len(strs), m=len(shortest string)
func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(strs[0]); i++ {
		c := strs[0][i]
		for j := 1; j < len(strs); j++ {
			if i == len(strs[j]) || strs[j][i] != c {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}

func main() {
	fmt.Println(LongestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Println(LongestCommonPrefix([]string{"dog", "racecar", "car"}))
	fmt.Println(LongestCommonPrefix([]string{"a"}))
}
```
