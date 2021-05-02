package main

import (
	"fmt"
	"net/http"
	"strconv"
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
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	for _, entry := range table {
		char, err := strconv.Unquote(`"\u` + entry.CodePoint + `"`)
		if err != nil {
			fmt.Fprintf(w, "failed to unquote %s\n", entry.CodePoint)
			continue
		}
		fmt.Fprintf(w, "%s: %s\n", char, entry.CharacterName)
	}
}
