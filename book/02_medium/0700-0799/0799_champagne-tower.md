# 0799 — Champagne Tower

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan tabel database dan diminta untuk menulis query SQL. Karena repo ini menggunakan Go, query SQL disimulasikan dengan struktur data Go (map untuk grouping, slice untuk sorting, struct untuk representasi row).

Soal tipe ini menguji kemampuanmu menganalisis data relasional — seperti yang kamu lakukan dengan SQL di pekerjaan backend sehari-hari.

**Konsep kunci:** GROUP BY, JOIN, aggregate (SUM, COUNT, AVG), window function (RANK, ROW_NUMBER), HAVING.

**Fungsi yang perlu kamu implementasikan:**
```go
func champagneTower(poured int, queryRow int, queryGlass int) float64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming

**Kompleksitas Waktu:** O(query_row^2)  
**Kompleksitas Ruang:** O(query_row)

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #799: Champagne Tower
// https://leetcode.com/problems/champagne-tower/
// Difficulty: Medium
// Time: O(query_row^2)
// Space: O(query_row)

import "fmt"

func main() {
	fmt.Println(champagneTower(1, 1, 1))
	fmt.Println(champagneTower(2, 1, 1))
	fmt.Println(champagneTower(100000009, 33, 17))
}

func champagneTower(poured int, queryRow int, queryGlass int) float64 {
	dp := make([]float64, queryRow+1)
	dp[0] = float64(poured)

	for row := 0; row < queryRow; row++ {
		next := make([]float64, queryRow+2)
		for col := 0; col <= row; col++ {
			if dp[col] > 1.0 {
				excess := (dp[col] - 1.0) / 2.0
				next[col] += excess
				next[col+1] += excess
			}
		}
		dp = next
	}

	if dp[queryGlass] > 1.0 {
		return 1.0
	}
	return dp[queryGlass]
}
```
