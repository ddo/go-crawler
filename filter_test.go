package crawler

import (
	"testing"
)

func TestUniqueFilter(t *testing.T) {
	f := &UniqueFilter{}

	if !f.Filter("http://ddict.me") {
		t.Error("UniqueFilter init url")
	}

	if f.Filter("http://ddict.me") {
		t.Error("UniqueFilter duplicated url")
	}

	if !f.Filter("http://ddo.me") {
		t.Error("UniqueFilter new url")
	}
}

func TestSameHostFilter(t *testing.T) {
	f := NewSameHostFilter("http://ddict.me")

	if f.Filter("http://ddict.me") {
		t.Error("SameHostFilter same url")
	}

	if f.Filter("http://ddict.me/haha") {
		t.Error("SameHostFilter same host")
	}

	if !f.Filter("http://ddo.me") {
		t.Error("SameHostFilter diff host")
	}

	if !f.Filter("http://dashboard.ddict.me") {
		t.Error("SameHostFilter sub domain")
	}
}
