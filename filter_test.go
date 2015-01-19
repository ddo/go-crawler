package crawler

import (
	"net/url"
	"testing"
)

func TestUrlFilter(t *testing.T) {
	f := &UrlFilter{}

	if !f.Filter(parseUrl("http://ddict.me")) {
		t.Error("UniqueFilter normal url")
	}

	if !f.Filter(parseUrl("http://ddict.me/haha")) {
		t.Error("UniqueFilter normal url")
	}

	if !f.Filter(parseUrl("/")) {
		t.Error("UniqueFilter normal url")
	}

	if f.Filter(parseUrl("ftp://ddo.me")) {
		t.Error("UniqueFilter ftp url")
	}

	if f.Filter(parseUrl("javascript:void()")) {
		t.Error("UniqueFilter javascript url")
	}
}

func TestUniqueFilter(t *testing.T) {
	f := &UniqueFilter{}

	if !f.Filter(parseUrl("http://ddict.me")) {
		t.Error("UniqueFilter init url")
	}

	if f.Filter(parseUrl("http://ddict.me")) {
		t.Error("UniqueFilter duplicated url")
	}

	if !f.Filter(parseUrl("http://ddo.me")) {
		t.Error("UniqueFilter new url")
	}

	if f.Filter(parseUrl("http://ddo.me#hello")) {
		t.Error("UniqueFilter ignore fragment")
	}
}

func TestSameHostFilter(t *testing.T) {
	f := &SameHostFilter{
		parseUrl("http://ddict.me"),
	}

	if !f.Filter(parseUrl("http://ddict.me")) {
		t.Error("SameHostFilter same url")
	}

	if !f.Filter(parseUrl("http://ddict.me/haha")) {
		t.Error("SameHostFilter same host")
	}

	if f.Filter(parseUrl("http://ddo.me")) {
		t.Error("SameHostFilter diff host")
	}

	if f.Filter(parseUrl("http://dashboard.ddict.me")) {
		t.Error("SameHostFilter sub domain")
	}
}

func parseUrl(u string) *url.URL {
	u_obj, _ := url.Parse(u)

	return u_obj
}
