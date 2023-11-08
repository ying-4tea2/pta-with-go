// package main

// import (
// 	"fmt"
// )

// func contains(elems []string, v string) bool {
// 	for _, s := range elems {
// 		if v == s {
// 			return true
// 		}
// 	}
// 	return false
// }

// func main() {
// 	var num int

// 	var handsome []string
// 	var lookupNum int

// 	idMap := make(map[string]int)

// 	var id string

// 	fmt.Scan(&num)

// 	for i := 0; i < num; i++ {
// 		var cols int
// 		fmt.Scan(&cols)
// 		for j := 0; j < cols; j++ {
// 			fmt.Scan(&id)
// 			if cols > 1 {
// 				idMap[id] += 1
// 			}

// 		}
// 	}

// 	fmt.Scan(&lookupNum)

// 	for i := 0; i < lookupNum; i++ {
// 		fmt.Scan(&id)
// 		_, ok := idMap[id]
// 		added := contains(handsome, id)

// 		if !ok {
// 			if !added {
// 				handsome = append(handsome, id)
// 			}
// 		}

// 	}

// 	if len(handsome) == 0 {
// 		fmt.Println("No one is handsome")
// 	} else {

// 		for _, v := range handsome {
// 			fmt.Printf("%v ", v)
// 		}
// 	}

// }
