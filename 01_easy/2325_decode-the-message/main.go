package main

// LeetCode #2325: Decode the Message
// https://leetcode.com/problems/decode-the-message/
// Difficulty: Easy
// Time O(n + m) | Space O(1)

import "fmt"

func main() {
	fmt.Println(DecodeTheMessage("the quick brown fox jumps over the lazy dog", "vkbs bs t suepuv")) // "this is a secret"
	fmt.Println(DecodeTheMessage("eljuxhpwnyrdgtqkviszcfmabo", "zwx hnfx lqantp mnoeius ycgk vcnjrdb")) // "the five boxing wizards jump quickly"
}

func DecodeTheMessage(key string, message string) string {
	mapping := make([]byte, 26)
	idx := byte(0)

	for i := 0; i < len(key); i++ {
		if key[i] != ' ' && mapping[key[i]-'a'] == 0 {
			mapping[key[i]-'a'] = 'a' + idx
			idx++
		}
	}

	res := make([]byte, len(message))
	for i := 0; i < len(message); i++ {
		if message[i] == ' ' {
			res[i] = ' '
		} else {
			res[i] = mapping[message[i]-'a']
		}
	}
	return string(res)
}
