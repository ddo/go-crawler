package crawler

import (
	"net/url"
)

type Filter interface {
	Filter(string) bool
}

//default Filter
type UniqueFilter struct {
	urls []*url.URL
}

func (f *UniqueFilter) Filter(new_url string) bool {
	u, err := parseUrl(new_url)

	if err != nil {
		return false
	}

	for _, old_u := range f.urls {
		if *u == *old_u {
			return false
		}
	}

	f.urls = append(f.urls, u)

	return true
}

//same hostname only
type SameHostFilter struct {
	root_url *url.URL
}

func NewSameHostFilter(url string) *SameHostFilter {
	root_url, _ := parseUrl(url)

	return &SameHostFilter{root_url}
}

func (f *SameHostFilter) Filter(new_url string) bool {
	u, err := parseUrl(new_url)

	if err != nil {
		return false
	}

	return f.root_url.Host != u.Host
}

//parse url helper
func parseUrl(new_url string) (*url.URL, error) {
	u, err := url.Parse(new_url)

	if err != nil {
		return nil, err
	}

	if u.Path == "" {
		u.Path = "/"
	}

	return u, nil
}
