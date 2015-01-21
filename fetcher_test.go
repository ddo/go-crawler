package crawler

import (
	"net/http"
	"net/url"
	"testing"
	"time"
)

func TestFetch(t *testing.T) {
	f := Fetcher{
		Client: &http.Client{
			Timeout: time.Second * 10,
		},
		Picker: &AnchorPicker{},
	}

	u, _ := url.Parse("http://facebook.com")

	urls, err := f.Fetch(u)

	if err != nil {
		t.Fail()
		return
	}

	if len(urls) < 1 {
		t.Fail()
		return
	}
}
