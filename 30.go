// package main

// import "fmt"

// func main() {
// 	var counter int
// 	var sex []int
// 	var girlsInd, boysInd []int
// 	var girlsName, boysName []string
// 	fmt.Scan(&counter)
// 	for i := 0; i < counter; i++ {
// 		var s int
// 		var n string
// 		fmt.Scan(&s, &n)
// 		sex = append(sex, s)
// 		if s == 0 {
// 			girlsName = append(girlsName, n)
// 			girlsInd = append(girlsInd, i)
// 		} else {
// 			boysName = append(boysName, n)
// 			boysInd = append(boysInd, i)
// 		}
// 	}

// 	girlsGrouped, boysGrouped := 0, 0
// 	for i := 0; i < counter>>1; i++ {

// 		if sex[i] == 0 {
// 			fmt.Printf("%s %s\n", girlsName[girlsGrouped], boysName[len(boysName)-1-girlsGrouped])
// 			girlsGrouped += 1

// 		} else {
// 			fmt.Printf("%s %s\n", boysName[boysGrouped], girlsName[len(boysName)-1-boysGrouped])
// 			boysGrouped += 1
// 		}

// 	}
// }
