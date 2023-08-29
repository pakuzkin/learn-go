# learn-go
https://pkg.go.dev/fmt#pkg-overview

### Vars

var intArray [5]int // declare int arr of size 5

aSlice := make([]string, 10) // dynamic arr aka slice

### Loop

for i := 0; i < 5; i++ {
    // Loop body
}

for index, value := range collection {
    // Loop body
}

for condition {
    // While body
}

### Copy array

```
func copyArr(x []float64) []float64 {
	myCopy := make([]float64, len(x))
	copy(myCopy, x)
	return myCopy
}
```


