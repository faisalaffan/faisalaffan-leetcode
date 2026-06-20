# 3015 — Count The Number Of Houses At A Certain Distance I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func countOfPairs(n int, x int, y int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n^2)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3015: Count the Number of Houses at a Certain Distance I
// https://leetcode.com/problems/count-the-number-of-houses-at-a-certain-distance-i/
// Difficulty: Medium
// Time: O(n^2) | Space: O(n)

import "fmt"

func main() {
	fmt.Println(countOfPairs(3, 1, 3))
	fmt.Println(countOfPairs(5, 2, 4))
	fmt.Println(countOfPairs(4, 1, 1))
}

func countOfPairs(n int, x int, y int) []int {
  // Alokasi slice integer
	ans := make([]int, n)
	if x > y {
		x, y = y, x
	}
	for i := 1; i <= n; i++ {
		for j := i + 1; j <= n; j++ {
			dist := j - i
			via := abs(i-x) + 1 + abs(y-j)
			if via < dist {
				dist = via
			}
			ans[dist-1] += 2
		}
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```
