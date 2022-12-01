package main

import "fmt"

func main() {

	fmt.Println("Pilihan Operasi")
	fmt.Println("===============")

	fmt.Println("1. Penjumlahan")
	fmt.Println("2. Pengurangan")
	fmt.Println("3. Pembagian")
	fmt.Println("4. Perkalian")

	var Operasi float32
	fmt.Println("Masukkan pilihan : ")
	fmt.Scanln(&Operasi)

	var a, b, c float32

	if Operasi == 1 {
		fmt.Println("Masukkan angka pertama : ")
		fmt.Scanln(&a)
		fmt.Println("Masukkan angka Kedua : ")
		fmt.Scanln(&b)
		c = a + b
		fmt.Println("Hasilnya adalah : ", c)
	} else if Operasi == 2 {
		fmt.Println("Masukkan angka pertama : ")
		fmt.Scanln(&a)
		fmt.Println("Masukkan angka Kedua : ")
		fmt.Scanln(&b)
		c = a - b
		fmt.Println("Hasilnya adalah : ", c)
	} else if Operasi == 3 {
		fmt.Println("Masukkan angka pertama : ")
		fmt.Scanln(&a)
		fmt.Println("Masukkan angka Kedua : ")
		fmt.Scanln(&b)
		if b == 0 {
			c = 0
		} else {
			c = a / b
		}
		fmt.Println("Hasilnya adalah : ", c)
	} else if Operasi == 4 {
		fmt.Println("Masukkan angka pertama : ")
		fmt.Scanln(&a)
		fmt.Println("Masukkan angka Kedua : ")
		fmt.Scanln(&b)
		c = a * b
		fmt.Println("Hasilnya adalah : ", c)
	}
	fmt.Println("=========================")

}

