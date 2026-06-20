# 2890 — Reshape Data Melt

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func ReshapeDataMelt(data [][]string) [][]string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n * m) where m is number of store columns  
**Kompleksitas Ruang:** O(n * m)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2890: Reshape Data: Melt
// https://leetcode.com/problems/reshape-data-melt/
// Difficulty: Easy
//
// Note: This is a Pandas-only problem on LeetCode (Python).
// In Go, we melt from wide format [product, store1, store2] to
// long format [product, store, price].

import "fmt"

func main() {
	// LeetCode name: meltTable
	// Input: [product_id, store1_price, store2_price]
	fmt.Println(ReshapeDataMelt([][]string{{"P1", "100", "200"}, {"P2", "150", "250"}}))
	// [[P1 store1 100] [P1 store2 200] [P2 store1 150] [P2 store2 250]]
}

// Time: O(n * m) where m is number of store columns | Space: O(n * m)
// LeetCode submission name: meltTable
func ReshapeDataMelt(data [][]string) [][]string {
	result := [][]string{}
	for _, row := range data {
		productID := row[0]
		for j := 1; j < len(row); j++ {
			store := fmt.Sprintf("store%d", j)
			result = append(result, []string{productID, store, row[j]})
		}
	}
	return result
}
```
