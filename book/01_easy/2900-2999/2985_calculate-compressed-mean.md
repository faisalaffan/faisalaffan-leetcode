# 2985 — Calculate Compressed Mean

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan tabel database dan diminta untuk menulis query SQL. Karena repo ini menggunakan Go, query SQL disimulasikan dengan struktur data Go (map untuk grouping, slice untuk sorting, struct untuk representasi row).

Soal tipe ini menguji kemampuanmu menganalisis data relasional — seperti yang kamu lakukan dengan SQL di pekerjaan backend sehari-hari.

**Konsep kunci:** GROUP BY, JOIN, aggregate (SUM, COUNT, AVG), window function (RANK, ROW_NUMBER), HAVING.

**Fungsi yang perlu kamu implementasikan:**
```go
func CalculateCompressedMean(data [][]int) float64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2985: Calculate Compressed Mean
// https://leetcode.com/problems/calculate-compressed-mean/
// Difficulty: Easy [Paid - Database]
//
// Note: This is a SQL problem on LeetCode. In Go, we implement the
// equivalent weighted average calculation.

import "fmt"

func main() {
	// LeetCode name: calculateCompressedMean
	// Input: item_count, order_occurrences pairs
	fmt.Println(CalculateCompressedMean([][]int{{1, 500}, {2, 1000}, {3, 800}, {4, 200}}))
	fmt.Println(CalculateCompressedMean([][]int{{1, 5}, {2, 10}, {3, 5}})) // 2
}

// Time: O(n) | Space: O(1)
// LeetCode submission name: calculateCompressedMean
// data[i] = [item_count, order_occurrences]
func CalculateCompressedMean(data [][]int) float64 {
	var totalItems, totalOrders float64
	for _, row := range data {
		totalItems += float64(row[0] * row[1])
		totalOrders += float64(row[1])
	}
	result := totalItems / totalOrders
	// Round to 2 decimal places
	return float64(int(result*100+0.5)) / 100
}
```
