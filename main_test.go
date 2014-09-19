package gaehoge

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"appengine/aetest"
	"fmt"
	"io/ioutil"
	"regexp"
)

func TestIndex(t *testing.T) {

	// test init
	c, err := aetest.NewContext(nil)
	if err != nil {
		t.Error(err)
	}
	defer c.Close()

	// test exec
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Error(err)
	}
	res := httptest.NewRecorder()

	index(res, req, c)

	// response check
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
	}
	h := string(data)
	fmt.Println(h)

	re := regexp.MustCompile("Sign Guestbook")
	if matched := re.MatchString(h); !matched {
		t.Error("unmatched")
	}
}
