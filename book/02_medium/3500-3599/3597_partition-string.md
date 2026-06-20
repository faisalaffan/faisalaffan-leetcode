# 3597 — Partition String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func PartitionString(s string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3597: Partition String
// https://leetcode.com/problems/partition-string/
// Difficulty: Medium
// Complexity: O(n) time, O(1) space

import "fmt"

func main() {
	// Test case 1
	fmt.Println("Test 1:", PartitionString("abac"))
	// Test case 2
	fmt.Println("Test 2:", PartitionString("aaaa"))
	// Test case 3
	fmt.Println("Test 3:", PartitionString("abc"))
}

func PartitionString(s string) int {
	// Partition into substrings with unique characters
	// Use greedy: start new partition when duplicate found
	count := 1
  // Membuat map (HashMap) — pencarian O(1)
	seen := make(map[byte]bool)
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(s); i++ {
		if seen[s[i]] {
			count++
			seen = make(map[byte]bool)
		}
		seen[s[i]] = true
	}
	return count
}
```
