package model

type Hall struct {
	ID       uint   `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	HallName string `json:"hallname" gorm:"type:varchar(255);not null"`
	Capacity uint   `json:"capacity"`
	CinemaID uint   `json:"cinemaid"`

	Price           uint `json:"Price"`
	VIPPrice        uint `json:"vipprice"`
	WeekendDiscount uint `json:"discount"`
}
type Cinema struct {
	ID         uint   `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	CinemaName string ` json:"cinemaname" gorm:"type:varchar(255);not null"`
	Halls      []Hall
}
