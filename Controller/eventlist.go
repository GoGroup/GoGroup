package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	_ "github.com/lib/pq/oid"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"

	"github.com/pkg/errors"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1234567890"
	dbname   = "Table1"
)

//url :=" postgresql://[postgres[:1234567890]@][localhost][:5432][/Table1][?param1=value1&...]"
const conString = "postgres://postgres:1234567890@localhost:5432/Table1"

// var temple = template.Must(template.ParseFiles("header.html", "footer.html"))

var (
	connectionString = flag.String("conn", getenvWithDefault("DATABASE_URL", conString), "PostgreSQL connection string")
	listenAddr       = flag.String("addr", getenvWithDefault("LISTENADDR", ":8080"), "HTTP address to listen on")
	db               *sqlx.DB
	tmpl             = template.New("")
)

type ContactFavorites struct {
	Colors []string `json:"colors"`
}

// Contact represents a Contact model in the database
type Contact struct {
	ID                   int
	Name, Location, Type string
}

func getenvWithDefault(name, defaultValue string) string {
	val := os.Getenv(name)
	if val == "" {
		val = defaultValue
	}

	return val
}

func handler(w http.ResponseWriter, r *http.Request) {
	// temple.Execute(w, nil)
	contacts, err := fetchContacts()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	tmpl.ExecuteTemplate(w, "eventlist.html", struct{ Contacts []*Contact }{contacts})

}

func fetchContacts() ([]*Contact, error) {
	contacts := []*Contact{}
	err := db.Select(&contacts, "select * from contacts")
	if err != nil {
		return nil, errors.Wrap(err, "Unable to fetch contacts")
	}

	return contacts, nil
}

func NewDB(connectionString string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("pgx", connectionString)

	if err != nil {
		return nil, err
	}

	return db, nil
}
func main() {
	flag.Parse()
	var err error
	tmpl.Funcs(template.FuncMap{"StringsJoin": strings.Join})
	_, err = tmpl.ParseGlob(filepath.Join(".", "templates", "*.html"))
	if err != nil {
		log.Fatalf("Unable to parse templates: %v\n", err)
	}

	if *connectionString == "" {
		log.Fatalln("Please pass the connection string using the -conn option")
	}

	db, err = sqlx.Connect("pgx", *connectionString)
	if err != nil {
		log.Fatalf("Unable to establish connection: %v\n", err)
	}

	fs := http.FileServer(http.Dir("asset"))
	http.Handle("/asset/", http.StripPrefix("/asset/", fs))
	http.HandleFunc("/", handler)
	log.Printf("listening on %s\n", *listenAddr)
	http.ListenAndServe(*listenAddr, nil)
}
