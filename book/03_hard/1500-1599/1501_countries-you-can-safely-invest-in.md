# 1501 — Countries You Can Safely Invest In

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func findSafeCountries(persons []Person, countries []Country, calls []Call) []SafeCountry`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Trie

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1501: Countries You Can Safely Invest In
// https://leetcode.com/problems/countries-you-can-safely-invest-in/
// Difficulty: Medium (listed here as Hard) [Paid]
//
// A country is safe to invest in if the average call duration
// of its residents is longer than the global average call duration.

import "fmt"

// Person represents a person record.
type Person struct {
	ID        int
	Name      string
	CountryID int
}

// Country represents a country record.
type Country struct {
	ID   int
	Name string
}

// Call represents a call record.
type Call struct {
	CallerID int
	CalleeID int
	Duration int // in seconds
}

// SafeCountry holds the result.
type SafeCountry struct {
	Name string
}

// findSafeCountries finds countries whose average call duration
// exceeds the global average.
func findSafeCountries(persons []Person, countries []Country, calls []Call) []SafeCountry {
	// Map country ID to country name
  // HashMap: O(1) lookup
	countryMap := make(map[int]string)
	for _, c := range countries {
		countryMap[c.ID] = c.Name
	}

	// Map person ID to country ID
  // HashMap: O(1) lookup
	personCountry := make(map[int]int)
	for _, p := range persons {
		personCountry[p.ID] = p.CountryID
	}

	// Global totals
	var globalTotalDuration int
	var globalTotalCalls int

	// Per-country totals
	type totals struct {
		duration int
		count    int
	}
  // HashMap: O(1) lookup
	countryTotals := make(map[int]*totals)

	for _, c := range calls {
		globalTotalDuration += c.Duration
		globalTotalCalls++

		// Caller's country
		if countryID, ok := personCountry[c.CallerID]; ok {
			if countryTotals[countryID] == nil {
				countryTotals[countryID] = &totals{}
			}
			countryTotals[countryID].duration += c.Duration
			countryTotals[countryID].count++
		}

		// Callee's country
		if countryID, ok := personCountry[c.CalleeID]; ok {
			if countryTotals[countryID] == nil {
				countryTotals[countryID] = &totals{}
			}
			countryTotals[countryID].duration += c.Duration
			countryTotals[countryID].count++
		}
	}

	if globalTotalCalls == 0 {
		return nil
	}
	globalAvg := float64(globalTotalDuration) / float64(globalTotalCalls)

	var result []SafeCountry
	for countryID, t := range countryTotals {
		if t.count == 0 {
			continue
		}
		countryAvg := float64(t.duration) / float64(t.count)
		if countryAvg > globalAvg {
			name := countryMap[countryID]
			result = append(result, SafeCountry{Name: name})
		}
	}

	return result
}

func main() {
	persons := []Person{
		{1, "Alice", 1},
		{2, "Bob", 2},
		{3, "Charlie", 1},
		{4, "David", 3},
		{5, "Eve", 2},
	}

	countries := []Country{
		{1, "USA"},
		{2, "Canada"},
		{3, "Mexico"},
	}

	calls := []Call{
		{1, 2, 300},  // USA->Canada, 300s
		{1, 3, 200},  // USA->USA, 200s
		{2, 4, 100},  // Canada->Mexico, 100s
		{3, 5, 400},  // USA->Canada, 400s
		{4, 5, 50},   // Mexico->Canada, 50s
	}

	// Global avg = (300+200+100+400+50)/5 = 1050/5 = 210
	// USA: (Alice as caller: 300+200=500, Charlie as caller: 400, Alice as callee: none, Charlie as callee: 200) = 1100/3 = 366.67
	// Wait, let me recalculate:
	// USA calls: caller Alice(300+200=500), callee Charlie(200), caller Charlie(400) = 500+200+400=1100, count=3, avg=366.67 > 210 -> SAFE
	// Canada calls: caller Bob(100+50=150), callee Alice(300), callee Bob(50+100=150)... wait, calls are bi-directional?

	// Actually looking at the problem: A call has a caller and callee. The person's country
	// is involved when they are either the caller or callee. Duration counts for the call.
	// Each call contributes its duration to BOTH the caller's country and the callee's country.

	// Call 1: caller=1(USA,300), callee=2(Canada,300)
	// Call 2: caller=1(USA,200), callee=3(USA,200)
	// Call 3: caller=2(Canada,100), callee=4(Mexico,100)
	// Call 4: caller=3(USA,400), callee=5(Canada,400)
	// Call 5: caller=4(Mexico,50), callee=5(Canada,50)

	// Global: 300+200+100+400+50 = 1050, count=5, avg=210

	// USA: calls 1(300),2(200),4(400) = 900, count=3, avg=300
	// Canada: calls 1(300),3(100),4(400),5(50) = 850, count=4, avg=212.5
	// Mexico: calls 3(100),5(50) = 150, count=2, avg=75

	// USA avg 300 > 210 -> Safe!
	// Canada avg 212.5 > 210 -> Safe!
	// Mexico avg 75 < 210 -> Not safe

	safe := findSafeCountries(persons, countries, calls)
	fmt.Println("Safe countries for investment:")
	for _, c := range safe {
		fmt.Printf("  %s\n", c.Name)
	}

	// Test 2: All below average
	persons2 := []Person{
		{1, "Alice", 1},
		{2, "Bob", 1},
	}
	countries2 := []Country{
		{1, "TestLand"},
	}
	calls2 := []Call{
		{1, 2, 10},
	}
	safe2 := findSafeCountries(persons2, countries2, calls2)
	fmt.Printf("\nTest 2 - All below avg: %d (expected 0)\n", len(safe2))

	// Test 3: No calls
	safe3 := findSafeCountries(persons, countries, nil)
	fmt.Printf("Test 3 - No calls: %v (expected [])\n", safe3)
}
```
