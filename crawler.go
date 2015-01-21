package crawler

import (
	"net/http"
	"net/url"
	"time"
)

type Crawler struct {
	u       *url.URL
	limit   int
	worker  int
	ch_url  chan string
	ch_err  chan error
	ch_done chan bool
	fetcher Fetcher
	filters []Filter
}

type Config struct {
	Url     string
	Limit   int
	Filters []Filter
	Client  *http.Client
}

func New(config *Config) *Crawler {
	u, _ := url.Parse(config.Url)

	limit := config.Limit
	filters := config.Filters
	client := config.Client

	//default limit - 50
	if limit == 0 {
		limit = 50
	}

	//default client - timeout 10s
	if client == nil {
		client = &http.Client{
			Timeout: time.Second * 10,
		}
	}

	fetcher := Fetcher{
		Client: client,
		Picker: &AnchorPicker{},
	}

	//default filters
	if filters == nil {
		//unique filter
		filter_url := &UrlFilter{}

		//unique filter
		filter_unique := &UniqueFilter{
			[]*url.URL{u},
		}

		//same host filter
		filter_samehost := &SameHostFilter{
			u,
		}

		filters = []Filter{filter_url, filter_unique, filter_samehost}
	}

	return &Crawler{
		u:       u,
		limit:   limit,
		ch_url:  make(chan string),
		ch_err:  make(chan error),
		ch_done: make(chan bool),
		fetcher: fetcher,
		filters: filters,
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

func (c *Crawler) crawl(u *url.URL) {
	c.worker++

	urls, err := c.fetcher.Fetch(u)

	if err != nil {
		c.ch_err <- err
	} else {

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
				println("done by limit")
				close(c.ch_done)
				return
			}

			c.ch_url <- child_u_obj.String()

			go c.crawl(child_u_obj)
		}
	}

	//need to add delay here ?
	c.worker--
	if c.worker == 0 {
		println("done by all worker done")
		close(c.ch_done)
	}
}
