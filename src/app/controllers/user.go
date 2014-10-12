package controllers

import (
	"net/http"

	"appengine"
	"appengine/user"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "GET request only", http.StatusMethodNotAllowed)
		return
	}

	c := appengine.NewContext(r)

	u := user.Current(c)
	if u == nil {
		url, err := user.LoginURL(c, r.URL.String())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Location", url)
		w.WriteHeader(http.StatusFound)
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "GET request only", http.StatusMethodNotAllowed)
		return
	}

	c := appengine.NewContext(r)

	u := user.Current(c)
	if u != nil {
		url, err := user.LogoutURL(c, r.URL.String())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Location", url)
		w.WriteHeader(http.StatusFound)
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)
}
