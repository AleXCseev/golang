package main

import (
	"fmt"
	"math"
)

func main() {
	const IMTPower = 2
	var userHeight float64
	var userKg float64
	fmt.Print("Enter your height: ")
	fmt.Scan(&userHeight)
	fmt.Print("Enter your weight: ")
	fmt.Scan(&userKg)
	IMT := userKg / math.Pow(userHeight / 100, IMTPower)
	fmt.Printf("IMT: %.0f", IMT)
}
