package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// LeetCode #1169: Invalid Transactions
// https://leetcode.com/problems/invalid-transactions/
// Difficulty: Medium

// A transaction is invalid if: amount > 1000, or same name within
// 60 min of another transaction at different city.

// Time: O(n^2) due to pairwise comparison
// Space: O(n)

type Transaction struct {
	name   string
	time   int
	amount int
	city   string
	idx    int
}

func invalidTransactions(transactions []string) []string {
	n := len(transactions)
	txs := make([]Transaction, n)

	for i, t := range transactions {
		parts := strings.Split(t, ",")
		tm, _ := strconv.Atoi(parts[1])
		amt, _ := strconv.Atoi(parts[2])
		txs[i] = Transaction{parts[0], tm, amt, parts[3], i}
	}

	invalid := make([]bool, n)

	for i := 0; i < n; i++ {
		if txs[i].amount > 1000 {
			invalid[txs[i].idx] = true
		}
		for j := i + 1; j < n; j++ {
			if txs[i].name == txs[j].name &&
				abs(txs[i].time-txs[j].time) <= 60 &&
				txs[i].city != txs[j].city {
				invalid[txs[i].idx] = true
				invalid[txs[j].idx] = true
			}
		}
	}

	result := make([]string, 0)
	for i, v := range invalid {
		if v {
			result = append(result, transactions[i])
		}
	}
	sort.Strings(result)
	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	t1 := []string{"alice,20,800,mtv", "alice,50,100,beijing"}
	fmt.Printf("%v (expected: [alice,20,800,mtv alice,50,100,beijing])\n", invalidTransactions(t1))

	t2 := []string{"alice,20,800,mtv", "bob,50,1200,mtv"}
	fmt.Printf("%v (expected: [bob,50,1200,mtv])\n", invalidTransactions(t2))

	t3 := []string{"alice,20,800,mtv", "alice,50,100,mtv"}
	fmt.Printf("%v (expected: [])\n", invalidTransactions(t3))
}
