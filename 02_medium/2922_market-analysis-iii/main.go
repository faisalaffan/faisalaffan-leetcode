package main

// LeetCode #2922: Market Analysis III
// https://leetcode.com/problems/market-analysis-iii/
// Difficulty: Medium (SQL problem — implemented in Go)
// Time: O(n) | Space: O(n)

import (
	"fmt"
	"sort"
)

type User struct {
	SellerID      int
	JoinDate      string
	FavoriteBrand string
}

type Item struct {
	ItemID    int
	ItemBrand string
}

type Order struct {
	OrderID   int
	OrderDate string
	ItemID    int
	SellerID  int
}

type SellerResult struct {
	SellerID int
	NumItems int
}

func marketAnalysisIII(users []User, items []Item, orders []Order) []SellerResult {
	// Build item_id -> item_brand map
	itemBrand := make(map[int]string)
	for _, item := range items {
		itemBrand[item.ItemID] = item.ItemBrand
	}

	// Build seller_id -> favorite_brand map
	sellerBrand := make(map[int]string)
	for _, u := range users {
		sellerBrand[u.SellerID] = u.FavoriteBrand
	}

	// For each seller, count distinct items where item_brand != favorite_brand
	sellerCounts := make(map[int]map[int]bool) // seller_id -> set of item_ids
	for _, o := range orders {
		brand, ok := itemBrand[o.ItemID]
		if !ok {
			continue
		}
		fav, ok := sellerBrand[o.SellerID]
		if !ok {
			continue
		}
		if brand != fav {
			if sellerCounts[o.SellerID] == nil {
				sellerCounts[o.SellerID] = make(map[int]bool)
			}
			sellerCounts[o.SellerID][o.ItemID] = true
		}
	}

	// Find max count
	maxCount := 0
	for _, itemsSet := range sellerCounts {
		if len(itemsSet) > maxCount {
			maxCount = len(itemsSet)
		}
	}

	// Collect sellers with max count
	var results []SellerResult
	for sid, itemsSet := range sellerCounts {
		if len(itemsSet) == maxCount {
			results = append(results, SellerResult{SellerID: sid, NumItems: len(itemsSet)})
		}
	}

	// Order by seller_id ASC
	sort.Slice(results, func(i, j int) bool {
		return results[i].SellerID < results[j].SellerID
	})

	return results
}

func main() {
	users := []User{
		{SellerID: 1, JoinDate: "2024-01-01", FavoriteBrand: "Apple"},
		{SellerID: 2, JoinDate: "2024-01-02", FavoriteBrand: "Samsung"},
		{SellerID: 3, JoinDate: "2024-01-03", FavoriteBrand: "Google"},
	}

	items := []Item{
		{ItemID: 1, ItemBrand: "Apple"},
		{ItemID: 2, ItemBrand: "Samsung"},
		{ItemID: 3, ItemBrand: "Google"},
		{ItemID: 4, ItemBrand: "Apple"},
	}

	orders := []Order{
		{OrderID: 1, OrderDate: "2024-02-01", ItemID: 2, SellerID: 1},  // seller1 buys Samsung != Apple
		{OrderID: 2, OrderDate: "2024-02-02", ItemID: 3, SellerID: 1},  // seller1 buys Google != Apple
		{OrderID: 3, OrderDate: "2024-02-03", ItemID: 1, SellerID: 2},  // seller2 buys Apple != Samsung
		{OrderID: 4, OrderDate: "2024-02-04", ItemID: 4, SellerID: 2},  // seller2 buys Apple != Samsung
		{OrderID: 5, OrderDate: "2024-02-05", ItemID: 3, SellerID: 2},  // seller2 buys Google != Samsung
		{OrderID: 6, OrderDate: "2024-02-06", ItemID: 1, SellerID: 3},  // seller3 buys Apple != Google
		{OrderID: 7, OrderDate: "2024-02-07", ItemID: 2, SellerID: 3},  // seller3 buys Samsung != Google
	}

	fmt.Println("Market Analysis III")
	fmt.Println("===================")
	fmt.Printf("%-12s %s\n", "Seller ID", "Num Items")
	fmt.Println("---------------------")

	results := marketAnalysisIII(users, items, orders)
	for _, r := range results {
		fmt.Printf("%-12d %d\n", r.SellerID, r.NumItems)
	}
}
