# 1445 — Apples Oranges

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan tabel database dan diminta untuk menulis query SQL. Karena repo ini menggunakan Go, query SQL disimulasikan dengan struktur data Go (map untuk grouping, slice untuk sorting, struct untuk representasi row).

Soal tipe ini menguji kemampuanmu menganalisis data relasional — seperti yang kamu lakukan dengan SQL di pekerjaan backend sehari-hari.

**Konsep kunci:** GROUP BY, JOIN, aggregate (SUM, COUNT, AVG), window function (RANK, ROW_NUMBER), HAVING.

**Fungsi yang perlu kamu implementasikan:**
```go
func applesOranges(sales []struct {
	saleDate string
	fruit    string
	soldNum  int
}) []diffResult
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n) for sorting  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1445: Apples & Oranges
// https://leetcode.com/problems/apples-oranges/
// Difficulty: Medium

import "fmt"

func main() {
	// SQL problem - simulating in Go
	result := applesOranges([]struct {
		saleDate string
		fruit    string
		soldNum  int
	}{
		{"2020-05-01", "apples", 10},
		{"2020-05-01", "oranges", 8},
		{"2020-05-02", "apples", 15},
		{"2020-05-02", "oranges", 15},
		{"2020-05-03", "apples", 20},
		{"2020-05-03", "oranges", 0},
		{"2020-05-04", "apples", 15},
		{"2020-05-04", "oranges", 16},
	})
	for _, r := range result {
		fmt.Printf("%s %d\n", r.date, r.diff)
	}
}

type diffResult struct {
	date string
	diff int
}

// Time: O(n log n) for sorting
// Space: O(n)
func applesOranges(sales []struct {
	saleDate string
	fruit    string
	soldNum  int
}) []diffResult {
	// Group by date
  // Membuat map (HashMap) — pencarian O(1)
	apples := make(map[string]int)
  // Membuat map (HashMap) — pencarian O(1)
	oranges := make(map[string]int)

  // Membuat map (HashMap) — pencarian O(1)
	dateSet := make(map[string]bool)
	for _, s := range sales {
		dateSet[s.saleDate] = true
		if s.fruit == "apples" {
			apples[s.saleDate] += s.soldNum
		} else {
			oranges[s.saleDate] += s.soldNum
		}
	}

	// Sort dates
	dates := make([]string, 0, len(dateSet))
	for d := range dateSet {
		dates = append(dates, d)
	}

	var result []diffResult
	for _, d := range dates {
		diff := apples[d] - oranges[d]
		result = append(result, diffResult{d, diff})
	}

	return result
}
```
