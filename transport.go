/*
Provide to `http.Client{Transport: handlertransport.New(my.HelloHandler)}` to call to the http.Handler instead of remote network http call.
*/
package handlertransport

import (
	"net/http"
	"net/http/httptest"
)

type ht struct {
	h http.Handler
}

func New(h http.Handler) http.RoundTripper {
	return &ht{h: h}
}

func (t *ht) RoundTrip(req *http.Request) (res *http.Response, err error) {
	rec := httptest.NewRecorder()
	t.h.ServeHTTP(rec, req)
	res = rec.Result()
	return
}
