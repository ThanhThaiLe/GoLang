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

	//Viết chương trình in các số từ 1 đến 1000 theo thứ tự tăng dần.
	fmt.Println("Viết chương trình in các số từ 1 đến 1000 theo thứ tự tăng dần.")
	for i := 1; i <= 1000; i++ {
		fmt.Printf("%d\n", i)
	}

	// Viết chương trình in các số từ 1 đến 1000 theo thứ tự giảm dần.
	fmt.Println("Viết chương trình in các số từ 1 đến 1000 theo thứ tự giảm dần.")
	for i := 1000; i >= 1; i-- {
		fmt.Printf("%d\n", i)
	}

	//Viết chương trình in bảng số từ 1 đến 200.
	fmt.Println("Viết chương trình in bảng số từ 1 đến 200.")
	for i := 0; i < 200; i++ {
		fmt.Printf("%d\n", i+1)
	}
	//Viết chương trình nhập một số nguyên, tìm bội số của số đó với các số từ 1 đến 20, sau đó in kết quả ra màn hình.
	fmt.Println("Viết chương trình nhập một số nguyên, tìm bội số của số đó với các số từ 1 đến 20, sau đó in kết quả ra màn hình. ")

	var number int

	fmt.Print("Nhập giá trị cho biến: ")
	fmt.Scanf("%d", &number)

	fmt.Printf("Giá trị của biến là: %d\n", number)
	for i := 0; i < 20; i++ {
		var value = number * (i + 1)
		fmt.Printf("Vị trí là: %d\n", (i + 1))
		fmt.Printf("Giá trị bội của n là: %d\n", value)
	}
}
