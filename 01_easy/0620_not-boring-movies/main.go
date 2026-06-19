package main

// LeetCode #620: Not Boring Movies
// https://leetcode.com/problems/not-boring-movies/
// Difficulty: Easy

import "fmt"

func NotBoringMovies() string {
	return "SELECT * FROM Cinema WHERE id % 2 = 1 AND description != 'boring' ORDER BY rating DESC"
}

func main() {
	fmt.Println(NotBoringMovies())
}
