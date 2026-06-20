# 1501 — Countries You Can Safely Invest In

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func FindSafeCountries(personCountry map[int]string, countryName map[string]string, calls [][3]int) []string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, Trie

**Kompleksitas Waktu:** O(N), Space: O(K) where N = calls, K = countries  
**Kompleksitas Ruang:** O(K) where N = calls, K = countries

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

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
  // Membuat map (HashMap) — pencarian O(1)
	countryDur := make(map[string]int)
  // Membuat map (HashMap) — pencarian O(1)
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
