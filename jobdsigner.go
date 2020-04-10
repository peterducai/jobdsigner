package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"

	"github.com/peterducai/jobdsigner/models"
)

func main() {
	ver := models.Version{0, 0, 1, "xeee1"}
	fmt.Printf("Job dSigner %d.%d.%d %s\n", ver.MAJOR,ver.MINOR,ver.PATCH,ver.HASH)
	
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
		fmt.Fprintf(w, "Job dSigner %d.%d.%d %s\n", ver.MAJOR,ver.MINOR,ver.PATCH,ver.HASH)
	})
	cfg := &tls.Config{
		MinVersion:               tls.VersionTLS12,
		CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
		PreferServerCipherSuites: true,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
			tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_RSA_WITH_AES_256_CBC_SHA,
		},
	}
	srv := &http.Server{
		Addr:         ":8443",
		Handler:      mux,
		TLSConfig:    cfg,
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
	}
	fmt.Println("server is running at https://localhost:8443")
	log.Fatal(srv.ListenAndServeTLS("tls.crt", "tls.key"))
	
}
