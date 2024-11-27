package main

import "fmt"

func main() {
	var x, y, i int
	var a bool
	fmt.Print("bilangan: ")
	fmt.Scan(&x)
	fmt.Print("faktor: ")
	for i = 1; i <= x; i++ {
		if x%i == 0 {
			y++
			fmt.Print(i, " ")
		}
	}
	fmt.Print("\nprima: ")
	if y <= 2 {
		a = true
		fmt.Print(a)
	} else {
		fmt.Print(a)
	}
}
