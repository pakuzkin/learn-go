package main

import (
	"fmt"

	mypackage "github.com/pakuzkin/learn-go/mymodule"
)

func main() {
	fmt.Println("Hello, Modules!")

	mypackage.PrintHello()
}
