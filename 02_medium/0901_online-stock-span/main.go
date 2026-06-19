package main

// LeetCode #901: Online Stock Span
// https://leetcode.com/problems/online-stock-span/
// Difficulty: Medium

import "fmt"

type StockSpanner struct {
	stk []pair
}

type pair struct {
	price, span int
}

func Constructor() StockSpanner {
	return StockSpanner{[]pair{}}
}

func (this *StockSpanner) Next(price int) int {
	span := 1
	for len(this.stk) > 0 && this.stk[len(this.stk)-1].price <= price {
		span += this.stk[len(this.stk)-1].span
		this.stk = this.stk[:len(this.stk)-1]
	}
	this.stk = append(this.stk, pair{price, span})
	return span
}

func main() {
	obj := Constructor()
	fmt.Println(obj.Next(100))
	fmt.Println(obj.Next(80))
	fmt.Println(obj.Next(60))
	fmt.Println(obj.Next(70))
	fmt.Println(obj.Next(60))
	fmt.Println(obj.Next(75))
	fmt.Println(obj.Next(85))
}
