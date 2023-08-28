package main

import (
	"fmt"

	"github.com/GoesToEleven/puppy"
	mypackage "github.com/pakuzkin/learn-go/mymodule"
)

func main() {
	fmt.Println("Hello, Modules!")

	mypackage.PrintHello()
	fmt.Println(puppy.Barks())

}
