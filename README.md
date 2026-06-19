# leetcode

Solusi LeetCode dalam Go — 3962 soal terstruktur per difficulty.

## Struktur

```
.
├── 01_easy/     (950 soal)
├── 02_medium/   (2069 soal)
└── 03_hard/     (943 soal)
```

Setiap soal dalam folder sendiri dengan `main.go` — bisa di-debug satu per satu.

## Cara Debug

```bash
go run ./01_easy/0001_two-sum/
go run ./02_medium/0015_3sum/
go run ./03_hard/0679_24-game/
```

## Konvensi

- **Folder**: `{4digit_id}_{slug}/main.go`
- **Package**: `main` — setiap soal bisa langsung `go run`
- **Function**: PascalCase dari slug: `two-sum` → `TwoSum`, `3sum` → `ThreeSum`, `01-matrix` → `ZeroOneMatrix`
- Metadata soal (link, difficulty) di komentar atas function

## Contoh

```go
// 01_easy/0001_two-sum/main.go
package main

import "fmt"

// LeetCode #1: Two Sum
// https://leetcode.com/problems/two-sum/
// Difficulty: Easy

func main() {
	fmt.Println(TwoSum([]int{2, 7, 11, 15}, 9))
}

func TwoSum(nums []int, target int) []int {
	seen := make(map[int]int)
	for i, n := range nums {
		if j, ok := seen[target-n]; ok {
			return []int{j, i}
		}
		seen[n] = i
	}
	return nil
}
```
