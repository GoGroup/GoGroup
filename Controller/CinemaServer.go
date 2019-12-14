package controller

// import (
// 	"html/template"
// 	"net/http"
// )

// var templ = template.Must(template.ParseFiles("CinemaName.html"))
// var temp2 = template.Must(template.ParseFiles("CinemaViewer.html"))

// func CinemaNameHandler(w http.ResponseWriter, r *http.Request) {
// 	templ.Execute(w, nil)
// }

// func CinemaViewerHandler(w http.ResponseWriter, r *http.Request) {
// 	temp2.Execute(w, nil)
// }

// func main() {
// 	mux := http.NewServeMux()
// 	fs := http.FileServer(http.Dir("assets"))
// 	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))
// 	mux.HandleFunc("/", CinemaViewerHandler)
// 	mux.HandleFunc("/CinemaName", CinemaNameHandler)
// 	http.ListenAndServe(":8080", mux)
// }
