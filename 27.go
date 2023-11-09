// package main

// import "fmt"

// func main() {
// 	var telstr string
// 	var tel, arr []int
// 	var j int
// 	fmt.Scanf("%s\n", &telstr)
// 	for i := 0; i < len(telstr); i++ {
// 		tel = append(tel, int(telstr[i]-'0'))
// 	}
// 	for j = 9; j >= 0; j-- {
// 		for i := 0; i < len(tel); i++ {
// 			if tel[i] == j {
// 				arr = append(arr, j)
// 				break
// 			}
// 		}
// 	}
// 	fmt.Printf("int[] arr = new int[]{")
// 	for i, v := range arr {
// 		if i == len(arr)-1 {
// 			fmt.Printf("%d", v)
// 		} else {
// 			fmt.Printf("%d,", v)
// 		}
// 	}
// 	fmt.Printf("};\n")

// 	fmt.Printf("int[] index = new int[]{")
// 	for i := 0; i < len(tel); i++ {
// 		for j := 0; j < len(arr); j++ {
// 			if tel[i] == arr[j] {
// 				if i == len(tel)-1 {
// 					fmt.Printf("%d", j)
// 				} else {
// 					fmt.Printf("%d,", j)
// 				}
// 			}
// 		}
// 	}
// 	fmt.Printf("};\n")

// }
