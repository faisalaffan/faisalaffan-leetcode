# 0367 — Valid Perfect Square

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string berisi tanda kurung: `()`, `[]`, `{}`. Tugasmu adalah memeriksa apakah string tersebut **valid** — setiap kurung buka harus ditutup oleh kurung yang sesuai dalam urutan benar.

Contoh valid: `()[]{}`, `({[]})`. Tidak valid: `(]`, `([)]`.

**Cara berpikir:** Gunakan Stack. Kurung buka → push. Kurung tutup → pop dan cek kecocokan. Di akhir, stack harus kosong.

**Fungsi Solusi:** `func ValidPerfectSquare(num int) bool`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(log n), Space: O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #367: Valid Perfect Square
// https://leetcode.com/problems/valid-perfect-square/
// Difficulty: Easy

import "fmt"

// Time: O(log n), Space: O(1)
func ValidPerfectSquare(num int) bool {
	lo, hi := 1, num
	for lo <= hi {
		mid := lo + (hi-lo)/2
		sq := mid * mid
		if sq == num {
			return true
		} else if sq < num {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return false
}

func main() {
	fmt.Println(ValidPerfectSquare(16))
	fmt.Println(ValidPerfectSquare(14))
	fmt.Println(ValidPerfectSquare(1))
}
```
