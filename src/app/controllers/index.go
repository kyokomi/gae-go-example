package controllers

import (
	"app/models"
	"app/views"
	"net/http"

	"appengine"
	"appengine/datastore"
	"appengine/user"
)

func Index(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "GET request only", http.StatusMethodNotAllowed)
		return
	}

	c := appengine.NewContext(r)

	q := datastore.NewQuery("Greeting").Ancestor(models.GuestBookKey(c)).Order("-Date").Limit(10)
	greetings := make([]models.Greeting, 0, 10)
	if _, err := q.GetAll(c, &greetings); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if us, err := user.CurrentOAuth(c, ""); err != nil {
		c.Debugf(err.Error())
	} else {
		c.Debugf("OAuthUser: %v", us)
	}

	if key, err := user.OAuthConsumerKey(c); err != nil {
		c.Debugf(err.Error())
	} else {
		c.Debugf("key: %s", key)
	}

	type exec struct {
		Author    string
		Greetings []models.Greeting
	}
	e := exec{
		Greetings: greetings,
	}
	if u := user.Current(c); u != nil {
		e.Author = u.String()
	}

	if err := views.GuestBookTemplate.Execute(w, e); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
