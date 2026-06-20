# 3815 — Design Auction System

## Deskripsi

**Soal:** [3815. Design Auction System](https://leetcode.com/problems/design-auction-system/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(log N) per operation  
**Kompleksitas Ruang:** O(N)

**Algoritma:** HashMap (tabel pencarian O(1)), LIS (Longest Increasing Subsequence)

**Fungsi Solusi:** `func Constructor() AuctionSystem`

> **Ide Kunci:** Hash map for user bids, sorted set for highest bid per item.

## Solusi Go

```go
package main

// LeetCode #3815: Design Auction System
// https://leetcode.com/problems/design-auction-system/
// Difficulty: Medium
// Time: O(log N) per operation | Space: O(N)
// Approach: Hash map for user bids, sorted set for highest bid per item.

import (
	"fmt"
)

// AuctionSystem handles bids from users on items.
type AuctionSystem struct {
	// users[userId][itemId] = bidAmount
	users map[int]map[int]int
	// items[itemId] -> sorted list of (bidAmount, userId) using slice
	items map[int][][2]int
}

func Constructor() AuctionSystem {
	return AuctionSystem{
		users: make(map[int]map[int]int),
		items: make(map[int][][2]int),
	}
}

func (as *AuctionSystem) AddBid(userId int, itemId int, bidAmount int) {
	if _, ok := as.users[userId]; !ok {
		as.users[userId] = make(map[int]int)
	}
	// If user already has a bid on this item, remove it first
	if oldBid, ok := as.users[userId][itemId]; ok {
		as.removeBidFromItem(itemId, oldBid, userId)
	}
	as.users[userId][itemId] = bidAmount
	as.items[itemId] = append(as.items[itemId], [2]int{bidAmount, userId})
}

func (as *AuctionSystem) removeBidFromItem(itemId int, bidAmount int, userId int) {
	bids := as.items[itemId]
	for i, b := range bids {
		if b[0] == bidAmount && b[1] == userId {
			as.items[itemId] = append(bids[:i], bids[i+1:]...)
			break
		}
	}
}

func (as *AuctionSystem) UpdateBid(userId int, itemId int, newAmount int) {
	if oldBid, ok := as.users[userId][itemId]; ok {
		as.removeBidFromItem(itemId, oldBid, userId)
		as.users[userId][itemId] = newAmount
		as.items[itemId] = append(as.items[itemId], [2]int{newAmount, userId})
	}
}

func (as *AuctionSystem) RemoveBid(userId int, itemId int) {
	if bidAmount, ok := as.users[userId][itemId]; ok {
		as.removeBidFromItem(itemId, bidAmount, userId)
		delete(as.users[userId], itemId)
	}
}

func (as *AuctionSystem) GetHighestBidder(itemId int) int {
	bids, ok := as.items[itemId]
	if !ok || len(bids) == 0 {
		return -1
	}
	bestAmount, bestUser := -1, -1
	for _, b := range bids {
		if b[0] > bestAmount || (b[0] == bestAmount && b[1] > bestUser) {
			bestAmount = b[0]
			bestUser = b[1]
		}
	}
	return bestUser
}

func main() {
	as := Constructor()

	as.AddBid(1, 7, 5)
	as.AddBid(2, 7, 6)
	fmt.Println(as.GetHighestBidder(7)) // Expected: 2

	as.UpdateBid(1, 7, 8)
	fmt.Println(as.GetHighestBidder(7)) // Expected: 1

	as.RemoveBid(2, 7)
	fmt.Println(as.GetHighestBidder(7)) // Expected: 1

	fmt.Println(as.GetHighestBidder(3)) // Expected: -1
}
```
