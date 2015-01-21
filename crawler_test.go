package crawler

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	c := New(&Config{
		Url: "http://facebook.com/",
	})

	if reflect.TypeOf(*c).String() != "crawler.Crawler" {
		t.Fail()
		return
	}
}

func TestStart(t *testing.T) {
	urls, errs := []string{}, []error{}

	c := New(&Config{
		Url: "http://facebook.com/",
	})

	c.Start(func(url string) {
		urls = append(urls, url)
	}, func(err error) {
		errs = append(errs, err)
	})

	if len(urls) != 10 {
		t.Error("TestStart limit should be 10")
		return
	}

	if len(errs) != 0 {
		t.Error("TestStart should be no error")
		return
	}
}
