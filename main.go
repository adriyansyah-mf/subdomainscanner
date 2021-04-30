package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"gopkg.in/alecthomas/kingpin.v2"
)

func getdata() {

	var argUrl = kingpin.Arg("url", "Target Url").Required().String()

	kingpin.Parse()
	resp, err := http.Get("http://api.hackertarget.com/hostsearch/?q=" + *argUrl)
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp)
}

func main() {
	getdata()
}
