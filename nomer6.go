package main

import "fmt"

func power(x, y int) int {
	if y == 0 {
		return 1
	}
	return x * power(x, y-1)
}

func main() {
	var x, y int
	fmt.Print("Masukkan bilangan x: ")
	fmt.Scanln(&x)
	fmt.Print("Masukkan bilangan y: ")
	fmt.Scanln(&y)
	fmt.Println("Hasil", x, "pangkat", y, "adalah", power(x, y))
}
