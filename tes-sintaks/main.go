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

	// tambah
	mahasiswa["Rizki"] = "Raja Jomok"
	
	// cek
	value, rilkah := mahasiswa["Rizki"]
	if rilkah {
		fmt.Println("Julukan:", value)
	} else {
		fmt.Println("Data tidak ditemukan")
	}

	// hapus
	delete(mahasiswa, "Ihya")

	// looping
	for nama, julukan := range mahasiswa {
		fmt.Println(nama, ":", julukan)
	}


}