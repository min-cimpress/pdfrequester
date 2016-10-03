package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	url    string
	output string
)

func panicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}

func init() {
	flag.StringVar(&url, "url", "", "url of pdf file")
	flag.StringVar(&output, "output", "", "pdf file output path")
}

func main() {
	flag.Parse()
	fmt.Println(url)
	res, err := http.Get(url)
	panicOnErr(err)

	pdf, err := ioutil.ReadAll(res.Body)
	panicOnErr(err)
	res.Body.Close()
	ioutil.WriteFile(output, pdf, 0644)
}
