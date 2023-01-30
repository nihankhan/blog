package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", route.Index)
	mux.HandleFunc("/login", route.Login)
	mux.HandleFunc("/logout", route.Logout)
	mux.HandleFunc("signup", route.Signup)
	mux.HandleFunc("/dashboard", route.Dashboard)

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

}
