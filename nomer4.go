package main

import "fmt"

func printSequence(n int) {
	if n == 1 {
		fmt.Print(n, " ")
		return
	}
	fmt.Print(n, " ")
	printSequence(n - 1)
	fmt.Print(n, " ")
}

func main() {
	var n int
	fmt.Print("Masukkan bilangan bulat positif N: ")
	fmt.Scanln(&n)

	printSequence(n)
	fmt.Println()
}
