# 1393 — Capital Gainloss

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func capitalGainLoss(stocks []struct {
	stockName    string
	operation    string
	operationDay int
	price        int
}) []gainLoss
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n) where n = number of transactions  
**Kompleksitas Ruang:** O(k) where k = number of unique stock names

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

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
  // Membuat map (HashMap) — pencarian O(1)
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
