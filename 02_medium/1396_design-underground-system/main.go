package main

// LeetCode #1396: Design Underground System
// https://leetcode.com/problems/design-underground-system/
// Difficulty: Medium

import "fmt"

type UndergroundSystem struct {
	checkins map[int]checkin
	travels  map[string]travel
}

type checkin struct {
	stationName string
	t           int
}

type travel struct {
	totalTime int
	count     int
}

func main() {
	us := Constructor()

	us.CheckIn(45, "Leyton", 3)
	us.CheckIn(32, "Paradise", 8)
	us.CheckIn(27, "Leyton", 10)
	us.CheckOut(45, "Waterloo", 15)
	us.CheckOut(27, "Waterloo", 20)
	us.CheckOut(32, "Cambridge", 22)
	fmt.Println(us.GetAverageTime("Paradise", "Cambridge")) // 14.0
	fmt.Println(us.GetAverageTime("Leyton", "Waterloo"))    // 11.0
	us.CheckIn(10, "Leyton", 24)
	fmt.Println(us.GetAverageTime("Leyton", "Waterloo"))    // 11.0
	us.CheckOut(10, "Waterloo", 38)
	fmt.Println(us.GetAverageTime("Leyton", "Waterloo"))    // 12.0

	us2 := Constructor()
	us2.CheckIn(1, "A", 1)
	us2.CheckIn(2, "A", 2)
	us2.CheckOut(1, "B", 10)
	us2.CheckOut(2, "B", 20)
	fmt.Println(us2.GetAverageTime("A", "B")) // 13.5
}

func Constructor() UndergroundSystem {
	return UndergroundSystem{
		checkins: make(map[int]checkin),
		travels:  make(map[string]travel),
	}
}

func (this *UndergroundSystem) CheckIn(id int, stationName string, t int) {
	this.checkins[id] = checkin{stationName, t}
}

// Time: O(1)
func (this *UndergroundSystem) CheckOut(id int, stationName string, t int) {
	ci := this.checkins[id]
	delete(this.checkins, id)

	key := ci.stationName + "->" + stationName
	tr := this.travels[key]
	tr.totalTime += t - ci.t
	tr.count++
	this.travels[key] = tr
}

// Time: O(1)
func (this *UndergroundSystem) GetAverageTime(startStation string, endStation string) float64 {
	key := startStation + "->" + endStation
	tr := this.travels[key]
	return float64(tr.totalTime) / float64(tr.count)
}
