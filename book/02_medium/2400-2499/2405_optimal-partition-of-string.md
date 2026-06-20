# 2405 — Optimal Partition Of String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func partitionString(s string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2405: Optimal Partition of String
// https://leetcode.com/problems/optimal-partition-of-string/
// Difficulty: Medium
// Time: O(n) | Space: O(1)
// Greedy: start new partition when duplicate char found.

import "fmt"

func main() {
	fmt.Println(partitionString("abacaba")) // 4
	fmt.Println(partitionString("ssssss"))  // 6
}

func partitionString(s string) int {
	ans := 1
  // Membuat map (HashMap) — pencarian O(1)
	seen := make(map[byte]bool)
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(s); i++ {
		if seen[s[i]] {
			ans++
			seen = make(map[byte]bool)
		}
		seen[s[i]] = true
	}
	return ans
}
```
