package controllers

import (
	"fmt"
	"net/http"

	"appengine"
	"appengine/delay"
)

func DelayFunc(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	testFunc.Call(c, FuncArg{"test", "fuga"})

	fmt.Fprintf(w, "ok")
}

type FuncArg struct {
	X, Y string
}

var testFunc = delay.Func("hoge", func(c appengine.Context, arg FuncArg) {
	c.Debugf("delayFunc %s %s", arg.X, arg.Y)
})
