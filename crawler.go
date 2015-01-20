package crawler

import (
	"net/url"
)

type Crawler struct {
	u       string
	limit   int
	ch_url  chan string
	ch_err  chan error
	ch_done chan bool
	picker  Picker
	filters []Filter
}

func New(u string, limit int) *Crawler {
	u_obj, _ := url.Parse(u)

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

	return &Crawler{
		u:       u,
		limit:   limit,
		ch_url:  make(chan string),
		ch_err:  make(chan error),
		ch_done: make(chan bool),
		picker:  &AnchorPicker{},
		filters: []Filter{filter_url, filter_unique, filter_samehost},
	}
}

type receiver_url func(string)
type receiver_err func(error)

func (c *Crawler) Start(r_url receiver_url, r_err receiver_err) {
	go c.crawl(c.u)

loop:
	for {
		select {
		case url := <-c.ch_url:
			r_url(url)
		case err := <-c.ch_err:
			r_err(err)
		case <-c.ch_done:
			break loop
		}
	}

	return
}

func (c *Crawler) crawl(u string) {
	urls, err := Fetch(u, c.picker)

	if err != nil {
		c.ch_err <- err
		return
	}

parent_loop:
	for _, child_u_obj := range urls {
		for _, f := range c.filters {
			if !f.Filter(child_u_obj) {
				continue parent_loop
			}
		}

		//valid url
		c.limit--

		select {
		case <-c.ch_done:
			return
		default:
		}

		if c.limit < 0 {
			close(c.ch_done)
			return
		}

		child_url := child_u_obj.String()

		c.ch_url <- child_url

		go c.crawl(child_url)
	}
}
