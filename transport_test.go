package handlertransport_test

import (
	"io"
	"net/http"
	"os"

	"github.com/sunfmin/handlertransport"
)

/*
	Provided to http.Client, But won't call a remote service.
*/
func ExampleNew_01Client() {
	hf := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world"))
	}

	client := &http.Client{
		Transport: handlertransport.New(http.HandlerFunc(hf)),
	}

	resp, err := client.Get("/")
	if err != nil {
		panic(err)
	}
	io.Copy(os.Stdout, resp.Body)
	// output:
	// Hello world
}
