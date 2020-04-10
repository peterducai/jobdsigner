package api

import (
	"fmt"
	"net/http"
)

//About adds job to pipeline
func About(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
	fmt.Fprintf(w, "Job dSigner 0.1") //%d.%d.%d %s\n", ver.MAJOR, ver.MINOR, ver.PATCH, ver.HASH)
}
