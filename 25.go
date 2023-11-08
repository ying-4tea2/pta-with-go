package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var result string
	var numStrs []string

	r := bufio.NewReader(os.Stdin)
	strin, _ := r.ReadString('\n')
	//the string returns by ReadString method includs the delim thus the last element should be deleted
	strin = strin[:len(strin)-1]
	numStrs = strings.Split(strin, " ")
	num1, err1 := strconv.ParseUint(numStrs[0], 10, 64)
	if num1 == 0 || num1 > 1000 {
		err1 = errors.New("Out of range")
	}
	num2, err2 := strconv.ParseUint(numStrs[1], 10, 64)
	if num2 == 0 || num2 > 1000 {
		err2 = errors.New("Out of range")
	}
	if err1 == nil && err2 == nil {
		result = strconv.Itoa(int(num1 + num2))
		fmt.Printf("%v + %v", num1, num2)
	} else {
		if err1 != nil {
			fmt.Printf("?")
			result = "?"
		} else {
			fmt.Printf("%v", num1)
		}
		fmt.Printf(" + ")
		if err2 != nil {
			fmt.Printf("?")
			result = "?"
		} else {
			fmt.Printf("%v", num2)
		}
	}

	fmt.Printf(" = %s", result)

}
