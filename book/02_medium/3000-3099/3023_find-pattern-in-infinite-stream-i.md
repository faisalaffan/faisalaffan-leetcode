# 3023 — Find Pattern In Infinite Stream I

## Deskripsi

**Soal:** [3023. Find Pattern In Infinite Stream I](https://leetcode.com/problems/find-pattern-in-infinite-stream-i/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(m) where m = len(pattern) ≤ 100

**Algoritma:** Bitmask (representasi himpunan dengan bit)

**Fungsi Solusi:** `func NewInfiniteStream(data []int) *InfiniteStream`

## Solusi Go

```go
package main

// LeetCode #3023: Find Pattern in Infinite Stream I (PAID)
// https://leetcode.com/problems/find-pattern-in-infinite-stream-i/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(m) where m = len(pattern) ≤ 100

// Given a binary array pattern and an InfiniteStream that yields bits via Next(),
// find the first index (0-indexed) where pattern finishes matching as a
// contiguous subsequence in the stream.

import "fmt"

// InfiniteStream simulates the LeetCode API. Internally backed by a repeating
// slice so we can test without an actual infinite source.
type InfiniteStream struct {
	data []int
	pos  int
}

func NewInfiniteStream(data []int) *InfiniteStream {
	return &InfiniteStream{data: data, pos: 0}
}

// Next returns the next bit. The internal data cycles through the given
// slice to simulate an infinite source. With pattern length ≤ 100 and
// typical data ≤ 10⁴, the pattern will repeat quickly.
func (s *InfiniteStream) Next() int {
	val := s.data[s.pos%len(s.data)]
	s.pos++
	return val
}

func main() {
	// Test: pattern appears starting at index 2
	stream := NewInfiniteStream([]int{0, 1, 0, 1, 0, 1})
	fmt.Println(findPattern(stream, []int{0, 1, 0})) // 2

	// Test: pattern at a later index
	stream2 := NewInfiniteStream([]int{1, 1, 1, 0, 1, 0, 0})
	fmt.Println(findPattern(stream2, []int{1, 0})) // 3

	// Test: single-bit pattern
	stream3 := NewInfiniteStream([]int{1, 0, 0, 0})
	fmt.Println(findPattern(stream3, []int{1})) // 0
}

func findPattern(stream *InfiniteStream, pattern []int) int {
	m := len(pattern)
	if m == 0 {
		return 0
	}

	// Build pattern bitmask (pattern[0] is shifted into the most
	// significant position of the mask).
	patternMask := 0
	for _, bit := range pattern {
		patternMask = (patternMask << 1) | bit
	}

	mask := (1 << m) - 1
	window := 0
	idx := 0
	// Safety upper bound: the pattern must appear at some point in an
	// underlying periodic sequence; we scan up to 10⁶ positions.
	maxIter := 1_000_000

	for idx < maxIter {
		bit := stream.Next()
		window = ((window << 1) | bit) & mask
		if idx >= m-1 && window == patternMask {
			return idx
		}
		idx++
	}
	return -1
}
```
