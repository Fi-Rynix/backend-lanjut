package service

import (
	"testing"

	"api-students/app/model"
)

func TestCountTotalPages(t *testing.T) {
	cases := []struct {
		total, limit, want int
	}{
		{0, 10, 0},
		{1, 10, 1},
		{10, 10, 1},
		{11, 10, 2},
		{137, 20, 7},
	}

	for _, tc := range cases {
		got := CountTotalPages(tc.total, tc.limit)
		if got != tc.want {
			t.Errorf("total=%d limit=%d: harap %d, dapat %d",
				tc.total, tc.limit, tc.want, got)
		}
	}
}

func TestApplyPatch(t *testing.T) {
	initial := model.Student{
		ID:       1,
		NIM:      434241001,
		Name:     "udin",
		Grade:    3.5,
		IsActive: true,
	}
	newName := "udin_updated"
	newGrade := 4.0

	result, errs := ApplyPatch(initial, model.PatchStudentRequest{
		Name:  &newName,
		Grade: &newGrade,
	})

	if len(errs) != 0 {
		t.Fatalf("tidak seharusnya ada error: %v", errs)
	}
	if result.Name != "udin_updated" {
		t.Errorf("name seharusnya berubah ke 'udin_updated', dapat '%s'", result.Name)
	}
	if result.Grade != 4.0 {
		t.Errorf("grade seharusnya berubah ke 4.0, dapat %v", result.Grade)
	}
	if result.NIM != 434241001 {
		t.Errorf("NIM seharusnya tidak berubah (434241001), dapat %d", result.NIM)
	}
	if !result.IsActive {
		t.Error("is_active seharusnya tidak berubah (true)")
	}
}

func TestIsEmptyPatch(t *testing.T) {
	empty := model.PatchStudentRequest{}
	if !IsEmptyPatch(empty) {
		t.Error("patch kosong seharusnya dianggap kosong")
	}

	newGrade := 3.5
	notEmpty := model.PatchStudentRequest{Grade: &newGrade}
	if IsEmptyPatch(notEmpty) {
		t.Error("patch dengan Grade seharusnya dianggap tidak kosong")
	}
}

func TestValidateCreate(t *testing.T) {
	req := model.CreateStudentRequest{
		NIM:      434241001,
		Name:     "udin",
		Grade:    3.5,
		IsActive: true,
	}
	errs := ValidateCreate(req)
	if len(errs) > 0 {
		t.Errorf("data valid seharusnya tidak ada error: %v", errs)
	}

	reqInvalid := model.CreateStudentRequest{
		NIM:      0,
		Name:     "udin",
		Grade:    3.5,
		IsActive: true,
	}
	errs = ValidateCreate(reqInvalid)
	if _, ok := errs["nim"]; !ok {
		t.Error("NIM <= 0 seharusnya memunculkan error nim")
	}

	reqEmptyName := model.CreateStudentRequest{
		NIM:      434241001,
		Name:     "",
		Grade:    3.5,
		IsActive: true,
	}
	errs = ValidateCreate(reqEmptyName)
	if _, ok := errs["name"]; !ok {
		t.Error("name kosong seharusnya memunculkan error name")
	}
}
