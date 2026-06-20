# 2582 — Pass The Pillow

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func PassThePillow(n int, time int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** —  |  **Ruang:** —


## 💻 Solusi Go

```go
package main

// LeetCode #2582: Pass the Pillow
// https://leetcode.com/problems/pass-the-pillow/
// Difficulty: Easy
// Time O(1) | Space O(1)

import "fmt"

func main() {
	fmt.Println(PassThePillow(4, 5)) // 2
	fmt.Println(PassThePillow(3, 2)) // 3
}

func PassThePillow(n int, time int) int {
	cycle := 2 * (n - 1)
	t := time % cycle
	if t < n {
		return t + 1
	}
	return n - (t - n + 1)
}
```
