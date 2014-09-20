package gaehoge

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"net/url"
	"strings"

	"appengine/aetest"
)

func TestIndex(t *testing.T) {

	// test init
	c, err := aetest.NewContext(nil)
	if err != nil {
		t.Error(err)
	}
	defer c.Close()

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
	fmt.Println("html: ", h)

	re := regexp.MustCompile("Sign Guestbook")
	if matched := re.MatchString(h); !matched {
		t.Error("not matched")
	}
}

func TestSign(t *testing.T) {

	// test init
	c, err := aetest.NewContext(nil)
	if err != nil {
		t.Error(err)
	}
	defer c.Close()

	val := url.Values{}
	val.Set("content", "hogehoge")
	req, err := http.NewRequest("POST", "/sign", strings.NewReader(val.Encode()))
	if err != nil {
		t.Error(err)
	}
	res := httptest.NewRecorder()

	sign(res, req, c)

	if res.Header().Get("Location") != "/" {
		t.Error("bad response location URL ", res.Header().Get("Location"))
	}

	if http.StatusFound != res.Code {
		t.Error("bad response status code ", res.Code)
	}
}
