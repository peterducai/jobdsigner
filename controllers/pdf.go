package controllers

import (
	"net/http"
)

var streamPDFbytes []byte

//PDF sends pdf
func PDF(w http.ResponseWriter, r *http.Request) {
	//
	// streamPDFbytes, err := ioutil.ReadFile("./receipt.pdf")
	//
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
	//
	// b := bytes.NewBuffer(streamPDFbyte)
	//
	// // stream straight to client(browser)
	// w.Header().Set("Content-type", "application/pdf")
	// if _, err := b.WriteTo(w); err != nil { // <----- here!
	// 	fmt.Fprintf(w, "%s", err)
	// }
	// w.Write([]byte("PDF Generated"))
}
