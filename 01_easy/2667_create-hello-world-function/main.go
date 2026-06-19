package main

// LeetCode #2667: Create Hello World Function
// https://leetcode.com/problems/create-hello-world-function/
// Difficulty: Easy
// Time: O(1) | Space: O(1)
// Note: JavaScript problem, adapted to Go. Returns a function that always returns "Hello World".

import "fmt"

func main() {
	f := CreateHelloWorldFunction()
	fmt.Println(f())
	fmt.Println(f())
}

func CreateHelloWorldFunction() func() string {
	return func() string {
		return "Hello World"
	}
}
