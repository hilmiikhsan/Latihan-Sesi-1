package main

import (
	"fmt"
)

func setMessage(car map[string]string) string {
	name := car["name"]
	color := car["color"]
	return "Mobil " + name + " berwarna " + color
}

func getMessage(message string) {
	fmt.Println(message)
}

func main() {
	car := make(map[string]string)
	car["name"] = "BWM"
	car["color"] = "Black"

	// buat 2 buah fungsi :
	// 1 => fungsi yang mengembalikan sebuah string
	// pada fungsi ini terjadi pengolahan kata sehingga menghasilkan kata : Mobil BMW berwarna Black

	// 2 => fungsi yang menampilkan hasil dari kembalian string
	// fungsi ini hanya bertugas untuk menampilkan kata

	// alur
	// simpan hasil dari return function kedalam sebuah variable message
	message := setMessage(car)

	// tampilkan hasil dari variable message
	getMessage(message)

	// Output: Mobil BWM berwarna Black
}
