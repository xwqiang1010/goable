package main

import "fmt"

func main() {
	//this is a test about data type
		var aa uint8
		v := 0b00101101
		aa = 33
		fmt.Printf("%v\n", aa)
		fmt.Printf("type:%T,%v\n", v, v)
		f := 3333.2363
		fmt.Printf("%.2f\n", f)
		//the follow is test about complex64 and complex128
		var c1 complex64
		var c2 complex64
		c1 = 3 + 6i
		c2 = 224 + 2i
		fmt.Println(c1 + c2)
		//this is a test about bool
		var b bool
		d1 := aa
		fmt.Println(b)
		fmt.Println(d1)
		//make a test about string
		s1 := "a"
		fmt.Println(s1)
		// make a long string
		s2 := `fsdfasjfd
	fsadffd
	dfsfddf
		`
		fmt.Println(s2)

	//test byte and rune byte include one char while rune include three char in utf-8
	s3 := "abc你好"
	fmt.Println(len(s3))
	for i := 0; i < len(s3); i++ {
		fmt.Printf("%c\n", s3[i])
	}

	for _, r := range s3 {
		fmt.Printf("%c\n", r)
	}

	//modify a string
	s4 := "abcdefg徐文强"
	byteS1 := []byte(s4)
	byteS2 := []rune(s4)
	byteS1[8] = 'x'
	fmt.Println(string(byteS1))
	byteS2[8] = '恒'
	fmt.Println(string(byteS2))

}
