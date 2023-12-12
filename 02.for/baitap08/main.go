package main

import "fmt"

func main() {
	fmt.Println("Nhập số nguyên n bất kỳ. Viết chương trình in các số lẻ từ 1 đến n.")
	var number int
	fmt.Print("Nhập giá trị cho biến: ")
	fmt.Scanf("%d", &number)
	fmt.Printf("Giá trị của biến là: %d\n", number)
	for i := 0; i <= number; i++ {
		if i%2 != 0 {
			fmt.Printf("Giá trị lẻ là %d\n", i)
		}
	}
}
