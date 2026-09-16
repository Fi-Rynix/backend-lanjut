package model

type JadwalKuliah struct {
	IDKuliah    int    `json:"id_kuliah"`
	IDStudent   int    `json:"id_student"`
	MataKuliah  string `json:"mata_kuliah"`
	Hari        string `json:"hari"`
}

