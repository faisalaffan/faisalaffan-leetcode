package main

// LeetCode #2043: Simple Bank System
// https://leetcode.com/problems/simple-bank-system/
// Difficulty: Medium
// Time: O(1) per operation | Space: O(n)

import "fmt"

type Bank struct {
	balance []int64
}

func Constructor(balance []int64) Bank {
	return Bank{balance: balance}
}

func (b *Bank) Transfer(account1 int, account2 int, money int64) bool {
	if account1 < 1 || account1 > len(b.balance) || account2 < 1 || account2 > len(b.balance) {
		return false
	}
	if b.balance[account1-1] < money {
		return false
	}
	b.balance[account1-1] -= money
	b.balance[account2-1] += money
	return true
}

func (b *Bank) Deposit(account int, money int64) bool {
	if account < 1 || account > len(b.balance) {
		return false
	}
	b.balance[account-1] += money
	return true
}

func (b *Bank) Withdraw(account int, money int64) bool {
	if account < 1 || account > len(b.balance) {
		return false
	}
	if b.balance[account-1] < money {
		return false
	}
	b.balance[account-1] -= money
	return true
}

func main() {
	bank := Constructor([]int64{10, 100, 20, 50, 30})
	fmt.Println("Test 1 Deposit(3, 10):", bank.Deposit(3, 10))   // true
	fmt.Println("Test 2 Transfer(5, 1, 20):", bank.Transfer(5, 1, 20)) // true
	fmt.Println("Test 3 Withdraw(5, 20):", bank.Withdraw(5, 20)) // true
	fmt.Println("Test 4 Transfer(3, 4, 15):", bank.Transfer(3, 4, 15)) // false (insufficient)
	fmt.Println("Test 5 Deposit(9, 10):", bank.Deposit(9, 10))   // false
}
