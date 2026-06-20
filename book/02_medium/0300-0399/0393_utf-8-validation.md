# 0393 — Utf 8 Validation

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string berisi tanda kurung: `()`, `[]`, `{}`. Tugasmu adalah memeriksa apakah string tersebut **valid** — setiap kurung buka harus ditutup oleh kurung yang sesuai dalam urutan benar.

Contoh valid: `()[]{}`, `({[]})`. Tidak valid: `(]`, `([)]`.

**Cara berpikir:** Gunakan Stack. Kurung buka → push. Kurung tutup → pop dan cek kecocokan. Di akhir, stack harus kosong.

**Fungsi Solusi:** `func validUtf8(data []int) bool`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #393: UTF-8 Validation
// https://leetcode.com/problems/utf-8-validation/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func validUtf8(data []int) bool {
	remaining := 0

	for _, b := range data {
		if remaining == 0 {
			if b>>3 == 0b11110 {
				remaining = 3
			} else if b>>4 == 0b1110 {
				remaining = 2
			} else if b>>5 == 0b110 {
				remaining = 1
			} else if b>>7 != 0 {
				return false
			}
		} else {
			if b>>6 != 0b10 {
				return false
			}
			remaining--
		}
	}
	return remaining == 0
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", validUtf8([]int{197, 130, 1}))
	// Expected: true

	// Test case 2
	fmt.Println("Test 2:", validUtf8([]int{235, 140, 4}))
	// Expected: false

	// Test case 3
	fmt.Println("Test 3:", validUtf8([]int{255}))
	// Expected: false
}
```
