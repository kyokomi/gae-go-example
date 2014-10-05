package gaehoge

import (
	"appengine"
	"appengine/datastore"
	"time"
)

type Greeting struct {
	Author  string
	Content string
	Date    time.Time
}

func guestBookKey(c appengine.Context) *datastore.Key {
	return datastore.NewKey(c, "Guestbook", "defualt_guestbook", 0, nil)
}
