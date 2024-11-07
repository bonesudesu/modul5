package main

import "fmt"

func faktor(n int, i int) []int {
	if i > n {
		return nil
	}
	if n%i == 0 {
		return append(faktor(n, i+1), i)
	}
	return faktor(n, i+1)
}

func main() {
	var n int
	fmt.Print("Masukkan bilangan bulat positif: ")
	fmt.Scanln(&n)

	faktorN := faktor(n, 1)
	fmt.Println("Faktor dari", n, ":", faktorN)
}
