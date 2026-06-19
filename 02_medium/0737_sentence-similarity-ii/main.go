package main

// LeetCode #737: Sentence Similarity II
// https://leetcode.com/problems/sentence-similarity-ii/
// Difficulty: Medium [Paid]
// Time: O(n * alpha(n))
// Space: O(n)

import "fmt"

func main() {
	pairs := [][]string{
		{"great", "fine"},
		{"drama", "acting"},
		{"fine", "good"},
	}
	fmt.Println(areSentencesSimilarTwo([]string{"great", "acting", "skills"}, []string{"fine", "drama", "talent"}, pairs))
}

func areSentencesSimilarTwo(words1 []string, words2 []string, pairs [][]string) bool {
	if len(words1) != len(words2) {
		return false
	}

	parent := make(map[string]string)

	var find func(x string) string
	find = func(x string) string {
		if _, ok := parent[x]; !ok {
			parent[x] = x
		}
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}

	union := func(x, y string) {
		px, py := find(x), find(y)
		if px != py {
			parent[px] = py
		}
	}

	for _, p := range pairs {
		union(p[0], p[1])
	}

	for i := 0; i < len(words1); i++ {
		if words1[i] == words2[i] {
			continue
		}
		if find(words1[i]) != find(words2[i]) {
			return false
		}
	}

	return true
}
