package controllers

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var streamSVGbytes []byte

func check(e error) {
	if e != nil {
		panic(e)
	}
}

//WriteGraph writes graph to svg string
func WriteGraph() {
	f, err := os.Create("/tmp/dat2")
	check(err)
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n4)
}

//SVG sends svg to http client
func SVG(w http.ResponseWriter, r *http.Request) {

	streamSVGbytes, err := ioutil.ReadFile("./receipt.svg")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	b := bytes.NewBuffer(streamSVGbytes)

	// stream straight to client(browser)
	w.Header().Set("Content-type", "application/pdf")
	if _, err := b.WriteTo(w); err != nil { // <----- here!
		fmt.Fprintf(w, "%s", err)
	}
	w.Write([]byte("PDF Generated"))
}
