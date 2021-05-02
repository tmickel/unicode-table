package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":80", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	table, err := FetchTable()
	if err != nil {
		fmt.Fprintf(w, "could not fetch current unicode table: %v", err)
	}
	fmt.Fprintf(w, "%#v", table)
}
