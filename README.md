# learn-go
https://pkg.go.dev/fmt#pkg-overview

### Vars

`var intArray [5]int` // declare int arr of size 5

`aSlice := make([]string, 10)` // dynamic arr aka slice

### Loop

```
for i := 0; i < 5; i++ {
    // Loop body
}
```

```
for index, value := range collection {
    // Loop body
}
```

```
for condition {
    // While body
}
```

### Copy slice

```
func copyArr(x []float64) []float64 {
	myCopy := make([]float64, len(x))
	copy(myCopy, x)
	return myCopy
}
```

### Delete from slice

Delete 5th element from slice `s`
```
newSlice := append(s[:5], s[6:]...)
```

### Map

```
    myMap := make(map[string]int)
```

```
    myMap := map[string]int{
        "Todd":   42,
        "Henry":  16,
        "Padget": 14,
    }
```

### Structures

```
    type person struct {
        first string
        last  string
        age   int
    }
```


### Functions

```
func dogYears(name string, age int) (string, int) {
	age *= 7
	return fmt.Sprint(name, " is this old in dog years "), age
}
```

### Methods 

```
// Define a struct named 'Rectangle'
type Rectangle struct {
    Width  float64
    Height float64
}

// Define a method for the 'Rectangle' type
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}
```

### Pointers

`&` the address of a variable, `*` the value at an address.

Values that are not refed by memory go to the stack. Referenced values go to the heap


```
y := &x // y is a pointer to the memory address of x
fmt.Println("*y:", *y) // Dereferencing the pointer
```


### I/O

Read
```
file, err := os.Open(filename)
if err != nil {
    fmt.Println("Error opening file:", err)
    return
}
defer file.Close()
```

Write
```
func main() {
	f, err := os.Create("output.txt")
	if err != nil {
		log.Fatalf("error %s", err)
	}
	defer f.Close()

	s := []byte("Hello gophers!")

	_, err = f.Write(s)
	if err != nil {
		log.Fatalf("error %s", err)
	}
}
```