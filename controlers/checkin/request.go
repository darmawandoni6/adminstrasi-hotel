package checkin

import (
	"administrasi-hotel/busieness/checkin"
	"administrasi-hotel/helpers/conversi"
)

type ReqCheckin struct {
	Name      string `json:"name"`
	Address   string `json:"address"`
	RoomId    int    `json:"room_id"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	Detail    []struct {
		FacilitiesId int `json:"facility_id"`
	} `json:"detail"`
}

func (req *ReqCheckin) ToDomain() *checkin.Domain {
	start, _ := conversi.Date(req.StartDate)
	end, _ := conversi.Date(req.EndDate)

	detail := []checkin.DomainDetail{}

	for i := 0; i < len(req.Detail); i++ {
		detail = append(detail, checkin.DomainDetail{
			FacilitiesId: req.Detail[i].FacilitiesId,
		})
	}

	return &checkin.Domain{
		Name:          req.Name,
		Address:       req.Address,
		RoomId:        req.RoomId,
		StartDate:     start,
		EndDate:       end,
		CheckinDetail: detail,
	}
}
