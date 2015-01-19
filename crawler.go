package crawler

import (
	"net/url"
)

type Crawler struct {
	limit int
}

func New(limit int) *Crawler {
	return &Crawler{
		limit: limit,
	}
}

func (c *Crawler) Start(u string, ch_url chan string, ch_err chan error, ch_done chan bool) {
	u_obj, _ := url.Parse(u)

	//anchor picker
	picker_anchor := &AnchorPicker{}

	//unique filter
	filter_url := &UrlFilter{}

	//unique filter
	filter_unique := &UniqueFilter{
		[]*url.URL{u_obj},
	}

	//same host filter
	filter_samehost := &SameHostFilter{
		u_obj,
	}

	c.crawl(u, []Picker{picker_anchor}, []Filter{filter_url, filter_unique, filter_samehost}, ch_url, ch_err, ch_done)
}

func (c *Crawler) crawl(u string, pickers []Picker, filters []Filter, ch_url chan string, ch_err chan error, ch_done chan bool) {

	urls, err := Fetch(u, pickers[0]) //hardcode

	if err != nil {
		ch_err <- err
		return
	}

parent_loop:
	for _, child_u_obj := range urls {
		for _, f := range filters {
			if !f.Filter(child_u_obj) {
				continue parent_loop
			}
		}

		//valid url
		c.limit--

		select {
		case <-ch_done:
			return
		default:
		}

		if c.limit < 0 {
			close(ch_done)
			return
		}

		child_url := child_u_obj.String()

		ch_url <- child_url

		go c.crawl(child_url, pickers, filters, ch_url, ch_err, ch_done)
	}
}
