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

func (f *UniqueFilter) Filter(u *url.URL) bool {
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

func (f *SameHostFilter) Filter(u *url.URL) bool {
	return f.root_url.Host != u.Host
}
