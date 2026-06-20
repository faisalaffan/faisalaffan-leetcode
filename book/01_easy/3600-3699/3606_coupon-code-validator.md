# 3606 — Coupon Code Validator

## Deskripsi

**Soal:** [3606. Coupon Code Validator](https://leetcode.com/problems/coupon-code-validator/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #3606: Coupon Code Validator
// https://leetcode.com/problems/coupon-code-validator/
// Difficulty: Easy

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(CouponCodeValidator(
		[]string{"SAVE20", "", "PHARMA5", "SAVE@20"},
		[]string{"restaurant", "grocery", "pharmacy", "restaurant"},
		[]bool{true, true, true, true},
	))
}

// Time: O(n log n)
// Space: O(n)
func CouponCodeValidator(code []string, businessLine []string, isActive []bool) []string {
	bizOrder := map[string]int{
		"electronics": 0,
		"grocery":     1,
		"pharmacy":    2,
		"restaurant":  3,
	}

	type entry struct {
		code string
		biz  string
	}
	var valid []entry

	for i, c := range code {
		if !isActive[i] {
			continue
		}
		if _, ok := bizOrder[businessLine[i]]; !ok {
			continue
		}
		if c == "" {
			continue
		}
		ok := true
		for _, ch := range c {
			if !isAlphanumeric(ch) && ch != '_' {
				ok = false
				break
			}
		}
		if ok {
			valid = append(valid, entry{c, businessLine[i]})
		}
	}

	sort.Slice(valid, func(i, j int) bool {
		if valid[i].biz != valid[j].biz {
			return bizOrder[valid[i].biz] < bizOrder[valid[j].biz]
		}
		return valid[i].code < valid[j].code
	})

  // Membuat slice untuk menyimpan hasil
	res := make([]string, len(valid))
	for i, e := range valid {
		res[i] = e.code
	}
	return res
}

func isAlphanumeric(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || (ch >= '0' && ch <= '9')
}
```
