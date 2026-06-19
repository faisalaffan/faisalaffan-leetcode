package main

// LeetCode #1479: Sales by Day of the Week
// https://leetcode.com/problems/sales-by-day-of-the-week/
// Difficulty: Hard [Paid]
//
// Calculate total sales for each item, each day of the week.

import "fmt"
import "time"

// Sale represents a sales record.
type Sale struct {
	ItemID   int
	Date     string // "YYYY-MM-DD"
	Quantity int
	Price    float64
}

// DaySales represents sales summary for an item on a day of the week.
type DaySales struct {
	ItemID int
	Day    string // Monday, Tuesday, etc.
	Total  float64
}

// calculateDaySales computes total sales by item and day of the week.
func calculateDaySales(sales []Sale) []DaySales {
	// We need timezone for date parsing
	loc := time.UTC

	// Aggregate: itemID -> dayOfWeek -> total
	type key struct {
		itemID int
		day    string
	}
	aggregate := make(map[key]float64)

	daysOrder := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

	for _, s := range sales {
		t, err := time.ParseInLocation("2006-01-02", s.Date, loc)
		if err != nil {
			continue
		}
		dayName := t.Weekday().String()
		// Normalize to "Monday", "Tuesday", etc.
		// Go's Weekday() returns Sunday=0, Monday=1, ..., Saturday=6
		aggregate[key{s.ItemID, dayName}] += float64(s.Quantity) * s.Price
	}

	// Collect results
	var results []DaySales
	for k, total := range aggregate {
		results = append(results, DaySales{k.itemID, k.day, total})
	}

	// Sort: by ItemID ASC, then by day of week order
	dayRank := make(map[string]int)
	for i, d := range daysOrder {
		dayRank[d] = i
	}
	// Also handle Go's Weekday() output: "Sunday", "Monday", etc.
	goDayRank := map[string]int{
		"Sunday":    0,
		"Monday":    1,
		"Tuesday":   2,
		"Wednesday": 3,
		"Thursday":  4,
		"Friday":    5,
		"Saturday":  6,
	}

	_ = dayRank // keep for reference

	for i := 0; i < len(results); i++ {
		for j := i + 1; j < len(results); j++ {
			swap := false
			if results[i].ItemID > results[j].ItemID {
				swap = true
			} else if results[i].ItemID == results[j].ItemID {
				if goDayRank[results[i].Day] > goDayRank[results[j].Day] {
					swap = true
				}
			}
			if swap {
				results[i], results[j] = results[j], results[i]
			}
		}
	}

	return results
}

func main() {
	sales := []Sale{
		{1, "2023-01-02", 5, 10.00},   // Monday
		{1, "2023-01-03", 3, 10.00},   // Tuesday
		{1, "2023-01-09", 2, 10.00},   // Monday
		{2, "2023-01-02", 4, 20.00},   // Monday
		{2, "2023-01-04", 1, 20.00},   // Wednesday
	}

	results := calculateDaySales(sales)
	fmt.Println("Sales by Day of Week:")
	for _, r := range results {
		fmt.Printf("  Item %d, %s: $%.2f\n", r.ItemID, r.Day, r.Total)
	}

	// Test 2: Weekend sales
	sales2 := []Sale{
		{1, "2023-01-07", 10, 5.00},   // Saturday
		{1, "2023-01-08", 8, 5.00},    // Sunday
		{2, "2023-01-07", 3, 15.00},   // Saturday
	}
	results2 := calculateDaySales(sales2)
	fmt.Println("\nTest 2 - Weekend sales:")
	for _, r := range results2 {
		fmt.Printf("  Item %d, %s: $%.2f\n", r.ItemID, r.Day, r.Total)
	}

	// Test 3: Empty
	results3 := calculateDaySales(nil)
	fmt.Printf("\nTest 3 - Empty: %d results\n", len(results3))

	// Test 4: Single sale
	sales4 := []Sale{
		{1, "2023-06-01", 2, 25.50},  // Thursday
	}
	results4 := calculateDaySales(sales4)
	fmt.Println("\nTest 4 - Single sale:")
	for _, r := range results4 {
		fmt.Printf("  Item %d, %s: $%.2f\n", r.ItemID, r.Day, r.Total)
	}
}
