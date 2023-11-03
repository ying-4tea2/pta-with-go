// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"math"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// // 公约数 x<y
// func GCD(x int, y int) int {

// 	if x > y {
// 		return GCD(y, x)
// 	} else if x == y {
// 		return x
// 	} else {
// 		if (y % x) != 0 {
// 			return GCD(x, y%x)
// 		} else {
// 			return x
// 		}
// 	}

// }

// func LCM(x int, y int) int {
// 	gcd := GCD(x, y)
// 	lcm := x * y / gcd
// 	return lcm
// }

// func main() {

// 	var count int
// 	fmt.Scanf("%d\n", &count)
// 	var numerator, denominator []int

// 	inputReader := bufio.NewReader(os.Stdin)
// 	line, _ := inputReader.ReadString('\n')
// 	numbers := strings.Fields(line)
// 	for _, v := range numbers {
// 		tmp := strings.Split(v, "/")
// 		ntmp, _ := strconv.Atoi(tmp[0])
// 		dtmp, _ := strconv.Atoi(tmp[1])
// 		numerator = append(numerator, ntmp)
// 		denominator = append(denominator, dtmp)
// 	}

// 	lcm := denominator[0]
// 	for i := 0; i < count-1; i++ {

// 		lcm = LCM(lcm, denominator[i+1])
// 	}
// 	//	fmt.Printf("lcm: %v\n", lcm)

// 	for i := 0; i < count; i++ {
// 		numerator[i] = numerator[i] * lcm / denominator[i]
// 	}
// 	//	fmt.Printf("numerator: %v\n", numerator)

// 	var sum_numerator, intPart int

// 	for i := 0; i < count; i++ {
// 		sum_numerator += numerator[i]
// 	}

// 	if sum_numerator == 0 {
// 		fmt.Printf("0\n")
// 		return
// 	}
// 	//fmt.Printf("sum_numerator: %v\n", sum_numerator)
// 	// 如果有整数部分
// 	if int(math.Abs(float64(sum_numerator))) >= lcm {
// 		intPart = int(sum_numerator / lcm)
// 		sum_numerator = sum_numerator % lcm
// 	}

// 	if intPart == 0 {

// 		fmt.Printf("%v/%v", sum_numerator, lcm)

// 	} else {
// 		if sum_numerator == 0 {
// 			fmt.Printf("%v", intPart)
// 		} else {
// 			gcd := GCD(sum_numerator, lcm)
// 			gcd = int(math.Abs(float64(gcd)))
// 			//fmt.Printf("sum_numerator: %v\n", sum_numerator)
// 			//fmt.Printf("lcm: %v\n", lcm)
// 			//fmt.Printf("gcd: %v\n", gcd)
// 			sum_numerator = sum_numerator / gcd
// 			lcm = lcm / gcd
// 			fmt.Printf("%v %v/%v", intPart, sum_numerator, lcm)
// 		}

// 	}

// }
