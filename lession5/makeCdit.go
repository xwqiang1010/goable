package main

import "fmt"

func main() {
	// // if else
	// // a1 := 29
	// // if a1 < 18 {
	// // 	fmt.Println("this is not a adult age!")
	// // } else if a1 == 18 {
	// // 	fmt.Println("the age equal 18.")
	// // } else {
	// // 	fmt.Println("this is a adult age")
	// // 	//
	// // }
	// // // special if
	// // if a2 := "aaasss"; len(a2) >= 3 {
	// // 	fmt.Println("the string has more than three chars")
	// // } else {
	// // 	fmt.Println("the lenth is short or equal than three")
	// // }
	// // // for
	// // for i := 0; i < 10; i++ {
	// // 	fmt.Println(i)
	// // }
	// // for range
	// arr1 := []int{1, 2, 3, 4}
	// for k, v := range arr1 {
	// 	fmt.Println(k, v)
	// }

	// for _, thisV := range arr1 {
	// 	fmt.Println(thisV)
	// }
	// // switch case
	// switch finger := 2; finger {
	// case 1:
	// 	fmt.Println("this is first finger")
	// case 2:
	// 	fmt.Println("this is second finger")
	// 	fallthrough //this can do next case
	// case 3:
	// 	fmt.Println("this is third finger")
	// case 4, 5:
	// 	fmt.Println("this is other firger")
	// default:
	// 	fmt.Println("this is not a finger")
	// }

	// //goto
	// testGoto()
	// //breakMark
	// testBreak()
	// //continue
	// continueDemo()
	//print 9*9
	printMulti()
}

//this is a test about goto
func testGoto() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 5; j++ {
			if j == 3 {
				goto breakTag
			}
			fmt.Printf("%v--%v\n", i, j)
		}
	}
breakTag:
	fmt.Println("finished")
}

//this is a test about break
func testBreak() {
markB:
	for i := 0; i < 10; i++ {
		if i == 5 {
			break markB
		}
		fmt.Printf("num:%v\n", i)
	}
}

func continueDemo() {
forloop1:
	for i := 0; i < 5; i++ {
		// forloop2:
		for j := 0; j < 5; j++ {
			if i == 2 && j == 2 {
				continue forloop1
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}
}

func printMulti() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			if i == j {
				fmt.Printf("%d*%d=%d\n", j, i, i*j)
			} else {
				fmt.Printf("%d*%d=%d ", j, i, i*j)
			}
		}
	}
}
