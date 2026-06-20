# 2084 — Drop Type 1 Orders For Customers With Type 0 Orders

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func dropOrders(orders []Order) []Order
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2084: Drop Type 1 Orders for Customers With Type 0 Orders
// https://leetcode.com/problems/drop-type-1-orders-for-customers-with-type-0-orders/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(n)

import "fmt"

type Order struct {
	OrderID  int
	Customer int
	OrderType int // 0 or 1
}

func dropOrders(orders []Order) []Order {
	// Find customers with type 0 orders
  // Membuat map (HashMap) — pencarian O(1)
	hasTypeZero := make(map[int]bool)
	for _, o := range orders {
		if o.OrderType == 0 {
			hasTypeZero[o.Customer] = true
		}
	}

	// Keep orders that don't need to be dropped
	result := []Order{}
	for _, o := range orders {
		if !(o.OrderType == 1 && hasTypeZero[o.Customer]) {
			result = append(result, o)
		}
	}
	return result
}

func main() {
	// Test case 1
	orders1 := []Order{
		{1, 1, 0},
		{2, 1, 1},
		{3, 2, 1},
		{4, 2, 0},
		{5, 3, 1},
	}
	result1 := dropOrders(orders1)
	fmt.Println("Test 1:")
	for _, o := range result1 {
		fmt.Printf("  Order %d (Customer %d, Type %d)\n", o.OrderID, o.Customer, o.OrderType)
	}
	// Expected: orders 1, 4, 5 (2 dropped because customer 1 has type 0)

	// Test case 2
	orders2 := []Order{
		{1, 1, 0},
		{2, 2, 0},
	}
	result2 := dropOrders(orders2)
	fmt.Println("Test 2:", len(result2))
	// Expected: 2 (no type 1 orders to drop)

	// Test case 3
	fmt.Println("Test 3:", len(dropOrders(nil)))
	// Expected: 0
}
```
