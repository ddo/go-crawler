package crawler

import (
	"net/url"
	"testing"
)

func TestFilterUrl(t *testing.T) {
	f := &FilterUrl{}

	if !f.Filter(parseUrl("http://facebook.com")) {
		t.Error("FilterUrl normal url")
	}

	if !f.Filter(parseUrl("http://facebook.com/haha")) {
		t.Error("FilterUrl normal url")
	}

	if !f.Filter(parseUrl("/")) {
		t.Error("FilterUrl normal url")
	}

	if f.Filter(parseUrl("ftp://ddo.me")) {
		t.Error("FilterUrl ftp url")
	}

	if f.Filter(parseUrl("javascript:void()")) {
		t.Error("FilterUrl javascript url")
	}
}

func TestFilterUnique(t *testing.T) {
	f := &FilterUnique{}

	if !f.Filter(parseUrl("http://facebook.com")) {
		t.Error("FilterUnique init url")
	}

	if f.Filter(parseUrl("http://facebook.com")) {
		t.Error("FilterUnique duplicated url")
	}

	if !f.Filter(parseUrl("http://ddo.me")) {
		t.Error("FilterUnique new url")
	}

	if f.Filter(parseUrl("http://ddo.me#hello")) {
		t.Error("FilterUnique ignore fragment")
	}
}

func TestFilterSameHost(t *testing.T) {
	f := &FilterSameHost{
		parseUrl("http://facebook.com"),
	}

	if !f.Filter(parseUrl("http://facebook.com")) {
		t.Error("FilterSameHost same url")
	}

	if !f.Filter(parseUrl("http://facebook.com/haha")) {
		t.Error("FilterSameHost same host")
	}

	if f.Filter(parseUrl("http://ddo.me")) {
		t.Error("FilterSameHost diff host")
	}

	if f.Filter(parseUrl("http://apps.facebook.com")) {
		t.Error("FilterSameHost sub domain")
	}
}

func parseUrl(u string) *url.URL {
	u_obj, _ := url.Parse(u)

	return u_obj
}
