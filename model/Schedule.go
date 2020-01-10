package model

type Schedule struct {
	ID           uint   `json: "id" gorm:"primary_key;AUTO_INCREMENT"`
	MoviezID     int    `json: "moviezid"`
	StartingTime string `json: "startingtime" gorm:"type:varchar(255);not null"`
	Dimension    string `json: "dimension" gorm:"type:varchar(255);not null"`
	HallID       int    `json: "hallid"`
	Day          string `json: "day" gorm:"type:varchar(255);not null"`
}
type Moviez struct {
	ID   uint `json: "id" gorm:"primary_key;AUTO_INCREMENT"`
	Tmdb int  `json: "tmdb"`
}
