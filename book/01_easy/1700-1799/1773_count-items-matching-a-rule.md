# 1773 — Count Items Matching A Rule

## Deskripsi

**Soal:** [1773. Count Items Matching A Rule](https://leetcode.com/problems/count-items-matching-a-rule/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func CountMatches(items [][]string, ruleKey string, ruleValue string) int`

## Solusi Go

```go
package main

// LeetCode #1773: Count Items Matching a Rule
// https://leetcode.com/problems/count-items-matching-a-rule/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(1)
func CountMatches(items [][]string, ruleKey string, ruleValue string) int {
	idx := 0
	switch ruleKey {
	case "color":
		idx = 1
	case "name":
		idx = 2
	}
	count := 0
	for _, item := range items {
		if item[idx] == ruleValue {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(CountMatches([][]string{{"phone", "blue", "pixel"}, {"computer", "silver", "lenovo"}, {"phone", "gold", "iphone"}}, "color", "silver"))
	fmt.Println(CountMatches([][]string{{"phone", "blue", "pixel"}, {"computer", "silver", "phone"}, {"phone", "gold", "iphone"}}, "type", "phone"))
}
```
