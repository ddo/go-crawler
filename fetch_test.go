package crawler

import (
	"testing"
)

func TestFetch(t *testing.T) {
	urls, err := Fetch("http://ddict.me", &AnchorPicker{})

	if err != nil {
		t.Fail()
		return
	}

	if len(urls) < 1 {
		t.Fail()
		return
	}
}
