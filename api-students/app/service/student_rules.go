package service

import (
	"strings"

	"api-students/app/model"
)

func ValidateCreate(req model.CreateStudentRequest) map[string]string {
	errs := map[string]string{}

	req.Name = strings.TrimSpace(req.Name)

	if req.NIM <= 0 {
		errs["nim"] = "NIM wajib diisi dan berupa angka positif"
	}
	if req.Name == "" {
		errs["name"] = "name wajib diisi"
	}

	return errs
}

func ValidateReplace(req model.UpdateStudentRequest) map[string]string {
	errs := map[string]string{}

	req.Name = strings.TrimSpace(req.Name)

	if req.NIM <= 0 {
		errs["nim"] = "NIM wajib diisi pada PUT"
	}
	if req.Name == "" {
		errs["name"] = "name wajib diisi pada PUT"
	}

	return errs
}

func ApplyPatch(current model.Student, req model.PatchStudentRequest) (model.Student, map[string]string) {
	errs := map[string]string{}

	if req.NIM != nil {
		if *req.NIM <= 0 {
			errs["nim"] = "NIM harus angka positif"
		} else {
			current.NIM = *req.NIM
		}
	}
	if req.Name != nil {
		if strings.TrimSpace(*req.Name) == "" {
			errs["name"] = "name tidak boleh kosong"
		} else {
			current.Name = strings.TrimSpace(*req.Name)
		}
	}
	if req.Grade != nil {
		current.Grade = *req.Grade
	}
	if req.IsActive != nil {
		current.IsActive = *req.IsActive
	}

	return current, errs
}

func IsEmptyPatch(req model.PatchStudentRequest) bool {
	return req.NIM == nil && req.Name == nil && req.Grade == nil && req.IsActive == nil
}

func CountTotalPages(total, limit int) int {
	if limit <= 0 {
		return 0
	}
	return (total + limit - 1) / limit
}
