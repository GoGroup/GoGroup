package main

import ("net/http"
		 "html/template")
var temple=template.Must(template.ParseFiles("index.html"))
func index(w http.ResponseWriter, r *http.Request){
temple.Execute(w,nil)
}
func main(){
	mux:=http.NewServeMux()
	style:=http.FileServer(http.Dir("css/"))
	image:=http.FileServer(http.Dir("image/"))
	js:=http.FileServer(http.Dir("js/"))
	mux.Handle("/js/",http.StripPrefix("/js/",js))
	mux.Handle("/css/",http.StripPrefix("/css/",style))
	mux.Handle("/image/",http.StripPrefix("/image/",image))
	mux.HandleFunc("/",index)
	http.ListenAndServe(":8080",mux)

}
