// package main

// import (
// 	"fmt"
// 	"sort"
// )

// func main() {
// 	numberMap := make(map[string]int)
// 	var number string
// 	fmt.Scanf("%s", &number)

// 	for i := 0; i < len(number); i++ {
// 		digit := string(number[i])
// 		_, ok := numberMap[digit]
// 		if ok {
// 			numberMap[digit] = numberMap[digit] + 1

// 		} else {
// 			numberMap[digit] = 1
// 		}
// 	}

// 	var keys []string
// 	for k, _ := range numberMap {
// 		keys = append(keys, k)
// 	}

// 	sort.Strings(keys)
// 	fmt.Printf("keys: %v\n", keys)

// 	for _, k := range keys {
// 		fmt.Printf("%s:%d\n", k, numberMap[k])
// 	}

// }
