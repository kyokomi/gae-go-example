package controllers

import (
	"net/http/httptest"
	"testing"

	"net/url"
	"strings"

	"app/models"
	"reflect"
	"time"

	"appengine"
	"appengine/aetest"
	"appengine/datastore"
	"appengine/user"
)

func TestSign(t *testing.T) {
	now := time.Now()

	inst, err := aetest.NewInstance(nil)
	if err != nil {
		t.Error(err)
	}
	defer inst.Close()

	testCases := []struct {
		method   string
		content  string
		user     *user.User
		code     int
		greeting *models.Greeting
	}{
		{
			method: "GET",
			code:   405,
		},
		{
			method:   "POST",
			content:  "Normal post",
			code:     303,
			greeting: &models.Greeting{Content: "Normal post"},
		},
		{
			method:   "POST",
			content:  "Post with user",
			user:     &user.User{Email: "hoge@gmail.com"},
			code:     303,
			greeting: &models.Greeting{Content: "Post with user", Author: "hoge@gmail.com"},
		},
	}

	for _, tt := range testCases {
		// create RequestData
		val := url.Values{
			"content": []string{tt.content},
		}

		req, err := inst.NewRequest(tt.method, "/guest-book/sign", strings.NewReader(val.Encode()))
		if err != nil {
			t.Error(err)
		}

		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		if tt.user != nil {
			aetest.Login(tt.user, req)
		}

		resp := httptest.NewRecorder()
		// exec request
		GuestSign(resp, req)

		if resp.Code != tt.code {
			t.Errorf("Got response code %d; want %d; body:\n%s", resp.Code, tt.code, resp.Body.String())
		}

		// Check the latest greeting against our expectation.
		c := appengine.NewContext(req)
		q := datastore.NewQuery("Greeting").Ancestor(models.GuestBookKey(c)).Order("-Date").Limit(1)
		var g models.Greeting
		_, err = q.Run(c).Next(&g)
		if err == datastore.Done {
			if tt.greeting != nil {
				t.Errorf("No greeting stored. Expected %v", tt.greeting)
			}
			continue
		}
		if err != nil {
			t.Errorf("Failed to fetch greeting: %v", err)
		}
		if tt.greeting == nil {
			if !g.Date.Before(now) {
				t.Errorf("Expected no new greeting, found: %v", g)
			}
			continue
		}

		if g.Date.Before(now) {
			t.Errorf("Greeting stored at %v, want at least %v", g.Date, now)
		}

		g.Date = time.Time{}
		if !reflect.DeepEqual(g, *tt.greeting) {
			t.Errorf("Greetings don't match. \nGot: %v\nWant: %v", g, *tt.greeting)
		}
	}
}
