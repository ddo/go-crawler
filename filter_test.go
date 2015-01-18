package crawler

import (
	"reflect"
	"testing"

	"."
)

func TestAnchorFilter(t *testing.T) {
	anchor := &crawler.AnchorFilter{}

	urls, err := anchor.Filter("<a href='http://ddo.me'>test</a><a href='http://ddict.me'>test</a>")

	if err != nil {
		t.Fail()
		return
	}

	if !reflect.DeepEqual(urls, []string{"http://ddo.me", "http://ddict.me"}) {
		t.Fail()
	}
}
