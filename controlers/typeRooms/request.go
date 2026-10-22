package typeRooms

import "administrasi-hotel/busieness/typeRooms"

type ReqTypeRooms struct {
	Name     string `json:"name"`
	IsActive bool   `json:"is_active"`
}

func (req *ReqTypeRooms) ToDomain() *typeRooms.Domain {
	return &typeRooms.Domain{
		Name:     req.Name,
		IsActive: req.IsActive,
	}
}
