package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", DisplayTable)
	http.ListenAndServe(":80", nil)
}

func DisplayTable(w http.ResponseWriter, r *http.Request) {
	table, err := FetchTable()
	if err != nil {
		fmt.Fprintf(w, "could not fetch current unicode table: %v", err)
	}
	for _, entry := range table {
		fmt.Fprintf(w, "%s\n", entry.CodePoint)
	}
}
