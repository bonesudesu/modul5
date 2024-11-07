package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masukkan bilangan bulat positif N: ")
	fmt.Scanln(&n)

	fmt.Println("Barisan bilangan ganjil:", printGanjil(n))
}

func printGanjil(n int) string {
	if n <= 0 {
		return ""
	} else {
		return fmt.Sprintf("%d %s", n, printGanjil(n-2))
	}
}
