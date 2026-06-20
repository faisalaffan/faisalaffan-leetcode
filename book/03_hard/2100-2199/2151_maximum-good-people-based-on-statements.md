# 2151 — Maximum Good People Based On Statements

## Deskripsi

**Soal:** [2151. Maximum Good People Based On Statements](https://leetcode.com/problems/maximum-good-people-based-on-statements/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Bitmask (representasi himpunan dengan bit)

## Solusi Go

```go
package main

// LeetCode #2151: Maximum Good People Based on Statements
// https://leetcode.com/problems/maximum-good-people-based-on-statements/
// Difficulty: Hard
//
// Bitmask enumeration: try all 2^n assignments of good (1) / bad (0) people.
// For each assignment, verify all non-2 statements from good people.

import "fmt"

func main() {
	fmt.Println(maximumGood([][]int{{2, 1, 2}, {1, 2, 2}, {2, 0, 2}})) // 2
	fmt.Println(maximumGood([][]int{{2, 0}, {0, 2}}))                  // 1
	fmt.Println(maximumGood([][]int{{2, 2, 2}, {2, 2, 2}, {2, 2, 2}})) // 3
}

func maximumGood(statements [][]int) int {
	n := len(statements)
	best := 0

	for mask := 0; mask < (1 << n); mask++ {
		valid := true
		cnt := 0
		for i := 0; i < n && valid; i++ {
			if mask&(1<<i) == 0 {
				continue
			}
			cnt++
			for j := 0; j < n; j++ {
				st := statements[i][j]
				if st == 2 {
					continue
				}
				isGoodJ := (mask >> j) & 1
				if st != isGoodJ {
					valid = false
					break
				}
			}
		}
		if valid && cnt > best {
			best = cnt
		}
	}
	return best
}

func MaximumGoodPeopleBasedOnStatements() any {
	return maximumGood([][]int{{2, 1, 2}, {1, 2, 2}, {2, 0, 2}})
}
```
