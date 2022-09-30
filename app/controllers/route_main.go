package controllers

import (
	"net/http"
)

func top(w http.ResponseWriter, r *http.Request) {
	generatedHTML(w, "Hello", "layout", "top")
}
