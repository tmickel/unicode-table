package main

import (
	"encoding/csv"
	"io"
	"net/http"
)

type Entry struct {
	CodePoint                     string
	CharacterName                 string
	GeneralCategory               string
	CanonicalCombiningClasses     string
	BidirectionalCategory         string
	CharacterDecompositionMapping string
	DecimalDigitValue             string
	DigitValue                    string
	NumericValue                  string
	Mirrored                      string
	Unicode1Name                  string
	ISO10646CommentField          string
	UppercaseMapping              string
	LowercaseMapping              string
	TitlecaseMapping              string
}

func FetchTable() ([]*Entry, error) {
	resp, err := http.Get("https://unicode.org/Public/UNIDATA/UnicodeData.txt")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	r := csv.NewReader(resp.Body)

	entries := []*Entry{}

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		entry := &Entry{
			CodePoint: record[0],
		}
		entries = append(entries, entry)
	}
	return entries, nil
}
