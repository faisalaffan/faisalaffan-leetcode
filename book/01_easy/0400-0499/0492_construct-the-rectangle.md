# 0492 — Construct The Rectangle

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func ConstructTheRectangle(area int) []int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(sqrt(n)), Space: O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #492: Construct the Rectangle
// https://leetcode.com/problems/construct-the-rectangle/
// Difficulty: Easy

import "fmt"

// Time: O(sqrt(n)), Space: O(1)
func ConstructTheRectangle(area int) []int {
	w := 1
	for i := 1; i*i <= area; i++ {
		if area%i == 0 {
			w = i
		}
	}
	return []int{area / w, w}
}

func main() {
	fmt.Println(ConstructTheRectangle(4))
	fmt.Println(ConstructTheRectangle(37))
	fmt.Println(ConstructTheRectangle(122122))
}
```
