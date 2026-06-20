# 3052 — Maximize Items

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan data tabel database. Tugasmu adalah menganalisis data tersebut. Karena repo ini Go, query SQL disimulasikan dengan map, slice, dan struct.

**Cara berpikir:** Tentukan SELECT, FROM, JOIN, GROUP BY, ORDER BY. Lalu terjemahkan ke Go.

**Fungsi Solusi:** `func maximizeItems(inventory []InventoryItem) []itemCountResult`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3052: Maximize Items (SQL simulation)
// https://leetcode.com/problems/maximize-items/
// Difficulty: Hard [Paid]
//
// Approach: Given warehouse space (500,000 sq ft) and inventory items grouped
// by type (prime_eligible, not_prime), maximize total items by filling the
// warehouse with whole batches of each type. Prime eligible items have priority.

import "fmt"

type InventoryItem struct {
	ItemID        int
	ItemType      string
	ItemCategory  string
	SquareFootage float64
}

type itemCountResult struct {
	ItemType string
	Count    int64
}

func maximizeItems(inventory []InventoryItem) []itemCountResult {
	const warehouseSpace = 500000.0
	type typeInfo struct {
		totalSqFt float64
		count     int64
	}
  // HashMap: O(1) lookup
	groups := make(map[string]*typeInfo)
	for _, item := range inventory {
		if groups[item.ItemType] == nil {
			groups[item.ItemType] = &typeInfo{}
		}
		groups[item.ItemType].totalSqFt += item.SquareFootage
		groups[item.ItemType].count++
	}
	var result []itemCountResult
	prime, primeExists := groups["prime_eligible"]
	nonPrime, nonPrimeExists := groups["not_prime"]
	if primeExists && prime.totalSqFt > 0 {
		primeBatches := int64(warehouseSpace / prime.totalSqFt)
		primeCount := primeBatches * prime.count
		result = append(result, itemCountResult{"prime_eligible", primeCount})
		remaining := warehouseSpace - float64(primeBatches)*prime.totalSqFt
		if nonPrimeExists && nonPrime.totalSqFt > 0 && remaining > 0 {
			nonPrimeBatches := int64(remaining / nonPrime.totalSqFt)
			nonPrimeCount := nonPrimeBatches * nonPrime.count
			result = append(result, itemCountResult{"not_prime", nonPrimeCount})
		} else if nonPrimeExists {
			result = append(result, itemCountResult{"not_prime", 0})
		}
	} else if nonPrimeExists {
		nonPrimeBatches := int64(warehouseSpace / nonPrime.totalSqFt)
		nonPrimeCount := nonPrimeBatches * nonPrime.count
		result = append(result, itemCountResult{"not_prime", nonPrimeCount})
	}
	return result
}

func main() {
	// Example 1: mixed types
	inventory1 := []InventoryItem{
		{1, "prime_eligible", "Watches", 100.0},
		{2, "prime_eligible", "Art", 200.0},
		{3, "not_prime", "Books", 50.0},
		{4, "not_prime", "Toys", 30.0},
	}
	fmt.Println("Example 1:")
	for _, r := range maximizeItems(inventory1) {
		fmt.Printf("  %s %d\n", r.ItemType, r.Count)
	}
	// Prime batch: 500000/(100+200) = 1666 batches, each with 2 items = 3332 prime items
	// Remainder: 500000 - 1666*300 = 500000 - 499800 = 200 sq ft
	// Non-prime batch: for [50,30] total 80 sq ft per batch
	// 200/80 = 2 batches, each with 2 items = 4 non-prime items

	// Example 2: only prime
	inventory2 := []InventoryItem{
		{1, "prime_eligible", "A", 500.0},
		{2, "prime_eligible", "B", 500.0},
	}
	fmt.Println("Example 2 (only prime):")
	for _, r := range maximizeItems(inventory2) {
		fmt.Printf("  %s %d\n", r.ItemType, r.Count)
	}
	// Prime batch: 500000/(500+500) = 500 batches, each with 2 items = 1000 prime items

	// Example 3: only non-prime
	inventory3 := []InventoryItem{
		{1, "not_prime", "X", 100.0},
	}
	fmt.Println("Example 3 (only non-prime):")
	for _, r := range maximizeItems(inventory3) {
		fmt.Printf("  %s %d\n", r.ItemType, r.Count)
	}
	// Non-prime batch: 500000/100 = 5000 batches, each with 1 item = 5000 items

	// Example 4: empty
	inventory4 := []InventoryItem{}
	fmt.Println("Example 4 (empty):")
	for _, r := range maximizeItems(inventory4) {
		fmt.Printf("  %s %d\n", r.ItemType, r.Count)
	}
}
```
