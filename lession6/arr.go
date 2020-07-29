package main

import "fmt"

func main() {
	//arrInit()
	//rangeArr()
	//manyDimensionsArr()
	//modifyArr()
	//arrPoint()
	arrSum := sumArr([...]int{1, 3, 5, 7, 8})
	fmt.Println(arrSum)
}

//test arr init
func arrInit() {
	a1 := [...]int{}
	a2 := [...]string{"a", "bbb", "ccc"}
	a3 := [...]int{1, 3: 5}
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
}

//range arr
func rangeArr() {
	b1 := [...]string{"jim", "jack", "tom"}
	for k, v := range b1 {
		fmt.Println(k, v)
	}
}

//many dimensions arr,just the first can use ...
func manyDimensionsArr() {
	c1 := [...][2]string{
		{"beijing", "shanghai"},
		{"guangzhou", "shenzhen"},
		{"guangzhou", "shenzhen"},
	}
	for _, v1 := range c1 {
		fmt.Println(v1)
	}
}

//array is a quote type,so you cann't change the value in array
func modifyArr() {
	d1 := [...]int{1, 2, 3}
	modifyArray(d1)
	fmt.Println(d1)
}

func modifyArray(x [3]int) {
	x[0] = 100
}

// compare [...]*int{...} and *[...]int{...}
func arrPoint() {
	e1 := 3
	e2 := 4
	e3 := 5
	e4 := [3]*int{&e1, &e2, &e3}
	fmt.Println(e4)
	f1 := [2]int{1, 2}
	f2 := &f1
	fmt.Printf("%T\n", f1)
	fmt.Printf("%T\n", *f2)
}

//求数组[1, 3, 5, 7, 8]所有元素的和
func sumArr(x [5]int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

//找出数组中和为指定值的两个元素的下标，比如从数组[1, 3, 5, 7, 8]中找出和为8的两个元素的下标分别为(0,3)和(1,2)
func findElement(x [5]int, sum int) (a int, b int) {
	return a, b
}
