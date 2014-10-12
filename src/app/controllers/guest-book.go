package controllers

import (
	"app/models"
	"appengine"
	"appengine/datastore"
	"appengine/user"
	"net/http"
	"time"
)

func GuestSign(w http.ResponseWriter, r *http.Request) {

	c := appengine.NewContext(r)

	g := models.Greeting{
		Content: r.FormValue("content"),
		Date:    time.Now(),
	}

	if u := user.Current(c); u != nil {
		g.Author = u.String()
	}

	key := datastore.NewIncompleteKey(c, "Greeting", models.GuestBookKey(c))
	_, err := datastore.Put(c, key, &g)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusFound)
}
