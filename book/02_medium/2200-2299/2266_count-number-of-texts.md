# 2266 — Count Number Of Texts

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func countTexts(pressedKeys string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2266: Count Number of Texts
// https://leetcode.com/problems/count-number-of-texts/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func countTexts(pressedKeys string) int {
	const mod = 1_000_000_007
	n := len(pressedKeys)
  // Alokasi slice integer
	dp := make([]int, n+1)
	dp[0] = 1

	for i := 1; i <= n; i++ {
		dp[i] = dp[i-1] // press once
		// Check for multiple presses of same digit
		maxPress := 3
		if pressedKeys[i-1] == '7' || pressedKeys[i-1] == '9' {
			maxPress = 4
		}
		for j := 2; j <= maxPress && j <= i; j++ {
			if pressedKeys[i-j] == pressedKeys[i-1] {
				dp[i] = (dp[i] + dp[i-j]) % mod
			} else {
				break
			}
		}
	}
	return dp[n]
}

func main() {
	// Test case 1
	fmt.Println(countTexts("22233"))
	// Expected: 8

	// Test case 2
	fmt.Println(countTexts("222222222222222222222222222222222222"))
	// Expected: 82876089

	// Test case 3
	fmt.Println(countTexts("33"))
	// Expected: 2
}
```
