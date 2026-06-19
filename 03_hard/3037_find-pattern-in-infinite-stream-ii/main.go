package main

// LeetCode #3037: Find Pattern in Infinite Stream II
// https://leetcode.com/problems/find-pattern-in-infinite-stream-ii/
// Difficulty: Hard [Paid]

import "fmt"

type InfiniteStream struct {
	bits []int
	pos  int
}

func NewInfiniteStream(bits []int) *InfiniteStream {
	return &InfiniteStream{bits: bits, pos: 0}
}

func (s *InfiniteStream) Next() int {
	val := s.bits[s.pos%len(s.bits)]
	s.pos++
	return val
}

func findPattern(stream *InfiniteStream, pattern []int) int {
	m := len(pattern)
	lps := make([]int, m)
	j := 0
	for i := 1; i < m; i++ {
		for j > 0 && pattern[i] != pattern[j] {
			j = lps[j-1]
		}
		if pattern[i] == pattern[j] {
			j++
		}
		lps[i] = j
	}
	j = 0
	for idx := 0; ; idx++ {
		bit := stream.Next()
		for j > 0 && bit != pattern[j] {
			j = lps[j-1]
		}
		if bit == pattern[j] {
			j++
		}
		if j == m {
			return idx - m + 1
		}
	}
}

func main() {
	stream := NewInfiniteStream([]int{1, 1, 1, 0, 1, 1, 1})
	fmt.Println(findPattern(stream, []int{0, 1}))
}
