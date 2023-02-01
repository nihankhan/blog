package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/nihankhan/blog/config"
	con "github.com/nihankhan/blog/controllers"
	"github.com/nihankhan/blog/models"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", con.Index)
	mux.HandleFunc("/login", con.Login)
	mux.HandleFunc("/logout", con.Logout)
	mux.HandleFunc("signup", con.Signup)
	mux.HandleFunc("/dashboard", con.Dashboard)

	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))

	log.Println(fmt.Sprintf("Your app is running on port %s.\n", os.Getenv("PORT")))
	log.Println(http.ListenAndServe(":"+os.Getenv("PORT"), mux))

}

func init() {
	con.Sessions = make(map[string]*models.User)

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	config.Connect()
}
