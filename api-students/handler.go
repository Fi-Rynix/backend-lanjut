package apistudents

import (
	"sort"
	"strconv"
	"strings"
	"github.com/gofiber/fiber/v2"
)

var students []Student
var nextID = 1

func findStudentIndex(id int) int {
	for i := range students {
		if students[i].ID == id {
			return i
		}
	}
	return -1
}

func cocokPencarian(s Student, kata string) bool {
	kata = strings.ToLower(kata)
	return strings.Contains(strings.ToLower(s.Name), kata) ||
		strings.Contains(strings.ToLower(strconv.Itoa(s.NIM)), kata)
}

func paramID(c *fiber.Ctx) (int, bool) {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil || id < 1 {
		return 0, false
	}
	return id, true
}

func listStudents(c *fiber.Ctx) error {
	q := parseListQuery(c)
	filtered := []Student{}

	for _, s := range students {
		if q.IsActive != nil && s.IsActive != *q.IsActive {
			continue
		}
		if q.Search != "" && !cocokPencarian(s, q.Search) {
			continue
		}
		filtered = append(filtered, s)
	}
	
	sort.SliceStable(filtered, func(i, j int) bool {
		var smaller bool
		switch q.Sort {
		case "nim":
			smaller = filtered[i].NIM < filtered[j].NIM
		case "name":
			smaller = strings.ToLower(filtered[i].Name) < strings.ToLower(filtered[j].Name)
		case "grade":
			smaller = filtered[i].Grade < filtered[j].Grade
		default:
			smaller = filtered[i].ID < filtered[j].ID
		}
		if q.Order == "desc" {
			return !smaller
		}
		return smaller
	})

	total := len(filtered)
	totalPages := (total + q.Limit - 1) / q.Limit
	if totalPages < 1 {
		totalPages = 1
	}
	start := (q.Page - 1) * q.Limit
	if start > total {
		start = total
	}
	end := start + q.Limit
	if end > total {
		end = total
	}

	return okList(c, "daftar student berhasil diambil", filtered[start:end], &Meta{
		Page:       q.Page,
		Limit:      q.Limit,
		Total:      total,
		TotalPages: totalPages,
	})
}

func getStudent(c *fiber.Ctx) error {
	id, valid := paramID(c)
	if !valid {
		return fail(c, fiber.StatusBadRequest, "id harus berupa angka positif")
	}

	i := findStudentIndex(id)
	if i == -1 {
		return fail(c, fiber.StatusNotFound, "student tidak ditemukan")
	}

	return ok(c, "student ditemukan", students[i])
}

func createStudent(c *fiber.Ctx) error {
	var req CreateStudentRequest
	if err := c.BodyParser(&req); err != nil {
		return fail(c, fiber.StatusBadRequest, "body harus berupa JSON yang valid")
	}

	errs := map[string]string{}
	if req.NIM <= 0 {
		errs["nim"] = "NIM wajib diisi dan berupa angka positif"
	}
	if strings.TrimSpace(req.Name) == "" {
		errs["name"] = "name wajib diisi"
	}
	for _, s := range students {
		if s.NIM == req.NIM {
			errs["nim"] = "NIM sudah dipakai"
		}
	}
	if len(errs) > 0 {
		return failValidation(c, errs)
	}

	baru := Student{
		ID:       nextID,
		NIM:      req.NIM,
		Name:     strings.TrimSpace(req.Name),
		Grade:    req.Grade,
		IsActive: req.IsActive,
	}
	students = append(students, baru)
	nextID++

	return created(c, "student berhasil dibuat", baru,
		"/api/v1/students/"+strconv.Itoa(baru.ID))
}

func replaceStudent(c *fiber.Ctx) error {
	id, valid := paramID(c)
	if !valid {
		return fail(c, fiber.StatusBadRequest, "id harus berupa angka positif")
	}

	i := findStudentIndex(id)
	if i == -1 {
		return fail(c, fiber.StatusNotFound, "student tidak ditemukan")
	}

	var req UpdateStudentRequest
	if err := c.BodyParser(&req); err != nil {
		return fail(c, fiber.StatusBadRequest, "body harus berupa JSON yang valid")
	}

	errs := map[string]string{}
	if req.NIM <= 0 {
		errs["nim"] = "NIM wajib diisi pada PUT"
	}
	if strings.TrimSpace(req.Name) == "" {
		errs["name"] = "name wajib diisi pada PUT"
	}
	if len(errs) > 0 {
		return failValidation(c, errs)
	}

	students[i].NIM = req.NIM
	students[i].Name = strings.TrimSpace(req.Name)
	students[i].Grade = req.Grade
	students[i].IsActive = req.IsActive

	return ok(c, "student berhasil diganti seluruhnya", students[i])
}

func patchStudent(c *fiber.Ctx) error {
	id, valid := paramID(c)
	if !valid {
		return fail(c, fiber.StatusBadRequest, "id harus berupa angka positif")
	}

	i := findStudentIndex(id)
	if i == -1 {
		return fail(c, fiber.StatusNotFound, "student tidak ditemukan")
	}

	var req PatchStudentRequest
	if err := c.BodyParser(&req); err != nil {
		return fail(c, fiber.StatusBadRequest, "body harus berupa JSON yang valid")
	}

	if req.NIM == nil && req.Name == nil && req.Grade == nil && req.IsActive == nil {
		return fail(c, fiber.StatusBadRequest, "tidak ada field yang diubah")
	}

	if req.NIM != nil {
		if *req.NIM <= 0 {
			return failValidation(c, map[string]string{"nim": "NIM harus angka positif"})
		}
		for _, s := range students {
			if s.ID != id && s.NIM == *req.NIM {
				return fail(c, fiber.StatusConflict, "NIM sudah dipakai student lain")
			}
		}
		students[i].NIM = *req.NIM
	}
	if req.Name != nil {
		if strings.TrimSpace(*req.Name) == "" {
			return failValidation(c, map[string]string{"name": "name tidak boleh kosong"})
		}
		students[i].Name = strings.TrimSpace(*req.Name)
	}
	if req.Grade != nil {
		students[i].Grade = *req.Grade
	}
	if req.IsActive != nil {
		students[i].IsActive = *req.IsActive
	}

	return ok(c, "student berhasil diperbarui sebagian", students[i])
}

func deleteStudent(c *fiber.Ctx) error {
	id, valid := paramID(c)
	if !valid {
		return fail(c, fiber.StatusBadRequest, "id harus berupa angka positif")
	}

	i := findStudentIndex(id)
	if i == -1 {
		return fail(c, fiber.StatusNotFound, "student tidak ditemukan")
	}

	students = append(students[:i], students[i+1:]...)

	return noContent(c)
}
