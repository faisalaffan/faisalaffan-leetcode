# 0780 — Reaching Points

## Deskripsi

**Soal:** [0780. Reaching Points](https://leetcode.com/problems/reaching-points/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

> **Ide Kunci:** Reverse modulo

## Solusi Go

```go
package main

// LeetCode #780: Reaching Points
// https://leetcode.com/problems/reaching-points/
// Difficulty: Hard
//
// From (x,y) can go to (x+y, y) or (x, x+y). Determine if (tx,ty)
// is reachable from (sx,sy).
//
// Approach: Reverse modulo
// Work backwards: if tx > ty, previous x = tx % ty (with adjustment
// for the starting constraint). Use modulo instead of repeated
// subtraction for efficiency.

import "fmt"

func main() {
	fmt.Println(reachingPoints(1, 1, 3, 5)) // true
	fmt.Println(reachingPoints(1, 1, 2, 2)) // false
	fmt.Println(reachingPoints(1, 1, 1, 1)) // true
	fmt.Println(reachingPoints(3, 3, 12, 9)) // true
	fmt.Println(reachingPoints(1, 2, 3, 5)) // true
}

func reachingPoints(sx, sy, tx, ty int) bool {
	for tx > sx && ty > sy {
		if tx > ty {
			tx %= ty
		} else {
			ty %= tx
		}
	}

	// Now one coordinate equals the start value
	if tx == sx {
		// Need (ty - sy) >= 0 and divisible by sx
		return (ty-sy) >= 0 && (ty-sy)%sx == 0
	}
	if ty == sy {
		return (tx-sx) >= 0 && (tx-sx)%sy == 0
	}

	return false
}
```
