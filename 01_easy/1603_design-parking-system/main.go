package main

// LeetCode #1603: Design Parking System
// https://leetcode.com/problems/design-parking-system/
// Difficulty: Easy

import "fmt"

type ParkingSystem struct {
	spots [3]int
}

func NewParkingSystem(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{spots: [3]int{big, medium, small}}
}

func (this *ParkingSystem) AddCar(carType int) bool {
	if this.spots[carType-1] > 0 {
		this.spots[carType-1]--
		return true
	}
	return false
}

func main() {
	ps := NewParkingSystem(1, 1, 0)
	fmt.Println(ps.AddCar(1))
	fmt.Println(ps.AddCar(2))
	fmt.Println(ps.AddCar(3))
	fmt.Println(ps.AddCar(1))
}
