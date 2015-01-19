package crawler

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

//get html via http.get
func get(u string) (string, error) {
	res, err := http.Get(u)

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

func Fetch(root_url string, f Picker) (urls_obj []*url.URL, err error) {
	html, err := get(root_url)

	if err != nil {
		return nil, err
	}

	//root url
	root_url_obj, _ := url.Parse(root_url)

	urls, err := f.Picker(html)

	if err != nil {
		return nil, err
	}

	//update non host url
	for _, v := range urls {
		url_obj, _ := url.Parse(v)

		urls_obj = append(urls_obj, root_url_obj.ResolveReference(url_obj))
	}

	return urls_obj, nil
}
