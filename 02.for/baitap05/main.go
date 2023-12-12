package main

import "fmt"

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

	//Go cung cấp một cách tốt hơn và ngắn gọn hơn
	//để lặp qua một mảng bằng cách sử dụng range của vòng lặp for. range trả về mục và cả giá trị tại mục đó.
	//Cùng viết lại đoạn code trên bằng cách sử dụng range. Chúng ta cũng sẽ tìm thấy tổng của tất cả các phần tử của mảng.
	sumA := 0
	for _, v := range a { //range returns both the index and value

		sumA += v
	}
	fmt.Println("Tổng của mảng A = ", sumA)

	sumB := 0
	for _, v := range b { //range returns both the index and value
		sumB += v
	}
	fmt.Println("Tổng của mảng B = ", sumB)

	sumC := 0
	for _, v := range c { //range returns both the index and value
		sumC += v
	}
	fmt.Println("Tổng của mảng C = ", sumC)

}
