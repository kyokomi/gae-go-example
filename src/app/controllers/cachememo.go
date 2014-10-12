package controllers

import (
	"net/http"
	"appengine"
	"appengine/memcache"
	"github.com/gorilla/mux"
	"fmt"
	"bytes"
	"time"
)

func MemoAdd(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	key := r.PostFormValue("key")
	value := r.PostFormValue("value")
	c.Debugf("key = %s, value = %s", key, value)

	if key == "" || value == "" {
		http.Error(w, "key or value param error", http.StatusInternalServerError)
		return
	}

	// 有効期限1分
	item := &memcache.Item{
		Key:   key,
		Value: bytes.NewBufferString(value).Bytes(),
		Expiration: (time.Minute * 1),
	}
	if err := memcache.Add(c, item); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func MemoShow(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	vars := mux.Vars(r)
	key := vars["key"]
	if key == "" {
		http.Error(w, "key param error", http.StatusInternalServerError)
	}

	if cache, err := memcache.Get(c, key); err == memcache.ErrCacheMiss {
		// データが存在しない
		c.Debugf("キャッシュなし")
		fmt.Fprintf(w, "key = %s キャッシュなし", key)

	} else if err != nil {
		// error
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		c.Debugf("キャッシュあり")
		fmt.Fprintf(w, "key = %s キャッシュあり %s", key, string(cache.Value))
	}
}

func MemoRemove(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	vars := mux.Vars(r)
	key := vars["key"]
	if key == "" {
		http.Error(w, "key param error", http.StatusInternalServerError)
	}

	if err := memcache.Delete(c, key); err == memcache.ErrCacheMiss {
		c.Debugf("キャッシュなし")
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else if err != nil {
		// error
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		// delete
		c.Debugf("キャッシュ削除")
		fmt.Fprintf(w, "key = %s キャッシュ削除", key)
	}
}
