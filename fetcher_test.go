package crawler

import (
	"net/http"
	"net/url"
	"testing"
	"time"
)

func TestFetch(t *testing.T) {
	f := Fetcher{
		&http.Client{
			Timeout: time.Second * 10,
		},
	}

	u, _ := url.Parse("http://ddict.me")

	urls, err := f.Fetch(u, &AnchorPicker{})

	if err != nil {
		t.Fail()
		return
	}

	if len(urls) < 1 {
		t.Fail()
		return
	}
}
