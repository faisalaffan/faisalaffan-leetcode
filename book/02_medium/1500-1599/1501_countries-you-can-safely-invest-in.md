# 1501 — Countries You Can Safely Invest In

## Deskripsi

**Soal:** [1501. Countries You Can Safely Invest In](https://leetcode.com/problems/countries-you-can-safely-invest-in/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(N), Space: O(K) where N = calls, K = countries  
**Kompleksitas Ruang:** O(K) where N = calls, K = countries

**Algoritma:** Trie (pohon awalan)

## Solusi Go

```go
package main

// LeetCode #1501: Countries You Can Safely Invest In
// https://leetcode.com/problems/countries-you-can-safely-invest-in/
// Difficulty: Medium [Paid]

import "fmt"

func main() {
	// SQL problem translated to Go: find countries whose average call duration
	// exceeds the global average.
	// Tables: Person(id, name, phone_number), Country(code, name), Calls(caller_id, callee_id, duration)

	// person ID -> country code
	personCountry := map[int]string{1: "US", 2: "US", 3: "UK", 4: "UK", 5: "IN"}
	// country code -> name
	countryName := map[string]string{"US": "USA", "UK": "UK", "IN": "India"}
	// calls: {caller_id, callee_id, duration}
	calls := [][3]int{{1, 2, 10}, {2, 1, 20}, {3, 4, 5}, {4, 3, 15}, {5, 1, 30}}

	result := FindSafeCountries(personCountry, countryName, calls)
	fmt.Println("Safe countries:", result)
}

func FindSafeCountries(personCountry map[int]string, countryName map[string]string, calls [][3]int) []string {
	// Time: O(N), Space: O(K) where N = calls, K = countries
  // Membuat map untuk pencarian O(1): key → value
	countryDur := make(map[string]int)
  // Membuat map untuk pencarian O(1): key → value
	countryCount := make(map[string]int)
	globalDur := 0
	globalCount := 0

	for _, c := range calls {
		caller, callee, dur := c[0], c[1], c[2]
		globalDur += dur
		globalCount++

		callerCode := personCountry[caller]
		calleeCode := personCountry[callee]

		countryDur[callerCode] += dur
		countryCount[callerCode]++
		if callerCode != calleeCode {
			countryDur[calleeCode] += dur
			countryCount[calleeCode]++
		}
	}

	if globalCount == 0 {
		return nil
	}
	globalAvg := float64(globalDur) / float64(globalCount)

  // Membuat slice untuk menyimpan hasil
	result := make([]string, 0)
	for code, dur := range countryDur {
		avg := float64(dur) / float64(countryCount[code])
		if avg > globalAvg {
			if name, ok := countryName[code]; ok {
				result = append(result, name)
			} else {
				result = append(result, code)
			}
		}
	}
	return result
}
```
