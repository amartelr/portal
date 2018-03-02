package webportal

import (
	"fmt"
	"net/http"
)

// Run pone servidor web en marcha
func Run(addr string) error {
	http.HandleFunc("/", roothandler)
	return http.ListenAndServe(addr, nil)
}

func roothandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Test web portal %s", r.RemoteAddr)
}
