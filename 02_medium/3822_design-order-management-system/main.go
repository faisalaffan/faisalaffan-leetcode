package main

// LeetCode #3822: Design Order Management System
// https://leetcode.com/problems/design-order-management-system/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

type Order struct {
	ID    int
	Type  string
	Price int
	Active bool
}

type OrderManagementSystem struct {
	orders      map[int]*Order
	priceLookup map[int]map[string]map[int]struct{}
}

func Constructor() OrderManagementSystem {
	return OrderManagementSystem{
		orders:      make(map[int]*Order),
		priceLookup: make(map[int]map[string]map[int]struct{}),
	}
}

func (o *OrderManagementSystem) AddOrder(orderID int, orderType string, price int) {
	o.orders[orderID] = &Order{ID: orderID, Type: orderType, Price: price, Active: true}

	if o.priceLookup[price] == nil {
		o.priceLookup[price] = make(map[string]map[int]struct{})
	}
	if o.priceLookup[price][orderType] == nil {
		o.priceLookup[price][orderType] = make(map[int]struct{})
	}
	o.priceLookup[price][orderType][orderID] = struct{}{}
}

func (o *OrderManagementSystem) ModifyOrder(orderID int, newPrice int) {
	order := o.orders[orderID]
	// Remove from old price bucket
	delete(o.priceLookup[order.Price][order.Type], orderID)
	if len(o.priceLookup[order.Price][order.Type]) == 0 {
		delete(o.priceLookup[order.Price], order.Type)
	}
	if len(o.priceLookup[order.Price]) == 0 {
		delete(o.priceLookup, order.Price)
	}

	// Add to new price bucket
	order.Price = newPrice
	if o.priceLookup[newPrice] == nil {
		o.priceLookup[newPrice] = make(map[string]map[int]struct{})
	}
	if o.priceLookup[newPrice][order.Type] == nil {
		o.priceLookup[newPrice][order.Type] = make(map[int]struct{})
	}
	o.priceLookup[newPrice][order.Type][orderID] = struct{}{}
}

func (o *OrderManagementSystem) CancelOrder(orderID int) {
	order := o.orders[orderID]
	order.Active = false

	delete(o.priceLookup[order.Price][order.Type], orderID)
	if len(o.priceLookup[order.Price][order.Type]) == 0 {
		delete(o.priceLookup[order.Price], order.Type)
	}
	if len(o.priceLookup[order.Price]) == 0 {
		delete(o.priceLookup, order.Price)
	}
}

func (o *OrderManagementSystem) GetOrdersAtPrice(orderType string, price int) []int {
	if o.priceLookup[price] == nil || o.priceLookup[price][orderType] == nil {
		return []int{}
	}
	result := make([]int, 0, len(o.priceLookup[price][orderType]))
	for id := range o.priceLookup[price][orderType] {
		result = append(result, id)
	}
	sort.Ints(result)
	return result
}

func main() {
	oms := Constructor()
	oms.AddOrder(1, "buy", 1)
	oms.AddOrder(2, "buy", 1)
	oms.AddOrder(3, "sell", 2)

	fmt.Println(oms.GetOrdersAtPrice("buy", 1)) // [1, 2]

	oms.ModifyOrder(1, 3)
	oms.ModifyOrder(2, 1)
	fmt.Println(oms.GetOrdersAtPrice("buy", 1)) // [2]

	oms.CancelOrder(3)
	oms.CancelOrder(2)
	fmt.Println(oms.GetOrdersAtPrice("buy", 1)) // []
}
