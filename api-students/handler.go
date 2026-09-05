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