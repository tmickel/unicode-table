package main

import (
	"io/ioutil"
	"net/http"
)

func FetchTable() (string, error) {
	resp, err := http.Get("https://unicode.org/Public/UNIDATA/UnicodeData.txt")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return "", err
    }

	return string(bodyBytes), nil
}
