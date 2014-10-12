package controllers

import (
	"io/ioutil"
	"net/http/httptest"
	"regexp"
	"testing"

	"appengine/aetest"
)

func TestIndex(t *testing.T) {

	inst, err := aetest.NewInstance(nil)
	if err != nil {
		t.Error(err)
	}
	defer inst.Close()

	req, err := inst.NewRequest("GET", "/", nil)
	if err != nil {
		t.Error(err)
	}

	res := httptest.NewRecorder()

	Index(res, req)

	// response check
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
	}
	h := string(data)
	//fmt.Println("html: ", h)

	re := regexp.MustCompile("Sign Guestbook")
	if matched := re.MatchString(h); !matched {
		t.Error("not matched")
	}
}
