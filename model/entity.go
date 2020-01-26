package model

import "github.com/jinzhu/gorm"

type UpcomingMovies struct {
	MovieList []Movie `json:"results"`
}
type Movie struct {
	Title       string `json:"title"`
	PosterPath  string `json:"poster_path"`
	Overview    string `json:"overview"`
	Id          int    `json:"id"`
	ReleaseDate string `json:"release_date"`
	Trailer     string
}
type MovieDetails struct {
	GenreList  []Genre `json:"genres"`
	Overview   string  `json:"overview"`
	Id         int     `json:"id"`
	Title      string  `json:"original_title"`
	RunTime    int     `json:"runtime"`
	PosterPath string  `json:"poster_path"`
	Trailer    string
}
type Genre struct {
	Name string `json:"name"`
}

type VideoLists struct {
	VList []Video `json:"results"`
}
type Video struct {
	Key string `json:"key"`
}
type Schedule struct {
	ID           uint   `json: "id" gorm:"primary_key;AUTO_INCREMENT"`
	MoviemID     int    `json: "moviemid"`
	StartingTime string `json: "startingtime" gorm:"type:varchar(255);not null"`
	Dimension    string `json: "dimension" gorm:"type:varchar(255);not null"`
	HallID       int    `json: "hallid"`
	Day          string `json: "day" gorm:"type:varchar(255);not null"`
}
type Moviem struct {
	TmdbID int `json: "TmdbID"  gorm:"unique"`
}

type HallSchedule struct {
	CinemaName string
	All        []BindedSchedule
}
type BindedSchedule struct {
	PosterPath, MovieName string
	Runtime               int
	HallName              string
	Day                   string
	StartTime             string
}
type ScheduleWithMovie struct {
	Sch       Schedule
	MovieName string
}
type Hall struct {
	ID              uint   `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	HallName        string `json:"hallname" gorm:"type:varchar(255);not null"`
	Capacity        uint   `json:"capacity"`
	CinemaID        uint   `json:"cinemaid"`
	schedules       []Schedule
	Price           uint `json:"Price"`
	VIPPrice        uint `json:"vipprice"`
	WeekendDiscount uint `json:"discount"`
}
type Cinema struct {
	ID         uint   `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	CinemaName string ` json:"cinemaname" gorm:"type:varchar(255);not null"`
	Halls      []Hall
}
type Role struct {
	ID   uint
	Name string `gorm:"type:varchar(255)"`
}
type User struct {
	ID       uint   `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	FullName string `gorm:"type:varchar(255);not null"`
	Email    string `gorm:"type:varchar(255);not null;unique"`
	Password string `gorm:"type:varchar(255)"`
	RoleID   uint
}
type Session struct {
	gorm.Model
	SessionId  string `gorm:"type:varchar(255);not null"`
	UUID       uint
	Expires    int64  `gorm:"type:varchar(255);not null"`
	SigningKey []byte `gorm:"type:varchar(255);not null"`
}
type Comment struct {
	gorm.Model
	UserID   uint
	UserName string
	MovieID  uint
	Message  string ` json:"message" gorm:"type:varchar(255);not null"`
}

type Event struct {
	ID          uint   `json:"id"  gorm:"primary_key;AUTO_INCREMENT"`
	Name        string `json:"name" gorm:"type:varchar(255);not null"`
	Description string `json:"description"`
	Location    string `json:"location"`
	Time        string `json:"time"`
	Image       string `json:"image" gorm:"type:varchar(255)"`
}
