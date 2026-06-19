package main

// LeetCode #1159: Market Analysis II
// https://leetcode.com/problems/market-analysis-ii/
// Difficulty: Hard [Paid]
//
// For each user, find their second purchase (by order date) and compare the
// item's brand to the user's favorite brand. Return users whose second
// purchase item has a different brand than their favorite brand.

import (
	"fmt"
	"sort"
)

// User represents a user with their favorite brand.
type User struct {
	UserID        int
	FavoriteBrand string
}

// OrderItem represents a single item in an order.
type OrderItem struct {
	OrderID  int
	ItemID   int
	Brand    string // not in raw schema but joined from Item
}

// SaleOrder represents a purchase order.
type SaleOrder struct {
	OrderID int
	UserID  int
	Date    string // "YYYY-MM-DD"
	ItemID  int
	Brand   string // denormalized for simplicity
}

// UserResult is the output format: user_id and whether they bought a different brand.
type UserResult struct {
	UserID          int
	DifferentBrand  bool // true if second purchase brand != favorite brand
}

func main() {
	users := []User{
		{1, "A"},
		{2, "B"},
		{3, "C"},
	}

	orders := []SaleOrder{
		{1, 1, "2019-01-01", 101, "A"},
		{2, 1, "2019-02-01", 102, "B"}, // second purchase: brand B != fav A
		{3, 2, "2019-01-01", 103, "B"},
		{4, 2, "2019-02-01", 104, "B"}, // second purchase: brand B == fav B
		{5, 3, "2019-01-01", 105, "C"},
		// only one purchase for user 3, no second purchase
	}

	results := getMarketAnalysis(users, orders)
	for _, r := range results {
		fmt.Printf("user_id=%d, different_brand=%t\n", r.UserID, r.DifferentBrand)
	}
}

// getMarketAnalysis returns for each user whether their second purchase's brand
// differs from their favorite brand. Users without a second purchase are excluded.
func getMarketAnalysis(users []User, orders []SaleOrder) []UserResult {
	// Build favorite brand lookup
	favBrand := make(map[int]string)
	for _, u := range users {
		favBrand[u.UserID] = u.FavoriteBrand
	}

	// Group orders by user
	userOrders := make(map[int][]SaleOrder)
	for _, o := range orders {
		userOrders[o.UserID] = append(userOrders[o.UserID], o)
	}

	// Sort each user's orders by date
	for uid := range userOrders {
		sort.Slice(userOrders[uid], func(i, j int) bool {
			return userOrders[uid][i].Date < userOrders[uid][j].Date
		})
	}

	var results []UserResult
	userIDs := make([]int, 0, len(users))
	for _, u := range users {
		userIDs = append(userIDs, u.UserID)
	}
	sort.Ints(userIDs)

	for _, uid := range userIDs {
		orders := userOrders[uid]
		if len(orders) < 2 {
			continue
		}
		secondBrand := orders[1].Brand
		results = append(results, UserResult{
			UserID:         uid,
			DifferentBrand: secondBrand != favBrand[uid],
		})
	}

	return results
}
