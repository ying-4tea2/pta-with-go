// package main

// import (
// 	"fmt"
// )

// func isPrime(num int) {

// 	if num == 1 {
// 		fmt.Println("No")
// 	} else {
// 		for i := 2; i*i <= num; i++ {
// 			if num%int(i) == 0 {
// 				fmt.Println("No")
// 				return
// 			}
// 		}
// 		fmt.Println("Yes")
// 	}
// }

// func main() {
// 	var counter int
// 	var num []int
// 	fmt.Scan(&counter)
// 	for i := 0; i < counter; i++ {
// 		var tmp int
// 		fmt.Scan(&tmp)
// 		num = append(num, tmp)
// 	}

// 	for _, v := range num {
// 		isPrime(v)
// 	}
// }
