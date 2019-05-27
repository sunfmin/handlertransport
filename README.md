

Provide to `http.Client{Transport: handlertransport.New(my.HelloHandler)}` to call to the http.Handler instead of remote network http call.



	Provided to http.Client, But won't call a remote service.
```go
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
```



