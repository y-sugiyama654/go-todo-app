package controllers

import (
	"log"
	"main/app/models"
	"net/http"
)

func signup(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		generatedHTML(w, nil, "layout", "public_navbar", "signup")
	} else if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Fatalln(err)
		}
		user := models.User{
			Name:     r.PostFormValue("name"),
			Email:    r.PostFormValue("email"),
			PassWord: r.PostFormValue("password"),
		}
		if err := user.CreateUser(); err != nil {
			log.Fatalln(err)
		}

		http.Redirect(w, r, "/", 302)
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	generatedHTML(w, nil, "layout", "public_navbar", "login")
}

func authenticate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	user, err := models.GetUserByEmail(r.PostFormValue("email"))
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/login", 302)
	}
	if user.PassWord == models.Encrypt(r.PostFormValue("password")) {
		session, err := user.CreateSession()
		if err != nil {
			log.Println(err)
		}
		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    session.UUID,
			HttpOnly: true,
		}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/", 302)
	} else {
		http.Redirect(w, r, "/login", 302) 
	}
}
