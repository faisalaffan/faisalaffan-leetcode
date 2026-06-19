package main

// LeetCode #1472: Design Browser History
// https://leetcode.com/problems/design-browser-history/
// Difficulty: Medium

import "fmt"

type BrowserHistory struct {
	history []string
	current int
}

func main() {
	bh := NewBrowserHistory("leetcode.com")
	bh.Visit("google.com")
	bh.Visit("facebook.com")
	bh.Visit("youtube.com")
	fmt.Println(bh.Back(1)) // "facebook.com"
	fmt.Println(bh.Back(1)) // "google.com"
	fmt.Println(bh.Forward(1)) // "facebook.com"
	bh.Visit("linkedin.com")
	fmt.Println(bh.Forward(2)) // "linkedin.com"
	fmt.Println(bh.Back(2)) // "google.com"
	fmt.Println(bh.Back(7)) // "leetcode.com"

	bh2 := NewBrowserHistory("a.com")
	bh2.Visit("b.com")
	fmt.Println(bh2.Back(1)) // "a.com"
	fmt.Println(bh2.Forward(1)) // "b.com"
}

func NewBrowserHistory(homepage string) BrowserHistory {
	return BrowserHistory{
		history: []string{homepage},
		current: 0,
	}
}

// Time: O(1)
func (this *BrowserHistory) Visit(url string) {
	this.current++
	this.history = this.history[:this.current]
	this.history = append(this.history, url)
}

// Time: O(1)
func (this *BrowserHistory) Back(steps int) string {
	this.current -= steps
	if this.current < 0 {
		this.current = 0
	}
	return this.history[this.current]
}

// Time: O(1)
func (this *BrowserHistory) Forward(steps int) string {
	this.current += steps
	if this.current >= len(this.history) {
		this.current = len(this.history) - 1
	}
	return this.history[this.current]
}
