package crawler

import (
	"net/http"
	"net/url"
	"testing"
	"time"
)

func TestFetch(t *testing.T) {
	anchor_picker := &PickerAttr{
		TagName: "a",
		Attr:    "href",
	}

	f := Fetcher{
		Client: &http.Client{
			Timeout: time.Second * 10,
		},
		Picker: anchor_picker,
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
