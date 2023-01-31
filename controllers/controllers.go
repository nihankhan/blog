package controllers

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/nihankhan/blog/helpers"
	"github.com/nihankhan/blog/models"
	"github.com/nihankhan/blog/views"
	"golang.org/x/crypto/bcrypt"
)

var data struct {
	Article  *models.Article
	Articles []*models.Article
}

var Sessions map[string]*models.User

func getCookie(r *http.Request) *http.Cookie {
	cookie := &http.Cookie{
		Name:  "session",
		Value: "",
	}

	for _, c := range r.Cookies() {
		if c.Name == "session" {
			cookie.Value = c.Value
			break
		}
	}

	return cookie
}

func permissson(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.URL.Path, "/")[1]
	cookie := getCookie(r)

	switch path {
	case "dashboard":
		if Sessions[cookie.Value] == nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		}
	case "login":
	case "register":
		if Sessions[cookie.Value] != nil {
			http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
		}
	}
}

func Index(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.URL.Path, "/")[1]

	if path == "" {
		data.Articles = models.Articles()

		var resp bytes.Buffer

		if err := views.Tmpl.ExecuteTemplate(&resp, "signup.tpl", data); err != nil {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)

			return
		}

		io.WriteString(w, resp.String())
	} else {
		data.Article = models.FindArticle(path)

		var resp bytes.Buffer

		if err := views.Tmpl.ExecuteTemplate(&resp, "single.tpl", data); err != nil {
			log.Fatal(err)
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)

			return
		}

		io.WriteString(w, resp.String())
	}
}

// Register is a router to register a new user.

func Signup(w http.ResponseWriter, r *http.Request) {
	permissson(w, r)

	switch r.Method {
	case http.MethodGet:
		var resp bytes.Buffer

		if err := views.Tmpl.ExecuteTemplate(&resp, "signup.tpl", data); err != nil {
			log.Fatal(err)
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)

			return
		}
		io.WriteString(w, resp.String())

	case http.MethodPost:
		user := &models.User{
			Name:  r.FormValue("name"),
			Email: r.FormValue("email"),
		}

		password, err := bcrypt.GenerateFromPassword([]byte(r.FormValue("password")), bcrypt.MinCost)

		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		if err := helpers.ValidateEmail(user.Email); err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		user.Password = password
		user = user.Find()

		if user.ID == 0 {
			user = user.Create()
		}

		sessionID := uuid.New().String()

		cookie := &http.Cookie{
			Name:  "session",
			Value: sessionID,
		}

		Sessions[sessionID] = user

		http.SetCookie(w, cookie)

		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
	}
}

func Login(w http.ResponseWriter, r *http.Request) {

}
