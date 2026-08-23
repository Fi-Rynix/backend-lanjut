package main

import "fmt"

func main() {
  // deklarasi variabel
	var nama string = "Ihya"
	var angka int = 67
	var panjang float64 = 67.67
	var status bool = true
	tinggi := 167.67
	ngawi := []string{"Amba", "Rusdi", "Vito", "Gatot"}

	fmt.Println("")
	fmt.Println("Nama:", nama)
	fmt.Println("Angka favorit:", angka)
	fmt.Println("Panjang:", panjang)
	fmt.Println("Tinggi:", tinggi)
	fmt.Println("Aku mahasiswa?:", status)
	fmt.Println("Tetangga Ngawi:", ngawi)
	
	// map
	mahasiswa := map[string]string{
		"Ihya": "Raja Iblis",
		"Udin": "Raja Karbit",
		"Vito": "Raja Suki",
	}
	fmt.Println("")
	fmt.Println("Isi full map:", mahasiswa)
	fmt.Println("")

	// tambah
	mahasiswa["Rizki"] = "Raja Jomok"
	
	// cek
	value, rilkah := mahasiswa["Rizki"]
	if rilkah {
		fmt.Println("Julukan:", value)
	} else {
		fmt.Println("Data tidak ditemukan")
	}

	value2 := mahasiswa["Vito"]
	fmt.Println("Julukan:", value2)

	value3 := mahasiswa["Fedo"]
	fmt.Println("Julukan:", value3)

	// hapus
	delete(mahasiswa, "Ihya")

	// looping
	fmt.Println("\nIsi map yang baru:")
	for nama, julukan := range mahasiswa {
		fmt.Println(nama, ":", julukan)
	}


}

func swapPointer(a, b *int) {
	temp := *a
	*a = *b
	*b = temp

	fmt.Println("nilai didalam fungsi swapPointer:")
	fmt.Println("a =", *a)
	fmt.Println("b =", *b)
}

func swapBiasa(a, b int) {
	temp := a
	a = b
	b = temp

	fmt.Println("nilai didalam fungsi swapBiasa:")
	fmt.Println("a =", a)
	fmt.Println("b =", b)
}