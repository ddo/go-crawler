package crawler

import (
	"testing"

	"."
)

func TestFetch(t *testing.T) {
	urls, err := crawler.Fetch("http://ddict.me", &crawler.AnchorFilter{})

	if err != nil {
		t.Fail()
		return
	}

	if len(urls) < 1 {
		t.Fail()
		return
	}
}
