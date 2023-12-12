package main

import "fmt"

func main() {
	fmt.Println("Viết chương trình nhập một câu bất kỳ, đếm số từ và ký tự trong câu đó, và in kết quả ra màn hình.")
	var text string
	fmt.Print("Nhập giá trị cho biến text: ")
	fmt.Scanf("%s", &text)
	fmt.Printf("Giá trị của biến là: %s\n", text)

}
