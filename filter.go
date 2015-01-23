package crawler

import (
	"net/url"
)

type Filter interface {
	Filter(*url.URL) bool
}

//is http url Filter
type FilterUrl struct{}

func (f *FilterUrl) Filter(u *url.URL) bool {
	return u.Scheme == "" || u.Scheme == "http" || u.Scheme == "https"
}

//no duplicated Filter
type FilterUnique struct {
	Us []*url.URL
}

func (f *FilterUnique) Filter(u *url.URL) bool {
	for _, old_u := range f.Us {
		if *u == *old_u {
			return false
		}

		another_u := *u
		another_old_u := *old_u

		another_u.Fragment = ""
		another_old_u.Fragment = ""

		if another_u == another_old_u {
			return false
		}
	}

	f.Us = append(f.Us, u)

	return true
}

//same hostname only
type FilterSameHost struct {
	U *url.URL
}

func (f *FilterSameHost) Filter(u *url.URL) bool {
	return f.U.Host == u.Host
}
