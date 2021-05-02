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
	r.Comma = ';'

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
			CodePoint:                     record[0],
			CharacterName:                 record[1],
			GeneralCategory:               record[2],
			CanonicalCombiningClasses:     record[3],
			BidirectionalCategory:         record[4],
			CharacterDecompositionMapping: record[5],
			DecimalDigitValue:             record[6],
			DigitValue:                    record[7],
			NumericValue:                  record[8],
			Mirrored:                      record[9],
			Unicode1Name:                  record[10],
			ISO10646CommentField:          record[11],
			UppercaseMapping:              record[12],
			LowercaseMapping:              record[13],
			TitlecaseMapping:              record[14],
		}
		entries = append(entries, entry)
	}
	return entries, nil
}
