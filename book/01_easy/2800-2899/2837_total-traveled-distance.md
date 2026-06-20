# 2837 — Total Traveled Distance

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func TotalTraveledDistance(rides []struct {
	UserID   int
	Distance int
}) map[int]int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2837: Total Traveled Distance
// https://leetcode.com/problems/total-traveled-distance/
// Difficulty: Easy [Paid]
// Time: O(n) | Space: O(n)
// Note: SQL problem, adapted to Go. Computes total travel distance per user.

import "fmt"

func main() {
	rides := []struct {
		UserID   int
		Distance int
	}{
		{1, 10},
		{1, 15},
		{2, 20},
		{3, 0},
	}
	fmt.Println(TotalTraveledDistance(rides))
}

func TotalTraveledDistance(rides []struct {
	UserID   int
	Distance int
}) map[int]int {
	total := map[int]int{}
	for _, r := range rides {
		total[r.UserID] += r.Distance
	}
	return total
}
```
