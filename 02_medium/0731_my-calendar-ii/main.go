package main

// LeetCode #731: My Calendar II
// https://leetcode.com/problems/my-calendar-ii/
// Difficulty: Medium
// Time: O(n^2) per booking
// Space: O(n)

import "fmt"

func main() {
	cal := ConstructorCalendar2()
	fmt.Println(cal.Book(10, 20))
	fmt.Println(cal.Book(50, 60))
	fmt.Println(cal.Book(10, 40))
	fmt.Println(cal.Book(5, 15))
	fmt.Println(cal.Book(5, 10))
	fmt.Println(cal.Book(25, 55))
}

type MyCalendarTwo struct {
	books    [][2]int
	overlaps [][2]int
}

func ConstructorCalendar2() MyCalendarTwo {
	return MyCalendarTwo{}
}

func (c *MyCalendarTwo) Book(start int, end int) bool {
	for _, o := range c.overlaps {
		if max(o[0], start) < min(o[1], end) {
			return false
		}
	}
	for _, b := range c.books {
		if max(b[0], start) < min(b[1], end) {
			c.overlaps = append(c.overlaps, [2]int{max(b[0], start), min(b[1], end)})
		}
	}
	c.books = append(c.books, [2]int{start, end})
	return true
}
