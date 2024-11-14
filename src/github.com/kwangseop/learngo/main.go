package main

import "fmt"

// func repeatMe(words ...string) {
// 	fmt.Println(words)
// }

//defer
/*
func lenAndUpper(name string) (length int, uppercase string) {
	//defer 함수 실행 후 동작
	defer fmt.Println("I'm done")
	length = len(name)
	uppercase = strings.ToUpper(name)
	return // 필수 X
}
*/

// loop (for) ex) for of, for in

/*
golang for,range, ...args
func superAdd(numbers ...int) int {
	loop with range
	for index, number := range numbers {
		fmt.Println(index, number)
	}

	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}

	total := 0
	for _, number := range numbers {
		total += number
	}

	return total
}

func main() {
	result := superAdd(1, 2, 3, 4, 5, 6)
	fmt.Println(result)
}
*/

func canIDrink(age int) bool {
	if kreanAge := age + 1; kreanAge < 18 {
		return false
	}

	return true
}

func main() {
	fmt.Println(canIDrink(16))
}
