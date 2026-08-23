package main

import "fmt"

	type Student struct {
	ID       int
	Name     string
	Grade    float64
	IsActive bool
	}

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

	fmt.Println("")

	// swap
	p := 10
	q := 20
	r := 30
	s := 40

	swapPointer(&p, &q)
	swapBiasa(r, s)

	println("\nnilai pqrs setelah kita run 2 fungsi tadi:")
	println("p =", p)
	println("q =", q)
	println("r =", r)
	println("s =", s)

	// update slice
	sukiLiar := []string{
        "Rizkimok",
        "Udin Karbit",
        "Vitomok",
    }

	fmt.Println("")
	updateSlice(&sukiLiar, "Reza Kecap")
	fmt.Println("After update:", sukiLiar)


	// struct
	mahasiswa1 := Student{
	ID:       1,
	Name:     "Ihya",
	Grade:    3.99,
	IsActive: false,
	}

	fmt.Println("")
	fmt.Println(mahasiswa1.GetInfo())

	fmt.Println("")
	mahasiswa1.UpdateGrade(4.00)
	fmt.Println("After:")
	fmt.Println(mahasiswa1.GetInfo())

	fmt.Println("")

	mahasiswa1.Activate()
	fmt.Println("Setelah after:")
	fmt.Println(mahasiswa1.GetInfo())

	mahasiswa1.Deactivate()
	fmt.Println("Setelah after after:")
	fmt.Println(mahasiswa1.GetInfo())

	fmt.Println("")
	
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

func updateSlice(sp *[]string, newNama string) {
    *sp = append(*sp, newNama)
}

func (sr Student) GetInfo() string {
	return fmt.Sprintf(
		"ID: %d, Name: %s, Grade: %.2f, Active: %t", sr.ID, sr.Name, sr.Grade, sr.IsActive,
	)
}

func (sr *Student) UpdateGrade(gradeBaru float64) {
    sr.Grade = gradeBaru
}

func (sr *Student) Activate() {
    sr.IsActive = true
}

func (sr *Student) Deactivate() {
    sr.IsActive = false
}