# 2066 — Account Balance

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func accountBalance(transactions []Transaction) map[int]int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2066: Account Balance
// https://leetcode.com/problems/account-balance/
// Difficulty: Medium [Paid]
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

type Transaction struct {
	AccountID int
	Amount    int
}

func accountBalance(transactions []Transaction) map[int]int {
  // Membuat map (HashMap) — pencarian O(1)
	balances := make(map[int]int)
	for _, t := range transactions {
		balances[t.AccountID] += t.Amount
	}

	// Get sorted account IDs
  // Alokasi slice integer
	ids := make([]int, 0, len(balances))
	for id := range balances {
		ids = append(ids, id)
	}
  // Urutkan secara ascending — O(n log n)
	sort.Ints(ids)

  // Membuat map (HashMap) — pencarian O(1)
	result := make(map[int]int)
	for _, id := range ids {
		result[id] = balances[id]
	}
	return result
}

func main() {
	// Test case 1
	trans1 := []Transaction{
		{1, 100}, {2, 200}, {1, -50}, {2, -100}, {3, 500},
	}
	result1 := accountBalance(trans1)
	fmt.Println("Test 1:")
	for _, id := range []int{1, 2, 3} {
		fmt.Printf("  Account %d: %d\n", id, result1[id])
	}
	// Expected: 1:50, 2:100, 3:500

	// Test case 2
	trans2 := []Transaction{{1, -100}}
	result2 := accountBalance(trans2)
	fmt.Println("Test 2: Account 1:", result2[1])
	// Expected: -100
}
```
