package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("ERROR:", err)
		os.Exit(1)
	}

	//fmt.Println(resp) //doesn't return body of get call

	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	//io.Copy(os.Stdout, resp.Body)

	lw := logWriter{}
	io.Copy(lw, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	//return 1, nil //garbage writer
	fmt.Println(string(bs))
	fmt.Println("just wrote this many bytes:", len(bs))
	return len(bs), nil
}
