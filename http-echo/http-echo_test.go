package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestHandler(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler(w, r)
	}))
	defer ts.Close()

	values := make(url.Values)
	values.Set("buf", "foo")
	values.Set("anotherVariable", "bar")

	r, _ := http.Post(ts.URL+"/?"+values.Encode(), "", nil)
	body, _ := ioutil.ReadAll(r.Body)
	r.Body.Close()

	if string(body) != "CLOUDWALK foo bar" {
		t.Errorf("got %s want CLOUDWALK foo bar", body)
	}
}
