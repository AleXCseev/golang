package main

import (
	"fmt"
	"math"
)

func main() {
	userHeight, userKg := getUserInput()
	IMT := calculateIMT(userHeight, userKg)
	outputResult(IMT)
}

func outputResult(IMT float64) {
	fmt.Printf("IMT: %.0f \n", IMT)
}

func calculateIMT(height float64, weight float64) float64 {
	const IMTPower = 2
	return weight / math.Pow(height / 100, IMTPower)
}

func getUserInput() (float64, float64) {
	var userHeight float64
	var userKg float64
	fmt.Print("Enter your height: ")
	fmt.Scan(&userHeight)
	fmt.Print("Enter your weight: ")
	fmt.Scan(&userKg)
	return userHeight, userKg
}