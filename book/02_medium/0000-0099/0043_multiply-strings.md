# 0043 — Multiply Strings

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func multiply(num1 string, num2 string) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(m*n)  
**Kompleksitas Ruang:** O(m+n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #43: Multiply Strings
// https://leetcode.com/problems/multiply-strings/
// Difficulty: Medium

import "fmt"

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	m, n := len(num1), len(num2)
	result := make([]byte, m+n)

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			prod := (num1[i]-'0')*(num2[j]-'0') + result[i+j+1]
			result[i+j+1] = prod % 10
			result[i+j] += prod / 10
		}
	}

	if result[0] == 0 {
		result = result[1:]
	}

  // Range loop: iterasi dengan indeks + nilai
	for i := range result {
		result[i] += '0'
	}

	return string(result)
}

func main() {
	// Test case 1
	fmt.Println(multiply("2", "3")) // "6"

	// Test case 2
	fmt.Println(multiply("123", "456")) // "56088"

	// Test case 3
	fmt.Println(multiply("999", "999")) // "998001"
}

// Time: O(m*n) | Space: O(m+n)
```
