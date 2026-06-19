package main

import (
	"fmt"
	"math/rand"
	"sort"
)

// LeetCode #710: Random Pick with Blacklist
// https://leetcode.com/problems/random-pick-with-blacklist/
// Difficulty: Hard
//
// Remap blacklisted numbers in [0, n-len(blacklist)) to non-blacklisted
// numbers in [n-len(blacklist), n). Pick() returns uniform random from
// valid set.

type Solution struct {
	mapping map[int]int
	size    int
}

func Constructor(n int, blacklist []int) Solution {
	blackSet := make(map[int]bool)
	for _, b := range blacklist {
		blackSet[b] = true
	}

	size := n - len(blacklist)
	mapping := make(map[int]int)

	last := size
	for _, b := range blacklist {
		if b < size {
			for last < n {
				if !blackSet[last] {
					mapping[b] = last
					last++
					break
				}
				last++
			}
		}
	}

	return Solution{mapping: mapping, size: size}
}

func (s *Solution) Pick() int {
	r := rand.Intn(s.size)
	if v, ok := s.mapping[r]; ok {
		return v
	}
	return r
}

func main() {
	// n=7, blacklist=[2,3,5] => valid picks: {0,1,4,6}
	s := Constructor(7, []int{2, 3, 5})
	counts := make(map[int]int)
	for i := 0; i < 10000; i++ {
		counts[s.Pick()]++
	}
	for k := range counts {
		if k >= 7 {
			fmt.Printf("ERROR: invalid pick %d\n", k)
		}
	}
	keys := make([]int, 0, len(counts))
	for k := range counts {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		fmt.Printf("%d: %d\n", k, counts[k])
	}
	fmt.Println("---")

	// n=4, blacklist=[] => valid: {0,1,2,3}
	s2 := Constructor(4, []int{})
	counts2 := make(map[int]int)
	for i := 0; i < 10000; i++ {
		counts2[s2.Pick()]++
	}
	keys2 := make([]int, 0, len(counts2))
	for k := range counts2 {
		keys2 = append(keys2, k)
	}
	sort.Ints(keys2)
	for _, k := range keys2 {
		fmt.Printf("%d: %d\n", k, counts2[k])
	}
	fmt.Println("---")

	// n=3, blacklist=[0,1] => valid: {2}
	s3 := Constructor(3, []int{0, 1})
	for i := 0; i < 100; i++ {
		v := s3.Pick()
		if v != 2 {
			fmt.Printf("ERROR: expected 2 got %d\n", v)
		}
	}
	fmt.Println("All picks = 2 (correct)")
}
