package main

// LeetCode #920: Number of Music Playlists
// https://leetcode.com/problems/number-of-music-playlists/
// Difficulty: Hard
// DP[i][j] = number of playlists of length j using exactly i distinct songs.
// DP[i][j] = DP[i-1][j-1] * (n-(i-1)) + DP[i][j-1] * max(0, i-k)
// Then DP[n][goal] is answer.

import "fmt"

func numMusicPlaylists(n int, goal int, k int) int {
	mod := int(1e9 + 7)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, goal+1)
	}
	dp[0][0] = 1

	for i := 1; i <= n; i++ {
		for j := i; j <= goal; j++ {
			// use a new song
			dp[i][j] = (dp[i][j] + dp[i-1][j-1]*(n-(i-1))) % mod
			// reuse an old song (need k different songs before repeating)
			if i > k {
				dp[i][j] = (dp[i][j] + dp[i][j-1]*(i-k)) % mod
			}
		}
	}
	return dp[n][goal]
}

func main() {
	fmt.Println(numMusicPlaylists(3, 3, 1)) // Expected: 6
	fmt.Println(numMusicPlaylists(2, 3, 0)) // Expected: 6
	fmt.Println(numMusicPlaylists(2, 3, 1)) // Expected: 2
	fmt.Println(numMusicPlaylists(1, 3, 0)) // Expected: 1
	fmt.Println(numMusicPlaylists(3, 3, 2)) // Expected: 6
}
