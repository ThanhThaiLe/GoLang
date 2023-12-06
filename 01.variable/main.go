package main

import (
	"fmt"
	"reflect"
)

// biến toàn cục
var d int = 4

func main() {
	//cú pháp
	//var + variable name + datatype + value
	var name string = "Lê Thanh Thái"
	fmt.Println(name)

	var first_name, last_name string = "Lê", " Thanh Thái"
	fmt.Println(first_name + last_name)

	var a int = 5
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)
	fmt.Println(b + c)

	// biến hiểu ngầm với giá trị khởi tạo

	var number = 5
	fmt.Println(number)
	fmt.Printf("%T: %d\n", number, number)
	fmt.Println("Demo")
	x := 42
	y := float32(43.3)
	z := "hello"

	xt := reflect.TypeOf(x).Kind()
	yt := reflect.TypeOf(y).Kind()
	zt := reflect.TypeOf(z).Kind()

	fmt.Printf("%T: %s\n", xt, xt)
	fmt.Printf("%T: %s\n", yt, yt)
	fmt.Printf("%T: %s\n", zt, zt)

	if xt == reflect.Int {
		println(">> x is int")
	}
	if yt == reflect.Float32 {
		println(">> y is float32")
	}
	if zt == reflect.String {
		println(">> z is string")
	}

	//biến cục bộ
	fmt.Printf("d=%d\n", d)
	//nếu trong 1 hàm có biến cục bộ và biến toàn cục trùng tên thì hệ thống sẽ ưu tiên lấy biến cục bộ để sử dụng
}
