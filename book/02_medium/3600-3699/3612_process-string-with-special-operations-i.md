# 3612 — Process String With Special Operations I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func ProcessStringWithSpecialOperationsI(s string, ops []int) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3612: Process String with Special Operations I
// https://leetcode.com/problems/process-string-with-special-operations-i/
// Difficulty: Medium
// Complexity: O(n) time, O(n) space

import "fmt"

func main() {
	// Test case 1
	fmt.Println("Test 1:", ProcessStringWithSpecialOperationsI("abc", []int{1, 0, 1}))
	// Test case 2
	fmt.Println("Test 2:", ProcessStringWithSpecialOperationsI("ab", []int{1, 1}))
	// Test case 3
	fmt.Println("Test 3:", ProcessStringWithSpecialOperationsI("x", []int{0}))
}

func ProcessStringWithSpecialOperationsI(s string, ops []int) string {
	b := []byte(s)
	for i, op := range ops {
		if i >= len(b) {
			break
		}
		if op == 1 {
			// toggle case
			if b[i] >= 'a' && b[i] <= 'z' {
				b[i] = b[i] - 'a' + 'A'
			} else if b[i] >= 'A' && b[i] <= 'Z' {
				b[i] = b[i] - 'A' + 'a'
			}
		}
		// 0 = no operation
	}
	return string(b)
}
```
