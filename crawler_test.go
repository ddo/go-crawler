package crawler

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	c, _ := New(&Config{
		Url: "http://facebook.com/",
	})

	if reflect.TypeOf(*c).String() != "crawler.Crawler" {
		t.Error("New should return *crawler.Crawler")
		return
	}
}

func TestNewFailure(t *testing.T) {
	_, err := New(&Config{
		Url: ":",
	})

	if err == nil {
		t.Error("New should return error on invalid url")
		return
	}
}

func TestStart(t *testing.T) {
	urls, errs := []string{}, []error{}

	c, _ := New(&Config{
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

func TestStartFailure(t *testing.T) {
	urls, errs := []string{}, []error{}

	c, _ := New(&Config{
		Url: "http://123456.com/",
	})

	c.Start(func(url string) {
		urls = append(urls, url)
	}, func(err error) {
		errs = append(errs, err)
	})

	if len(urls) > 0 {
		t.Error("TestStartFailure urls should be 0")
		return
	}

	if len(errs) != 1 {
		t.Error("TestStartFailure errs be 1")
		return
	}
}
