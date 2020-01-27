package main

import "fmt"

var a int = 10
var b int = 20
var c float64 = 1.234
var d float64 = 5.678

func main() {
	//hello()
	//calc(a,b,c,d)
	fmt.Println(square(10))
	samplerange()
}
func hello() {
    fmt.Println("hello, world")
}
func calc(a int, b int, c float64, d float64) {
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(c * d)
	fmt.Println(c / d)
}
func square(x int) int {
	return x * x
}
func samplerange() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    sum1 := 0
    for i := 0; i < len(a); i++ {
        sum1 += a[i]
    }
    fmt.Println(sum1)
    sum2 := 0
    for _, v := range a {
		fmt.Println(v)
        sum2 += v
    }
    fmt.Println(sum2)
}
