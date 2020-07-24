package main

import "fmt"

var (
	name  string  = "aaaa"
	age   int8    = 10
	money float32 = 3000.33
)

const (
	pi = 3.141592653
	n1 = iota
	n2
	n3
	n4
)

func main() {
	a := 3
	b := 3.14
	c := true
	d := "hello golang"
	e := &a
	f := [3]int{1, 2, 4}
	g := []int{}
	h := map[string]int{}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(money)
	fmt.Println(pi)
	i, _ := testAnoy(3, 4)
	fmt.Println(i)
	fmt.Println(n1, n2, n3, n4)
}

//this is a test about anonymous variable
func testAnoy(x int, y int) (int, int) {
	return y, x
}
