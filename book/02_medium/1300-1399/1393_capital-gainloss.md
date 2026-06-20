# 1393 — Capital Gainloss

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func capitalGainLoss(stocks []struct { stockName string operation string operationDay int price int }) []gainLoss`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n) where n = number of transactions  |  **Ruang:** O(k) where k = number of unique stock names

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1393: Capital Gain/Loss
// https://leetcode.com/problems/capital-gainloss/
// Difficulty: Medium

import "fmt"

func main() {
	result := capitalGainLoss(
		[]struct {
			stockName    string
			operation    string
			operationDay int
			price        int
		}{
			{"Leetcode", "Buy", 1, 1000},
			{"Corona Masks", "Buy", 2, 10},
			{"Leetcode", "Sell", 5, 9000},
			{"Handbags", "Buy", 17, 30000},
			{"Corona Masks", "Sell", 3, 1010},
			{"Corona Masks", "Buy", 4, 1000},
			{"Corona Masks", "Sell", 5, 500},
			{"Corona Masks", "Buy", 6, 1000},
			{"Corona Masks", "Sell", 7, 500},
			{"Handbags", "Sell", 29, 7000},
		},
	)
	for _, r := range result {
		fmt.Printf("%s %d\n", r.name, r.gain)
	}
}

type gainLoss struct {
	name string
	gain int
}

// Time: O(n) where n = number of transactions
// Space: O(k) where k = number of unique stock names
func capitalGainLoss(stocks []struct {
	stockName    string
	operation    string
	operationDay int
	price        int
}) []gainLoss {
  // HashMap: O(1) lookup
	holdings := make(map[string]int) // net cost of buys
	for _, s := range stocks {
		if s.operation == "Buy" {
			holdings[s.stockName] -= s.price
		} else {
			holdings[s.stockName] += s.price
		}
	}

	var result []gainLoss
	for name, gain := range holdings {
		result = append(result, gainLoss{name, gain})
	}
	return result
}
```
