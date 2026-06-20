# 2525 — Categorize Box According To Criteria

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func CategorizeBoxAccordingToCriteria(length int, width int, height int, mass int) string`

## 🔍 Petunjuk Penyelesaian

**Waktu:** —  |  **Ruang:** —


## 💻 Solusi Go

```go
package main

// LeetCode #2525: Categorize Box According to Criteria
// https://leetcode.com/problems/categorize-box-according-to-criteria/
// Difficulty: Easy
// Time O(1) | Space O(1)

import "fmt"

func main() {
	fmt.Println(CategorizeBoxAccordingToCriteria(1000, 35, 700, 300)) // "Heavy"
	fmt.Println(CategorizeBoxAccordingToCriteria(200, 50, 800, 50))   // "Neither"
}

func CategorizeBoxAccordingToCriteria(length int, width int, height int, mass int) string {
	volume := length * width * height
	isBulky := length >= 10000 || width >= 10000 || height >= 10000 || volume >= 1000000000
	isHeavy := mass >= 100

	if isBulky && isHeavy {
		return "Both"
	}
	if isBulky {
		return "Bulky"
	}
	if isHeavy {
		return "Heavy"
	}
	return "Neither"
}
```
