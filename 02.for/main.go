package main

import (
	"fmt"
)

func main() {

	//cú pháp
	//var name [max-length] datatype
	//var name [max-length] datatype{value1, value2,...,valueN}
	//name:= [max-length] datatype
	//name:= [max-length] datatype{value1, value2,...,valueN}
	var a [3]int //int array with length 3
	fmt.Println(a)

	var b [3]int //int array with length 3
	b[0] = 12    // array index starts at 0
	b[1] = 78
	b[2] = 50
	fmt.Println(b)

	c := [3]int{100, 300, 50} // short hand declaration to create array
	fmt.Println(c)

	// length of array

	var lengthA int = len(a)
	fmt.Println("length array a:", lengthA)
	var lengthB int = len(b)
	fmt.Println("length array b:", lengthB)
	var lengthC int = len(c)
	fmt.Println("length array c:", lengthC)

	//in từng phần tử trong mảng
	fmt.Println("Các giá trị của mảng a")
	for i := 0; i < len(a); i++ { //looping from 0 to the length of the array
		fmt.Printf("Giá trị vị trí %d của mảng a là %d\n", i+1, a[i])
	}
	fmt.Println("Các giá trị của mảng b")
	for i := 0; i < len(b); i++ { //looping from 0 to the length of the array
		fmt.Printf("Giá trị vị trí %d của mảng b là %d\n", i+1, b[i])
	}
	fmt.Println("Các giá trị của mảng c")
	for i := 0; i < len(c); i++ { //looping from 0 to the length of the array
		fmt.Printf("Giá trị vị trí %d của mảng c là %d\n", i+1, c[i])
	}
	fmt.Println("Các giá trị của mảng a use for range")
	for i, v := range a { //range returns both the index and value
		fmt.Printf("Giá trị vị trí %d của mảng a là %d\n", i+1, v)
	}
	fmt.Println("Các giá trị của mảng b use for range")
	for i, v := range b { //range returns both the index and value
		fmt.Printf("Giá trị vị trí %d của mảng b là %d\n", i+1, v)
	}
	fmt.Println("Các giá trị của mảng c use for range")
	for i, v := range c { //range returns both the index and value
		fmt.Printf("Giá trị vị trí %d của mảng c là %d\n", i+1, v)
	}
}
