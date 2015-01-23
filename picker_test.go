package crawler

import (
	"reflect"
	"strings"
	"testing"
)

func TestPickerAttr(t *testing.T) {
	anchor := &PickerAttr{
		TagName: "a",
		Attr:    "href",
	}

	a, err := anchor.Pick(strings.NewReader("<a href='http://ddo.me'>test</a><a href='http://ddict.me'>test</a>"))

	if err != nil {
		t.Fail()
		return
	}

	if !reflect.DeepEqual(a, []string{"http://ddo.me", "http://ddict.me"}) {
		t.Fail()
	}

	script := &PickerAttr{
		TagName: "script",
		Attr:    "src",
	}

	s, err := script.Pick(strings.NewReader("<script src='haha'></script>"))

	if err != nil {
		t.Fail()
		return
	}

	if !reflect.DeepEqual(s, []string{"haha"}) {
		t.Fail()
	}
}
