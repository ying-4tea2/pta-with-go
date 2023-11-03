package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var str, cutStr string
	inputReader := bufio.NewReader(os.Stdin)
	str, _ = inputReader.ReadString('\n')
	cutStr, _ = inputReader.ReadString('\n')

	for _, v := range cutStr {
		str = strings.Replace(str, string(v), "", -1)
	}
	fmt.Printf("%s", str)
}
