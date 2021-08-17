package facilities

import "administrasi-hotel/busieness/facilities"

type ReqFacilities struct {
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	IsActive bool    `json:"is_active"`
}

func (req *ReqFacilities) ToDomain() *facilities.Domain {
	return &facilities.Domain{
		Name:     req.Name,
		Price:    req.Price,
		IsActive: req.IsActive,
	}
}
