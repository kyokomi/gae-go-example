package gaehoge

import (
	"net/http"
	"time"

	"appengine"
	"appengine/datastore"
	"appengine/user"
)

func init() {
	http.HandleFunc("/", doIndex)

	http.HandleFunc("/guest-book/sign", doSign)
}

func doIndex(w http.ResponseWriter, r *http.Request) {

	c := appengine.NewContext(r)

	q := datastore.NewQuery("Greeting").Ancestor(GuestBookKey(c)).Order("-Date").Limit(10)
	greetings := make([]Greeting, 0, 10)
	if _, err := q.GetAll(c, &greetings); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	type exec struct {
		Author    string
		Greetings []Greeting
	}
	e := exec{
		Greetings: greetings,
	}
	if u := user.Current(c); u != nil {
		e.Author = u.String()
	}

	if err := GuestBookTemplate.Execute(w, e); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func doSign(w http.ResponseWriter, r *http.Request) {

	c := appengine.NewContext(r)

	g := Greeting{
		Content: r.FormValue("content"),
		Date:    time.Now(),
	}

	if u := user.Current(c); u != nil {
		g.Author = u.String()
	}

	key := datastore.NewIncompleteKey(c, "Greeting", GuestBookKey(c))
	_, err := datastore.Put(c, key, &g)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusFound)
}
