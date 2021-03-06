package webportal

import (
	"fmt"
	"net/http"
)

// RunWebPortal starts running the dino web portal on address addr
func RunWebPortal(addr string) error  {
	http.HandleFunc("/", rootHandler)
	return 	 http.ListenAndServe(addr,nil)

}

func rootHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w,"Hello from root handler web portal %s", r.RemoteAddr)
}