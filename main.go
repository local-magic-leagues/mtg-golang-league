package main

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

const (
	DB_USER     = "mtguser"
	DB_PASSWORD = "mtguser"
	DB_NAME     = "test"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", Home).Methods("GET")
	router.HandleFunc("/home", Home).Methods("GET")
	router.HandleFunc("/getuser", GetUser).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}

// Home will return Welocm Home string
func Home(w http.ResponseWriter, r *http.Request) {
	test := []byte("Welcome home")
	w.Write(test)
}

// GetUser will return users from postgres
func GetUser(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Getting users...\n")

	dbInfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)

	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	users, err := db.Query("select * from public.user")
	if err != nil {
		panic(err)
	}
	for users.Next() {
		var (
			userid   string
			username string
		)
		if err := users.Scan(&userid, &username); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("UserID: %s \nUserName: %s", string(userid), username)
	}
}
