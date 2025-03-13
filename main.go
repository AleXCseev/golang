package main

import (
	"fmt"
	"math"
)

func main() {
	const IMTPower = 2
	userHeight := 1.72
	userKg := 85.0
	IMT := userKg / math.Pow(userHeight, IMTPower)
	fmt.Println("IMT: ", IMT)
}
