package main

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

/*
// golang if else
func canIDrink(age int) bool {
	if kreanAge := age + 1; kreanAge < 18 {
		return false
	}

	return true
}

func main() {
	fmt.Println(canIDrink(16))
}
*/

// Swtich

func canIDrink(age int) bool {
	/*
		switch age {
		case 10:
			return false
		case 18:
			return true
		}
	*/
	/*
		switch {
		case age < 18:
			return false
		case age == 18:
			return true
		case age > 50:
			return false
		}
	*/

	switch koreanAge := age + 2; koreanAge {
	case 10:
		return false
	case 18:
		return true
	}
	return false
}

/*
func main() {
	fmt.Println(canIDrink(18))
}


*/

/*
// Pointers
func main() {
	a := 2
	b := &a
	*b = 202020
	fmt.Println(a)
}

*/

/*
// Arrays and Slices
func main() {
	names := []string{"nico", "lynn", "dal"}
	names = append(names, "kseop")
	fmt.Println(names)
}
*/

// Maps

/*
// key : value
func main() {
	nico := map[string]string{"name": "kseop", "age": "12"}
	for key, _ := range nico {
		fmt.Println(key)
	}
}
*/

/*
// Structs
type person struct {
	name    string
	age     int
	favFood []string
}

func main() {
	favFood := []string{"kimchi", "raman"}
	seop := person{name: "kseop", age: 22, favFood: favFood}
	fmt.Println(seop.name)
}
*/

// Account + NewAccount
