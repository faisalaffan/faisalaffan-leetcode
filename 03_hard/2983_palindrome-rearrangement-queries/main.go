package main

// LeetCode #2983: Palindrome Rearrangement Queries
// https://leetcode.com/problems/palindrome-rearrangement-queries/
// Difficulty: Hard
//
// Approach:
// Split the string into two halves: left = s[:mid], right = reverse(s[mid:]).
// For each query [a,b,c,d], map the ranges to the first half indices.
// Use prefix sums to count character frequencies and differences.

import (
	"fmt"
)

func canMakePalindromeQueries(s string, queries [][]int) []bool {
	n := len(s)
	mid := n / 2
	a := s[:mid]
	b := reverseStr(s[mid:])
	prefA := make([][26]int, mid+1)
	prefB := make([][26]int, mid+1)
	for i := 0; i < mid; i++ {
		prefA[i+1] = prefA[i]
		prefA[i+1][a[i]-'a']++
		prefB[i+1] = prefB[i]
		prefB[i+1][b[i]-'a']++
	}
	diff := make([]int, mid+1)
	for i := 0; i < mid; i++ {
		diff[i+1] = diff[i]
		if a[i] != b[i] {
			diff[i+1]++
		}
	}
	cntRange := func(pref [][26]int, l, r int) [26]int {
		var res [26]int
		if l > r {
			return res
		}
		for i := 0; i < 26; i++ {
			res[i] = pref[r+1][i] - pref[l][i]
		}
		return res
	}
	ans := make([]bool, len(queries))
	for qi, q := range queries {
		qa, qb, qc, qd := q[0], q[1], q[2], q[3]
		rb := n - 1 - qd
		re := n - 1 - qc
		l1, r1 := qa, qb
		l2, r2 := rb, re
		if l1 > l2 {
			l1, r1, l2, r2 = l2, r2, l1, r1
		}
		if diff[l1] > 0 || diff[mid]-diff[max2(r1, r2)+1] > 0 {
			ans[qi] = false
			continue
		}
		inBoth := max2(l1, l2) <= min2(r1, r2)
		if !inBoth {
			cntA1 := cntRange(prefA, l1, r1)
			cntB1 := cntRange(prefB, l1, r1)
			cntA2 := cntRange(prefA, l2, r2)
			cntB2 := cntRange(prefB, l2, r2)
			ok := true
			for c := 0; c < 26; c++ {
				if cntA1[c] != cntB1[c] || cntA2[c] != cntB2[c] {
					ok = false
					break
				}
			}
			ans[qi] = ok
		} else {
			lBoth := max2(l1, l2)
			rBoth := min2(r1, r2)
			leftL, leftR := l1, l2-1
			rightL, rightR := r2+1, r1
			cntAleft := cntRange(prefA, leftL, leftR)
			cntBleft := cntRange(prefB, leftL, leftR)
			cntAright := cntRange(prefA, rightL, rightR)
			cntBright := cntRange(prefB, rightL, rightR)
			cntBothA := cntRange(prefA, lBoth, rBoth)
			cntBothB := cntRange(prefB, lBoth, rBoth)
			ok := true
			for c := 0; c < 26; c++ {
				if cntAleft[c] != cntBleft[c] || cntAright[c] != cntBright[c] {
					ok = false
					break
				}
			}
			if ok {
				for c := 0; c < 26; c++ {
					if cntBothA[c] != cntBothB[c] {
						ok = false
						break
					}
				}
			}
			ans[qi] = ok
		}
	}
	return ans
}

func reverseStr(s string) string {
	r := []byte(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func min2(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max2(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Example from problem: s="abcabc", queries=[[1,1,3,5],[0,2,5,5]] -> [true,true]
	res := canMakePalindromeQueries("abcabc", [][]int{{1, 1, 3, 5}, {0, 2, 5, 5}})
	fmt.Println(res)
	// Single revision on matching halves
	res2 := canMakePalindromeQueries("ab", [][]int{{0, 0, 1, 1}})
	fmt.Println(res2)
}
