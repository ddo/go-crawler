package crawler

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

type Fetcher struct {
	Client *http.Client
	Picker Picker
}

func (f *Fetcher) get(u *url.URL) (string, error) {
	res, err := f.Client.Get(u.String())

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

func (f *Fetcher) Fetch(u *url.URL) (urls_obj []*url.URL, err error) {
	html, err := f.get(u)

	if err != nil {
		return nil, err
	}

	urls, err := f.Picker.Pick(html)

	if err != nil {
		return nil, err
	}

	//update non host url
	for _, v := range urls {
		u_child, _ := url.Parse(v)

		urls_obj = append(urls_obj, u.ResolveReference(u_child))
	}

	return urls_obj, nil
}
