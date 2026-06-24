package main

// LeetCode #242: Valid Anagram
// https://leetcode.com/problems/valid-anagram/
// Difficulty: Easy
//
// LOGIC:
// 1. Jika panjang s != t, langsung return false
// 2. Buat array [26]int sebagai counter karakter a-z
// 3. Iterasi bersamaan: karakter di s → counter++, karakter di t → counter--
// 4. Jika semua counter == 0, s dan t adalah anagram
// 5. Intuisi: karakter yang muncul di s harus dinetralkan oleh karakter yang sama di t
//    Kalau satu string punya kelebihan suatu huruf, counternya tidak akan nol

import "fmt"

// Time: O(n) | Space: O(1) (fixed 26 chars)
func IsAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	count := [26]int{}
	for i := 0; i < len(s); i++ {
		count[s[i]-'a']++
		count[t[i]-'a']--
	}
	return count == [26]int{}
}

func main() {
	fmt.Println(IsAnagram("anagram", "nagaram"))
	fmt.Println(IsAnagram("rat", "car"))
}
