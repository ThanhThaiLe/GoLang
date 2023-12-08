package main

import "fmt"

func main() {

	// Viết chương trình in các số từ 1 đến 1000 theo thứ tự giảm dần.
	fmt.Println("Viết chương trình in các số từ 1 đến 1000 theo thứ tự giảm dần.")
	for i := 1000; i >= 1; i-- {
		fmt.Printf("%d\n", i)
	}

}
