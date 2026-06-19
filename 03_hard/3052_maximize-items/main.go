package main

// LeetCode #3052: Maximize Items (SQL simulation)
// https://leetcode.com/problems/maximize-items/
// Difficulty: Hard [Paid]

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
	inventory := []InventoryItem{
		{1, "prime_eligible", "Watches", 100.0},
		{2, "prime_eligible", "Art", 200.0},
		{3, "not_prime", "Books", 50.0},
		{4, "not_prime", "Toys", 30.0},
	}
	for _, r := range maximizeItems(inventory) {
		fmt.Printf("%s %d\n", r.ItemType, r.Count)
	}
}
