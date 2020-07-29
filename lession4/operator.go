package main

import "fmt"

func main() {
	// +
	// a1 := 1
	// a2 := 2
	// a3 := a1 + a2
	// fmt.Println(a3)
	// // -
	// b1 := 10
	// b2 := 2
	// b3 := b1 - b2
	// fmt.Println(b3)
	// // *
	// c1 := 20
	// c2 := 3.2
	// c3 := float64(c1) * c2
	// fmt.Println(c3)
	// // /
	// d1 := 65
	// d2 := 5.3
	// d3 := float64(d1) / d2
	// fmt.Printf("%.2f\n", d3)
	// // %
	// e1 := 555
	// e2 := 33
	// e3 := e1 % e2
	// fmt.Println(e3)
	// // ==
	// f1 := 0
	// f2 := 0.0
	// if float64(f1) == f2 {
	// 	fmt.Println("aaa")
	// } else {
	// 	fmt.Println("bbb")
	// }
	// if there has a most figure,and they most has appear twice,please find the number has appear once
	g1 := []int{1, 2, 3, 4, 5, 6, 6, 5, 4, 3, 1}
	newMap := make(map[int]int)
	for _, val := range g1 {
		if newMap[val] != 0 {
			newMap[val]++
		} else {
			newMap[val] = 1
		}
	}
	for k, v := range newMap {
		if v == 1 {
			fmt.Println(k)
		}
	}
}
