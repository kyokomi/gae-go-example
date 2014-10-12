package models

import (
	"time"

	"appengine"
	"appengine/datastore"
)

type Greeting struct {
	Author  string
	Content string
	Date    time.Time
}

func GuestBookKey(c appengine.Context) *datastore.Key {
	return datastore.NewKey(c, "Guestbook", "defualt_guestbook", 0, nil)
}
