package handlertransport_test

import (
	"io"
	"net/http"
	"os"

	"github.com/sunfmin/handlertransport"
)

/*
Here is an example, this can be used with http client library like https://github.com/machinebox/graphql to test a graphql handler.
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
