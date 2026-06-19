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
	userProductQty := make(map[[2]int]int) // [user_id, product_id] -> total qty

	for _, s := range sales {
		productID, userID, qty := s[1], s[2], s[3]
		key := [2]int{userID, productID}
		userProductQty[key] += qty
	}

	// For each user, find product with max qty
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
