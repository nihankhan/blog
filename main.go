package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	con "github.com/nihankhan/blog/controllers"
	"github.com/nihankhan/blog/models"
)

const PORT = ":8000"

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", con.Index)
	mux.HandleFunc("/login", con.Login)
	mux.HandleFunc("/logout", con.Logout)
	mux.HandleFunc("/signup", con.Signup)
	mux.HandleFunc("/dashboard", con.Dashboard)

	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))

	log.Println(fmt.Sprintf("Your app is Running on PORT %s\n", PORT))

	log.Fatal(http.ListenAndServe(PORT, mux))

}

func init() {
	con.Sessions = make(map[string]*models.User)

	models.Db, models.Err = sql.Open("mysql", "root:nihan@tcp(127.0.0.1:3306)/blog")
	if models.Err != nil {
		log.Fatal(models.Err)
	}

	models.Err = models.Db.Ping()
	if models.Err != nil {
		log.Fatal(models.Err)
	}
}
