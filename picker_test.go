package crawler

import (
	"reflect"
	"testing"
)

func TestAnchorPicker(t *testing.T) {
	anchor := &AnchorPicker{}

	urls, err := anchor.Pick("<a href='http://ddo.me'>test</a><a href='http://ddict.me'>test</a>")

	if err != nil {
		t.Fail()
		return
	}

	if !reflect.DeepEqual(urls, []string{"http://ddo.me", "http://ddict.me"}) {
		t.Fail()
	}
}
