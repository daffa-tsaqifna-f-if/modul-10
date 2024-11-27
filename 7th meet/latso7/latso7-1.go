package main

import "fmt"

func main() {
	var x, y, z, b1, b2, bt int
	fmt.Print("Berat parcel (gram): ")
	fmt.Scan(&x)
	y = x / 1000
	z = x % 1000
	b1 = y * 10000
	fmt.Println("Detai berat:  ", y, "Kg + ", z, " gr")
	if z >= 500 {
		b2 = z * 5
	} else {
		b2 = z * 15
	}
	fmt.Println("Detail biaya: Rp.", b1, " + Rp. ", b2)
	if y >= 10 {
		bt = b1
		fmt.Println("Total biaya: Rp. ", bt)
	} else {
		bt = b1 + b2
		fmt.Println("Total biaya: Rp. ", bt)
	}
}
