package controllers

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/gosimple/slug"
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
	permissson(w, r)

	switch r.Method {
	case http.MethodGet:
		var resp bytes.Buffer

		if err := views.Tmpl.ExecuteTemplate(&resp, "login.tpl", nil); err != nil {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		io.WriteString(w, resp.String())

	case http.MethodPost:
		user := &models.User{
			Email: r.FormValue("email"),
		}

		user = user.Find()

		if user.ID == 0 {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)

			return
		}

		err := bcrypt.CompareHashAndPassword(user.Password, []byte(r.FormValue("password")))
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)

			return
		}

		sessionsID := uuid.New().String()
		cookie := &http.Cookie{
			Name:  "session",
			Value: sessionsID,
		}
		Sessions[sessionsID] = user

		http.SetCookie(w, cookie)
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	cookie := getCookie(r)

	if Sessions[cookie.Value] != nil {
		delete(Sessions, cookie.Value)
	}

	cookie = &http.Cookie{
		Name:  "session",
		Value: "",
	}
	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func Dashboard(w http.ResponseWriter, r *http.Request) {
	permissson(w, r)

	cookie := getCookie(r)
	user := Sessions[cookie.Value]
	model := r.URL.Query().Get("model")
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	delete := r.URL.Query().Get("delete")

	switch r.Method {
	case http.MethodGet:
		switch model {
		case "article":
			if delete != "" && id != 0 {
				article := &models.Article{
					ID: id,
				}

				user.DeleteArticle(article)

				http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
			} else {
				data.Article = &models.Article{
					Image:   "",
					Title:   "",
					Content: "",
				}

				if user != nil && id != 0 {
					data.Article = user.FindArticle(id)
				}

				var resp bytes.Buffer

				if err := views.Tmpl.ExecuteTemplate(&resp, "article.tpl", data); err != nil {
					http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)

					return
				}

				io.WriteString(w, resp.String())
			}

		default:
			if user != nil {
				data.Articles = user.FindArticles()
			}

			var resp bytes.Buffer

			if err := views.Tmpl.ExecuteTemplate(&resp, "dashboard.tpl", data); err != nil {
				http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)

				return
			}

			io.WriteString(w, resp.String())
		}
	case http.MethodPost:
		switch model {
		case "article":
			if user != nil {
				if id == 0 {
					article := &models.Article{
						Image: r.FormValue("image"),
						Slug:  slug.Make(r.FormValue("title")),
						Title: r.FormValue("title"),
						Content: r.FormValue("content"),
						Author: *user,
						CreatedAt: time.Now(),
					}

					user.CreateArticle(article)
				} else {
					article := &models.Article{
						ID: id,
						Image: r.FormValue("image"),
						Slug: slug.Make(r.FormValue("title")),
						Title: r.FormValue("title"),
						Content: r.FormValue("content"),
					}

					user.UpdateArticle(article)
				}
			}

			http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
		}
	}
}
