package model

type UpcomingMovies struct {
	MovieList []Movie `json:"results"`
}
type Movie struct {
	Title       string `json:"title "`
	PosterPath  string `json:"poster_path"`
	Overview    string `json:"overview"`
	Id          int    `json:"id"`
	ReleaseDate string `json:"release_date"`
	Trailer     string
}
type MovieDetails struct {
	GenreList []Genre `json:"genres"`
	Overview  string  `json:"overview"`
	Title     string  `json:"original_title"`
	Trailer   string
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
