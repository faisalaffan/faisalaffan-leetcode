# leetcode

Solusi LeetCode dalam Go — 3962 soal terstruktur per difficulty.

## Struktur

```
.
├── easy/    (950 soal)
├── medium/  (2069 soal)
└── hard/    (943 soal)
```

## Konvensi

- **1 file = 1 soal**, nama: `{4digit_id}_{slug}.go`
- **Package = difficulty** (`easy`, `medium`, `hard`)
- **Function = PascalCase dari slug**: `two-sum` → `TwoSum`, `3sum` → `ThreeSum`, `01-matrix` → `ZeroOneMatrix`
- Metadata soal (link, difficulty) di komentar atas function

## Contoh

```go
// easy/0001_two-sum.go
package easy

// LeetCode #1: Two Sum
// https://leetcode.com/problems/two-sum/
// Difficulty: Easy

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
