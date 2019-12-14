package model

type UpcomingMovies struct {
	MovieList []Movie `json:"results"`
}
type Movie struct {
	Title       string `json:"title"`
	PosterPath  string `json:"poster_path"`
	Overview    string `json:"overview"`
	Id          int    `json:"id"`
	ReleaseDate string `json:"release_date"`
}
