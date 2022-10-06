package controllers

import (
	"fmt"
	"main/app/models"
	"main/config"
	"net/http"
	"regexp"
	"strconv"
	"text/template"
)

func generatedHTML(w http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("app/views/templates/%s.html", file))
	}
	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, "layout", data)
}

func session(w http.ResponseWriter, r *http.Request) (ses models.Session, err error) {
	cookie, err := r.Cookie("_cookie")
	if err == nil {
		ses = models.Session{UUID: cookie.Value}
		if ok, _ := ses.CheckSession(); !ok {
			err = fmt.Errorf("invalid session")
		}
	}
	return ses, err
}

var validPath = regexp.MustCompile("^/(todos|me)/(edit|update|delete)/([0-9]+)$")

func parseURL(fn func(http.ResponseWriter, *http.Request, int)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		q := validPath.FindStringSubmatch(r.URL.Path)
		if q == nil {
			http.NotFound(w, r)
			return
		}
		id, _ := strconv.Atoi(q[2])
		fn(w, r, id)
	}
}

func StartMainServer() error {
	files := http.FileServer(http.Dir(config.Config.Static))
	http.Handle("/static/", http.StripPrefix("/static/", files))
	http.HandleFunc("/", top)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/authenticate", authenticate)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/todos", index)
	http.HandleFunc("/todos/new", todoNew)
	http.HandleFunc("/todos/save", todoSave)
	http.HandleFunc("/todos/delete/", parseURL(todoDelete))
	http.HandleFunc("/todos/edit/", parseURL(todoEdit))
	http.HandleFunc("/todos/update/", parseURL(todoUpdate))
	http.HandleFunc("/me", myPage)
	http.HandleFunc("/me/update/", parseURL(myPageUpdate))
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
