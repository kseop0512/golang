package main

import (
	"fmt"
	"strings"
)

// func repeatMe(words ...string) {
// 	fmt.Println(words)
// }

func lenAndUpper(name string) (length int, uppercase string) {
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func main() {
	// repeatMe("nico", "lynn", "dal", "marl", "flynn")
	totalLength, up := lenAndUpper("kseop")
	fmt.Println(totalLength, up)

}
