package main

import (
	"fmt"
)

func main() {

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
