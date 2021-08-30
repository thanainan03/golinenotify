package main

import (
	"fmt"
	"golinenotify/golinenotify"
	"log"
)

func main() {
	url := golinenotify.GetAuthorizeUrl("S36gTcAhUTXAHFBQ2G23t1", "http://localhost:3333/line-notify/callback", "444")
	fmt.Println(url)

	a, err := golinenotify.GetAccessToken("S36gTcAhUTXAHFBQ2G23t1", "TLW3cBjdohAHuYLw3tnCL188IX1xTsSHUNE3aZalY3x", "http://localhost:3333/line-notify/callback", "tqKQNNfIT2XkZWpJfSae4m")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(a)

	err2 := golinenotify.Send("bCjP5KfzRF48sdpOkzx9WKyfkDtMR05vG4Z5Oor5NeH", "test")
	if err != nil {
		log.Fatal(err2)
	}
}
