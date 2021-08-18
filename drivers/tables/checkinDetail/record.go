package checkinDetail

import (
	"administrasi-hotel/drivers/tables/facilities"
	"time"
)

type CheckinDetail struct {
	Id           int
	CheckinId    int
	FacilitiesId int
	Facilities   facilities.Facilities `gorm:"foreignKey:FacilitiesId"`
	IsCheckout   bool
	CreatedAt    time.Time `gorm:"<-:create"`
	UpdatedAt    time.Time
}
