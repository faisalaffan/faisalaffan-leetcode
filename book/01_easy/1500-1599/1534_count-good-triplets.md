# 1534 — Count Good Triplets

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func countGoodTriplets(arr []int, a int, b int, c int) int

import "fmt"

func main()
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n^3), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1534: Count Good Triplets
// https://leetcode.com/problems/count-good-triplets/
// Difficulty: Easy
//
// LeetCode submission: func countGoodTriplets(arr []int, a int, b int, c int) int

import "fmt"

func main() {
	fmt.Println(CountGoodTriplets([]int{3, 0, 1, 1, 9, 7}, 7, 2, 3)) // 4
	fmt.Println(CountGoodTriplets([]int{1, 1, 2, 2, 3}, 0, 0, 1))   // 0
}

// Time: O(n^3), Space: O(1)
func CountGoodTriplets(arr []int, a int, b int, c int) int {
	n, count := len(arr), 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if abs(arr[i]-arr[j]) > a {
				continue
			}
			for k := j + 1; k < n; k++ {
				if abs(arr[j]-arr[k]) <= b && abs(arr[i]-arr[k]) <= c {
					count++
				}
			}
		}
	}
	return count
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```
