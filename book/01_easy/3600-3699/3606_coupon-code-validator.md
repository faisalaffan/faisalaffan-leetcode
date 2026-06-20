# 3606 — Coupon Code Validator

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string berisi tanda kurung: `()`, `[]`, `{}`. Tugasmu adalah memeriksa apakah string tersebut **valid** — setiap kurung buka harus ditutup oleh kurung yang sesuai dalam urutan benar.

Contoh valid: `()[]{}`, `({[]})`. Tidak valid: `(]`, `([)]`.

**Cara berpikir:** Gunakan Stack. Kurung buka → push. Kurung tutup → pop dan cek kecocokan. Di akhir, stack harus kosong.

**Fungsi Solusi:** `func CouponCodeValidator(code []string, businessLine []string, isActive []bool) []string`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Sorting

**Waktu:** O(n log n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

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

  // Custom sort
	sort.Slice(valid, func(i, j int) bool {
		if valid[i].biz != valid[j].biz {
			return bizOrder[valid[i].biz] < bizOrder[valid[j].biz]
		}
		return valid[i].code < valid[j].code
	})

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
