// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strings"
// )

// func main() {
// 	var studentsCount int
// 	var testNum string
// 	var preTestSeatNum, testSeatNum string
// 	preTest2testNum := make(map[string]string)
// 	preTest2testSeatNum := make(map[string]string)

// 	fmt.Scanf("%d\n", &studentsCount)

// 	// inputReader := bufio.NewReader(os.Stdin)
// 	// line, _ := inputReader.ReadString('\n')
// 	// arr := strings.Fields(line)
// 	// fmt.Printf("arr: %v\n", arr)

// 	for i := 0; i < studentsCount; i++ {
// 		//fmt.Printf("Input the %vth value\n", i)
// 		fmt.Scanf("%s %s %s\n", &testNum, &preTestSeatNum, &testSeatNum)
// 		preTest2testNum[preTestSeatNum] = testNum
// 		preTest2testSeatNum[preTestSeatNum] = testSeatNum
// 	}

// 	var lookupCount int
// 	fmt.Scanf("%d\n", &lookupCount)
// 	inputReader := bufio.NewReader(os.Stdin)
// 	line, _ := inputReader.ReadString('\n')
// 	arr := strings.Fields(line)

// 	for _, v := range arr {
// 		fmt.Printf("%s %s\n",preTest2testNum[v], preTest2testSeatNum[v])
// 	}
// }
