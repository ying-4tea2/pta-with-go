// package main

// import (
// 	"fmt"
// 	"math"
// )

// func fitCheck(h int, w int) {
// 	var std float64
// 	std = float64(h-100) * 1.8
// 	if math.Abs(float64(w)-std) < std*0.1 {
// 		fmt.Printf("You are wan mei!\n")
// 	} else if float64(w)-std >= std*0.1 {
// 		fmt.Printf("You are tai pang le!\n")
// 	} else {
// 		fmt.Printf("You are tai shou le!\n")
// 	}
// }

// func main() {
// 	var counter int
// 	var height []int
// 	var weight []int
// 	fmt.Scan(&counter)
// 	for i := 0; i < counter; i++ {
// 		var h, w int
// 		fmt.Scan(&h, &w)
// 		height = append(height, h)
// 		weight = append(weight, w)
// 	}
// 	for i := 0; i < counter; i++ {
// 		fitCheck(height[i], weight[i])
// 	}
// }
