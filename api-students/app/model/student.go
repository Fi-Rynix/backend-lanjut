package model

import "time"

type Student struct {
	ID        int
	NIM       int
	Name      string
	Grade     float64
	IsActive  bool
	CreatedAt time.Time
}

type CreateStudentRequest struct {
	NIM      int
	Name     string
	Grade    float64
	IsActive bool
}

type UpdateStudentRequest struct {
	NIM      int
	Name     string
	Grade    float64
	IsActive bool
}

type PatchStudentRequest struct {
	NIM      *int     `json:"nim,omitempty"`
	Name     *string  `json:"name,omitempty"`
	Grade    *float64 `json:"grade,omitempty"`
	IsActive *bool    `json:"is_active,omitempty"`
}

type WebResponse struct {
	Success bool
	Message string
	Data    any
	Meta    *Meta
	Errors  any
}

type Meta struct {
	Page       int
	Limit      int
	Total      int
	TotalPages int
}

type ListQuery struct {
	Page     int
	Limit    int
	Search   string
	Sort     string
	Order    string
	IsActive *bool
}

func (q ListQuery) Offset() int {
	return (q.Page - 1) * q.Limit
}
