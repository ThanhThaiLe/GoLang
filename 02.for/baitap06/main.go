package main

import "fmt"

func main() {
	fmt.Println("Viết chương trình in bảng cửu chương rút gọn của các số từ 1 đến 100.")
	var number int

	fmt.Print("Nhập giá trị cho biến: ")
	fmt.Scanf("%d", &number)

	fmt.Printf("Giá trị của biến là: %d\n", number)
	for i := 1; i <= number; i++ {
		fmt.Printf("Bảng cửu chương %d\n", i)
		for j := 0; j <= 10; j++ {
			var value = i * j
			fmt.Printf("%d * %d = %d\n", i, j, value)
		}
		fmt.Printf("\n")
	}
}
