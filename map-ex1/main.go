package main

import "fmt"

func main() {
	myMap := make(map[string][]string)
	myMap["James Bond"] = []string{"shaken", "not", "stirred"}
	myMap["Dr No"] = []string{"cats", "ice cream", "sunset"}

	fmt.Println(myMap)
	for k, v := range myMap {
		fmt.Println(k, v)
	}
}
