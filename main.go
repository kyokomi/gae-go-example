package gaehoge

import (
	"fmt"
	"net/http"
	"appengine"
	"appengine/user"
)

func init() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/logout", doLogoutHandler)
}

func handler(w http.ResponseWriter, r *http.Request) {

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
	fmt.Fprint(w, "Hello, %v!", u)
}

func doLogoutHandler(w http.ResponseWriter, r *http.Request) {

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
	http.Redirect(w, r, "/", 200)
}
