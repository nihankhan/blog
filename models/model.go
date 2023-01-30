package models

import (
	"database/sql"
	"log"
	"time"
)

//User is a model for user

type User struct {
	ID       int
	Name     string
	Email    string
	Password []byte
}

// Article is a model for articles.

type Article struct {
	ID        int
	Image     string
	Slug      string
	Title     string
	Content   string
	Author    User
	CreatedAt time.Time
}

var (
	// db is a database connection.
	db *sql.DB

	//err is an error returned.
	err error
)

// FindArticle is to print a article

func FindArticle(slug string) *Article {
	rows, err := db.Query(`SELECT articles.image, articles.title, articles.content, users.name, articles.created_at FROM articles JOIN users ON users.id = articles.author WHERE slug = ?`, slug)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var createdAt []byte

	user := &User{}
	article := &Article{}

	for rows.Next() {
		err = rows.Scan(&article.Image, &article.Title, &article.Content, &user.Name, &createdAt)
		if err != nil {
			log.Fatal(err)
		}

		parsedCreatedAt, err := time.Parse("2023-01-30 08:44:07", string(createdAt))
		if err != nil {
			log.Fatal(err)
		}

		article.CreatedAt = parsedCreatedAt
		article.Author = *user
	}

	return article
}

func Articles() []*Article {
	var articles []*Article

	rows, err := db.Query(`SELECT articles.id, articles.image, articles.slug, articles.title, articles.content, users.name, articles.created_at FROM articles JOIN users ON users.id = articles.author`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			id        int
			image     string
			slug      string
			title     string
			content   string
			author    string
			createdAt []byte
		)

		err = rows.Scan(&id, &image, &slug, &title, &content, &author, &createdAt)
		if err != nil {
			log.Fatal(err)
		}

		parsedCreatedAt, err := time.Parse("2023-01-30 08:44:07", string(createdAt))
		if err != nil {
			log.Fatal(err)
		}

		user := User {
			Name : author,
		}
		articles = append(articles, &Article{id, image, slug, title, content, user, parsedCreatedAt})
	}

	return articles
}

func (user User) Find() *User {
	rows, err := db.Query(`SELECT * FROM users WHERE email = ?`, user.Email)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
		if err != nil {
			log.Fatal(err)
		}
	}

	return &user
}

func (user User) FindArticle(id int) *Article {
	rows, err := db.Query(`SELECT image, slug, title, content, created_at FROM articles WHERE id = ? AND author = ?`, id, user.ID)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var createdAt []byte

	article := &Article{
		ID : id,
		Author: user,
	}

	for rows.Next() {
		err = rows.Scan(&article.Image, &article.Slug, &article.Title, &article.Content, &createdAt)
		if err != nil {
			log.Fatal(err)
		}

		parsedCreatedAt, err := time.Parse("2023-01-30 08:44:07", string(createdAt))
		if err != nil {
			log.Fatal(err)
		}

		article.CreatedAt = parsedCreatedAt
	}

	return article
}


