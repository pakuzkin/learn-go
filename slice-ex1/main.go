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
}
