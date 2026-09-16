package service

import (
	"testing"

	"api-students/app/model"
)

func TestValidateRegister(t *testing.T) {
	t.Run("data valid", func(t *testing.T) {
		req := model.RegisterRequest{
			Username: "udin123",
			Email:    "udin@example.com",
			Password: "rahasia123",
		}
		errs := ValidateRegister(req)
		if len(errs) > 0 {
			t.Errorf("data valid seharusnya tidak ada error, dapat: %v", errs)
		}
	})

	t.Run("username kosong", func(t *testing.T) {
		req := model.RegisterRequest{
			Username: "",
			Email:    "udin@example.com",
			Password: "rahasia123",
		}
		errs := ValidateRegister(req)
		if _, ok := errs["username"]; !ok {
			t.Error("username kosong seharusnya ada error")
		}
	})

	t.Run("username terlalu pendek", func(t *testing.T) {
		req := model.RegisterRequest{
			Username: "ud",
			Email:    "udin@example.com",
			Password: "rahasia123",
		}
		errs := ValidateRegister(req)
		if _, ok := errs["username"]; !ok {
			t.Error("username 2 karakter seharusnya error")
		}
	})

	t.Run("email tidak valid", func(t *testing.T) {
		req := model.RegisterRequest{
			Username: "udin123",
			Email:    "bukan-email",
			Password: "rahasia123",
		}
		errs := ValidateRegister(req)
		if _, ok := errs["email"]; !ok {
			t.Error("email tanpa @ seharusnya error")
		}
	})

	t.Run("password lemah", func(t *testing.T) {
		req := model.RegisterRequest{
			Username: "udin123",
			Email:    "udin@example.com",
			Password: "password1",
		}
		errs := ValidateRegister(req)
		if _, ok := errs["password"]; !ok {
			t.Error("password umum seharusnya error")
		}
	})

	t.Run("password terlalu pendek", func(t *testing.T) {
		req := model.RegisterRequest{
			Username: "udin123",
			Email:    "udin@example.com",
			Password: "abc",
		}
		errs := ValidateRegister(req)
		if _, ok := errs["password"]; !ok {
			t.Error("password < 8 karakter seharusnya error")
		}
	})

	t.Run("password tanpa angka", func(t *testing.T) {
		req := model.RegisterRequest{
			Username: "udin123",
			Email:    "udin@example.com",
			Password: "hanyahuruf",
		}
		errs := ValidateRegister(req)
		if _, ok := errs["password"]; !ok {
			t.Error("password tanpa angka seharusnya error")
		}
	})

	t.Run("username karakter tidak valid", func(t *testing.T) {
		req := model.RegisterRequest{
			Username: "udin@123",
			Email:    "udin@example.com",
			Password: "rahasia123",
		}
		errs := ValidateRegister(req)
		if _, ok := errs["username"]; !ok {
			t.Error("username dengan @ seharusnya error")
		}
	})
}

func TestValidateLogin(t *testing.T) {
	t.Run("data valid", func(t *testing.T) {
		req := model.LoginRequest{
			Username: "udin",
			Password: "rahasia123",
		}
		errs := ValidateLogin(req)
		if len(errs) > 0 {
			t.Errorf("data valid seharusnya tidak ada error, dapat: %v", errs)
		}
	})

	t.Run("username kosong", func(t *testing.T) {
		req := model.LoginRequest{
			Username: "",
			Password: "rahasia123",
		}
		errs := ValidateLogin(req)
		if _, ok := errs["username"]; !ok {
			t.Error("username kosong seharusnya error")
		}
	})

	t.Run("password kosong", func(t *testing.T) {
		req := model.LoginRequest{
			Username: "udin",
			Password: "",
		}
		errs := ValidateLogin(req)
		if _, ok := errs["password"]; !ok {
			t.Error("password kosong seharusnya error")
		}
	})
}

func TestIsValidEmail(t *testing.T) {
	t.Run("email valid", func(t *testing.T) {
		emails := []string{
			"udin@example.com",
			"a@b.c",
			"user.name@example.co.id",
		}
		for _, e := range emails {
			if !isValidEmail(e) {
				t.Errorf("'%s' seharusnya valid", e)
			}
		}
	})

	t.Run("email tidak valid", func(t *testing.T) {
		emails := []string{
			"",
			"bukan-email",
			"@example.com",
			"udin@",
			"udin@example",
		}
		for _, e := range emails {
			if isValidEmail(e) {
				t.Errorf("'%s' seharusnya tidak valid", e)
			}
		}
	})
}

func TestWeakPassword(t *testing.T) {
	t.Run("password lemah dideteksi", func(t *testing.T) {
		weak := []string{"password1", "12345678", "qwerty123", "ADMIN123"}
		for _, p := range weak {
			if checkPasswordStrength(p) == "" {
				t.Errorf("'%s' seharusnya dianggap lemah", p)
			}
		}
	})

	t.Run("password kuat tidak ditolak", func(t *testing.T) {
		strong := []string{"RahasiaBanget123", "udinSecure2024", "MyP@ssw0rd"}
		for _, p := range strong {
			if checkPasswordStrength(p) != "" {
				t.Errorf("'%s' seharusnya diterima, dapat error: %s", p, checkPasswordStrength(p))
			}
		}
	})
}
