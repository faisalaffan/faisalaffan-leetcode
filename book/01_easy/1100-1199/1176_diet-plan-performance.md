# 1176 — Diet Plan Performance

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func dietPlanPerformance(calories []int, k int, lower int, upper int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1176: Diet Plan Performance
// https://leetcode.com/problems/diet-plan-performance/
// Difficulty: Easy [Paid]
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(dietPlanPerformance([]int{1, 2, 3, 4, 5}, 1, 3, 3)) // 0
	fmt.Println(dietPlanPerformance([]int{3, 2}, 2, 0, 1))          // 1
}

// LeetCode submission: dietPlanPerformance
func dietPlanPerformance(calories []int, k int, lower int, upper int) int {
	n := len(calories)
	if n < k {
		return 0
	}
	sum := 0
	for i := 0; i < k; i++ {
		sum += calories[i]
	}
	ans := score(sum, lower, upper)
	for i := k; i < n; i++ {
		sum += calories[i] - calories[i-k]
		ans += score(sum, lower, upper)
	}
	return ans
}

func score(sum, lower, upper int) int {
	if sum < lower {
		return -1
	}
	if sum > upper {
		return 1
	}
	return 0
}
```
