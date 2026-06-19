package main

// LeetCode #1125: Smallest Sufficient Team
// https://leetcode.com/problems/smallest-sufficient-team/
// Difficulty: Hard
//
// Bitmask DP. Map each required skill to a bit position. Each person has a
// mask of their skills. DP[mask] = smallest team (list of person indices)
// that covers that skill mask. Iterate people and update DP.

import "fmt"

func main() {
	reqSkills := []string{"java", "nodejs", "reactjs"}
	people := [][]string{{"java"}, {"nodejs"}, {"nodejs", "reactjs"}}
	fmt.Println(smallestSufficientTeam(reqSkills, people))
}

func smallestSufficientTeam(reqSkills []string, people [][]string) []int {
	skillIdx := make(map[string]int)
	for i, s := range reqSkills {
		skillIdx[s] = i
	}

	m := len(reqSkills)
	fullMask := (1 << m) - 1

	// Convert people skills to bitmasks
	personMask := make([]int, len(people))
	for i, skills := range people {
		mask := 0
		for _, s := range skills {
			if idx, ok := skillIdx[s]; ok {
				mask |= 1 << idx
			}
		}
		personMask[i] = mask
	}

	// dp[mask] = list of person indices covering this mask
	dp := make([][]int, 1<<m)
	dp[0] = []int{} // empty team covers no skills

	for i, pMask := range personMask {
		if pMask == 0 {
			continue
		}
		for mask := 0; mask <= fullMask; mask++ {
			if dp[mask] == nil {
				continue
			}
			newMask := mask | pMask
			if dp[newMask] == nil || len(dp[mask])+1 < len(dp[newMask]) {
				newTeam := make([]int, len(dp[mask]))
				copy(newTeam, dp[mask])
				dp[newMask] = append(newTeam, i)
			}
		}
	}

	return dp[fullMask]
}
