package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(arr, "len", len(arr), "type")
	for i, v := range arr {
		fmt.Println(i, v)
	}

	deleted := append(arr[:3], arr[4:]...)
	fmt.Println(deleted)

	fmt.Println("arr - len", len(arr), "cap", cap(arr))
	fmt.Println("deleted - len", len(deleted), "cap", cap(deleted))

	added := append(deleted, 20, 21, 22, 23)
	fmt.Println(added)

	fmt.Println("added - len", len(added), "cap", cap(added))

	// ex 42
	ex42 := make([]int, 10)
	for i := 0; i < 10; i++ {
		ex42[i] = i + 42
	}
	fmt.Println(ex42)
	for i, v := range ex42 {
		fmt.Println(i, v)
	}

	// ex 43
	fmt.Println(ex42[:5])
	fmt.Println(ex42[5:])
	fmt.Println(ex42[2:7])
	fmt.Println(ex42[1:6])

	// ex 44
	ex44 := append(ex42, 53, 54, 55)
	fmt.Println(ex42)
	fmt.Println(ex44)

	// ex 45
	ex45 := append(ex44[:5], ex44[7:]...)
	fmt.Println(ex45)

}
