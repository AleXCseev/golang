package main

import (
	"fmt"
	"math"
)

const IMTPower = 2

func main() {
	userHeight, userKg := getUserInput()
	IMT := calculateIMT(userHeight, userKg)
	outputResult(IMT)

	if IMT < 16 {
		fmt.Println("Сильный дефицит массы тела")
	} else if IMT < 18.5 {
		fmt.Println("Недостаточная масса тела")
	} else if IMT < 25 {
		fmt.Println("Норма")
	} else if IMT < 30 {
		fmt.Println("Избыточная масса тела")
	} else if IMT < 35 {
		fmt.Println("Ожирение первой степени")
	} else if IMT < 40 {
		fmt.Println("Ожирение второй степени")
	} else if IMT >= 40 {
		fmt.Println("Ожирение третьей степени")
	} else {
		fmt.Println("Неверные данные")
		return
	}
}

func outputResult(IMT float64) {
	fmt.Printf("IMT: %.0f \n", IMT)
}

func calculateIMT(height float64, weight float64) float64 {
	return weight / math.Pow(height/100, IMTPower)
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
