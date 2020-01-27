package model

import "time"

var UserMock = User{
	ID:       1,
	FullName: "user 1",
	Email:    "user@gmail.com",
	Password: "12345678",
	RoleID:   1,
	Amount:   300,
}
var SessionMock = Session{
	ID:         1,
	SessionId:  "session_one",
	UUID:       1,
	SigningKey: []byte("Cinevpark"),
	Expires:    0,
}
var RoleMock = Role{
	ID:   1,
	Name: "Mock Role 01",
}
var CommentMock = Comment{
	ID:        1,
	UserID:    1,
	UserName:  "Mock user",
	MovieID:   1,
	Message:   "THIS IS A MESSAGE",
	CreatedAt: time.Time{},
}
var HallMock = Hall{
	ID:          1,
	HallName:    "Mock Hall1",
	Capacity:    90,
	CinemaID:    1,
	schedules:   []Schedule{},
	Price:       70,
	VIPPrice:    120,
	VIPCapacity: 20,
}

var CinemaMock = Cinema{
	ID:         1,
	CinemaName: "Mock Cinema1",
	Halls:      []Hall{},
}

var BookingMock = Booking{
	ID:      1,
	UserID:  1,
	MovieID: 1,
}
var EvenMock = Event{
	ID:          1,
	Name:        "MockName 1",
	Description: "This is Mock Event",
	Location:    "Addis Ababa",
	Time:        "12:03",
	Image:       "This is  a path",
}

var ScheduleMock = Schedule{
	ID:           1,
	MoviemID:     1,
	StartingTime: "12:09",
	Dimension:    "2D",
	HallID:       1,
	Day:          "Monday",
	Booked:       5,
}

var MockMovie = Moviem{
	TmdbID: 284053,
}

var MockBindedSchedule = BindedSchedule{
	PosterPath: "path/thor.jpg",
	MovieName:  "thor",
	Runtime:    120,
	HallName:   "Mock HallName",
	Day:        "Monday",
	StartTime:  "12:05",
	Dimension:  "2:01",
}
var MockScheduleWithMovie = ScheduleWithMovie{
	Sch:       Schedule{},
	MovieName: "thor",
}

var MockHallSchedule = HallSchedule{
	CinemaName: "Mock Cinema 1",
	All:        []BindedSchedule{},
}
var  MockUpcomingMovies = UpcomingMovies{
	MovieList :[]Movie{},
}
var MockMovie1= Movie {
	Title       :"Thor",
	PosterPath  :"path/thor",
	Overview    :"this is an overview",
	Id          :1,
	ReleaseDate :"01/01/2020",
	Trailer     "yes",
}
var  MockMovieDetails =MovieDetails {
	GenreList  :[]Genre{},
	Overview   :"this is an overview",
	Id         :1,
	Title      :"thor",
	RunTime    :"120",
	PosterPath :"path/thor",
	Trailer    :"yes",
}
type Genre struct {
	Name :"Mock Genre 01"
}
