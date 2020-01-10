package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

var db *gorm.DB //used to store connection object

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "Bangtan123"
	dbname   = "Movie_and_Events"
)

//crud operations
// func Inserthall(name string, id int, cap int, price int, vip int, discount int) {
// 	var stat = "INSERT INTO hall (hall_name,cinema_id,capacity,price,vip_price,weekend_percent)VALUES($1,$2,$3,$4,$5,$6)"
// 	_, err := db.Exec(stat, name, id, cap, price, vip, discount)
// 	if err != nil {
// 		panic(err)
// 	}
// }

// func Read(db *gorm.DB) string {
// 	var stat = "SELECT full_name FROM users WHERE id=$1"
// 	row := db.QueryRow(stat, 2)
// 	var name string
// 	switch err := row.Scan(&name); err {
// 	case sql.ErrNoRows:
// 		fmt.Println("no rows")
// 		return "no"
// 	case nil:
// 		return name

// 	default:
// 		panic(err)
// 	}

// }

// func Retrive(db *sql.DB) (int, string, string) {
//   //retrive
//   sqlStatement := SELECT id, email, phone FROM users WHERE id=$1;
//   var email string
//   var id int
//   var phone string

//   row := db.QueryRow(sqlStatement, 1)
//   switch err := row.Scan(&id, &email, &phone); err {
//   case sql.ErrNoRows:
//     fmt.Println("No rows were returned!")
//     return 0, "", ""
//   case nil:
//     return id, email, phone
//   default:
//     panic(err)
//   }
// }

// func Update(db *sql.DB) bool {
//   //update
//   sqlStatement := UPDATE users SET phone = $2 WHERE id = $1;

//   _, err := db.Exec(sqlStatement, 1, "+251911121314")
//   if err != nil {
//     panic(err)
//   }

//   return true
// }

func dbConn() (db *gorm.DB) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	//defer db.Close()
	fmt.Println("DB Connected sucessfully !")
	return db

}
func init() {
	db = dbConn()
}

// func main() {

// 	//Insert(db)
// 	//println("i = ", i)

// 	// i, e, p := Retrive(db)
// 	i := Read(db)
// 	println(i)

// 	// b := Update(db)
// 	// i, e, p := Retrive(db)
// 	// println(b, i, e, p)

// 	// d := Delete(db)
// 	// i, e, p := Retrive(db)
// 	// println(d, i, e, p)
// 	db = dbConn()
// }
func main() {

}
