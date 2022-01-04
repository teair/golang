package main

import (
	"fmt"
)

func numWaterBottles(numBottles int, numExchange int) int {

	var drinkBottles = numBottles

	for numBottles >= numExchange {

		numBottles = numBottles - numExchange + 1

		drinkBottles++

	}

	return drinkBottles
}

func main() {
	fmt.Println(numWaterBottles(9, 3))
}
