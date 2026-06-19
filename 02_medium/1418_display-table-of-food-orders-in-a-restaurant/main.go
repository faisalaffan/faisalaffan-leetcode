package main

// LeetCode #1418: Display Table of Food Orders in a Restaurant
// https://leetcode.com/problems/display-table-of-food-orders-in-a-restaurant/
// Difficulty: Medium

import "fmt"
import "sort"
import "strconv"

func main() {
	// Test case 1
	fmt.Println(displayTable([][]string{
		{"David", "3", "Ceviche"},
		{"Corina", "10", "Beef Burrito"},
		{"David", "3", "Fried Chicken"},
		{"Carla", "5", "Water"},
		{"Carla", "5", "Ceviche"},
		{"Rous", "3", "Ceviche"},
	}))

	// Test case 2
	fmt.Println(displayTable([][]string{
		{"James", "12", "Fried Chicken"},
		{"Ratesh", "12", "Fried Chicken"},
		{"Amadeus", "12", "Fried Chicken"},
		{"Adam", "1", "Canadian Waffles"},
		{"Brianna", "1", "Canadian Waffles"},
	}))
}

// Time: O(n log n) where n = number of orders
// Space: O(n) for maps
func displayTable(orders [][]string) [][]string {
	foodItems := make(map[string]bool)
	tableOrders := make(map[int]map[string]int)

	for _, o := range orders {
		table, _ := strconv.Atoi(o[1])
		food := o[2]

		foodItems[food] = true
		if tableOrders[table] == nil {
			tableOrders[table] = make(map[string]int)
		}
		tableOrders[table][food]++
	}

	// Sort food items (excluding "Table" header)
	foods := make([]string, 0, len(foodItems))
	for f := range foodItems {
		foods = append(foods, f)
	}
	sort.Strings(foods)

	// Sort table numbers
	tables := make([]int, 0, len(tableOrders))
	for t := range tableOrders {
		tables = append(tables, t)
	}
	sort.Ints(tables)

	// Build result
	result := make([][]string, 0, len(tables)+1)
	header := make([]string, 0, len(foods)+1)
	header = append(header, "Table")
	header = append(header, foods...)
	result = append(result, header)

	for _, t := range tables {
		row := make([]string, 0, len(foods)+1)
		row = append(row, strconv.Itoa(t))
		for _, f := range foods {
			row = append(row, strconv.Itoa(tableOrders[t][f]))
		}
		result = append(result, row)
	}

	return result
}
