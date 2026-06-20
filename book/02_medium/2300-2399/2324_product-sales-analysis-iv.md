# 2324 — Product Sales Analysis Iv

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func productSalesAnalysis(sales [][]int) [][]int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2324: Product Sales Analysis IV
// https://leetcode.com/problems/product-sales-analysis-iv/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(n)

import "fmt"

type Sale struct {
	UserID    int
	ProductID int
	Quantity  int
}

func productSalesAnalysis(sales [][]int) [][]int {
	// sales[i] = [sale_id, product_id, user_id, quantity]
  // Membuat map (HashMap) — pencarian O(1)
	userProductQty := make(map[[2]int]int) // [user_id, product_id] -> total qty

	for _, s := range sales {
		productID, userID, qty := s[1], s[2], s[3]
		key := [2]int{userID, productID}
		userProductQty[key] += qty
	}

	// For each user, find product with max qty
  // Membuat map (HashMap) — pencarian O(1)
	userMax := make(map[int]struct{ qty int; product int })
	for key, qty := range userProductQty {
		userID, productID := key[0], key[1]
		if prev, ok := userMax[userID]; !ok || qty > prev.qty || (qty == prev.qty && productID > prev.product) {
			userMax[userID] = struct{ qty int; product int }{qty, productID}
		}
	}

	result := [][]int{}
	// Sort by userID
	for u := 1; u <= len(userMax); u++ {
		if v, ok := userMax[u]; ok {
			result = append(result, []int{u, v.product, v.qty})
		}
	}
	return result
}

func main() {
	// Test case 1
	fmt.Println(productSalesAnalysis([][]int{{1, 1, 1, 5}, {2, 1, 1, 3}, {3, 2, 1, 4}, {4, 2, 2, 2}}))
	// Expected: [[1, 1, 8], [2, 2, 2]]

	// Test case 2
	fmt.Println(productSalesAnalysis([][]int{{1, 1, 1, 1}, {2, 2, 1, 1}}))
	// Expected: [[1, 2, 1]]
}
```
