// package main

// import (
// 	"fmt"
// )

// func gpltPrint(gplt []int) {

// 	if gplt[0] > 0 {
// 		fmt.Printf("G")
// 		gplt[0] -= 1
// 	}
// 	if gplt[1] > 0 {
// 		fmt.Printf("P")
// 		gplt[1] -= 1
// 	}
// 	if gplt[2] > 0 {
// 		fmt.Printf("L")
// 		gplt[2] -= 1
// 	}
// 	if gplt[3] > 0 {
// 		fmt.Printf("T")
// 		gplt[3] -= 1
// 	}
// 	if gplt[0] == 0 && gplt[1] == 0 && gplt[2] == 0 && gplt[3] == 0 {
// 		return
// 	} else {
// 		gpltPrint(gplt)
// 	}
// }

// func main() {
// 	gplt := []int{0, 0, 0, 0}
// 	var input string

// 	fmt.Scan(&input)
// 	for i := 0; i < len(input); i++ {
// 		if input[i] == 'G' || input[i]-'G' == 32 {
// 			gplt[0] += 1
// 		}
// 		if input[i] == 'P' || input[i]-'P' == 32 {
// 			gplt[1] += 1
// 		}
// 		if input[i] == 'L' || input[i]-'L' == 32 {
// 			gplt[2] += 1
// 		}
// 		if input[i] == 'T' || input[i]-'T' == 32 {
// 			gplt[3] += 1
// 		}

// 	}
// 	gpltPrint(gplt)

// }
