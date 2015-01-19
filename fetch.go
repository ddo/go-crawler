package crawler

import (
	"io/ioutil"
	"net/http"
)

//get html via http.get
func get(url string) (string, error) {
	res, err := http.Get(url)

	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	html, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return "", err
	}

	return string(html), nil
}

func Fetch(url string, f Picker) ([]string, error) {
	html, err := get(url)

	if err != nil {
		return nil, err
	}

	urls, err := f.Picker(html)

	if err != nil {
		return nil, err
	}

	return urls, nil
}
