package main

import "fmt"

func main() {
	//variabel tanpa value
	var celcius int
	var suhu string
	var hasil float32

	//untuk deklarasi value celcius & suhu
	fmt.Print("masukkan suhu Celcius: ")
	fmt.Scan(&celcius)
	fmt.Print("convert ke mana? (huruf kecil semua) \n Jawab:")
	fmt.Scan(&suhu)

	//penggunaan switch case untuk suhu yang diinginkan
	switch suhu {
	case "fahrenheit":
		hasil = float32(celcius)*9/5 + 32
		fmt.Print("Fahrenheit: ", hasil)

	case "reamur":
		hasil = float32(celcius) * 4 / 5
		fmt.Print("Reamur: ", hasil)

	case "kelvin":
		hasil = float32(celcius) + 273
		fmt.Print("Kelvin: ", hasil)

	default:
		fmt.Print("Kamu salah ketik")
	}

}
