package main

import (
	"fmt"
	"strings"
)

// func repeatMe(words ...string) {
// 	fmt.Println(words)
// }

func lenAndUpper(name string) (length int, uppercase string) {
	//defer 함수 실행 후 동작
	defer fmt.Println("I'm done")
	length = len(name)
	uppercase = strings.ToUpper(name)
	return // 필수 X
}

func main() {
	// repeatMe("nico", "lynn", "dal", "marl", "flynn")
	totalLength, up := lenAndUpper("kseop")
	fmt.Println(totalLength, up)

}

//defer
