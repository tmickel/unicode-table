package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
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
	fmt.Fprint(w, "Tim's Complete Unicode Table, fetched live from https://unicode.org/Public/UNIDATA/UnicodeData.txt\n\n")
	for _, entry := range table {
		encoded := fmt.Sprintf("%08s", entry.CodePoint)
		char, err := strconv.Unquote(`"\U` + encoded + `"`)
		if err != nil {
			fmt.Fprintf(w, "failed to unquote %s\n", entry.CodePoint)
			continue
		}
		charName := strings.Title(strings.ToLower(entry.CharacterName))
		oldName := strings.Title(strings.ToLower(entry.Unicode1Name))
		fmt.Fprintf(w, "%s | %s | %s | %s\n", entry.CodePoint, char, charName, oldName)
	}
}
