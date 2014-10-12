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
	"app/controllers"
)

type handlerTest struct {
	in string
	handler func(http.ResponseWriter, *http.Request)
	out string
}

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

	controllers.Index(res, req)

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

	inst, err := aetest.NewInstance(nil)
	if err != nil {
		t.Error(err)
	}
	defer inst.Close()

	val := url.Values{}
	val.Set("content", "hogehoge")
	req, err := inst.NewRequest("POST", "/sign", strings.NewReader(val.Encode()))
	if err != nil {
		t.Error(err)
	}

	res := httptest.NewRecorder()

	controllers.GuestSign(res, req)

	if res.Header().Get("Location") != "/" {
		t.Error("bad response location URL ", res.Header().Get("Location"))
	}

	if http.StatusFound != res.Code {
		t.Error("bad response status code ", res.Code)
	}
}
