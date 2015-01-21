package main

import (
	"fmt"
	"net/url"

	"github.com/ddo/go-crawler"
)

func main() {
	url_str := "http://facebook.com/"

	//counter, just for better log
	no := 0

	////////////////////define no external urls filter/////////////////
	u, _ := url.Parse(url_str)

	//url filter, http url only (no ftp etc)
	filter_url := &crawler.UrlFilter{}

	//unique filter, no duplicated url
	filter_unique := &crawler.UniqueFilter{
		[]*url.URL{u},
	}

	//same host filter, NO EXTERNAL links
	filter_samehost := &crawler.SameHostFilter{
		u,
	}

	filters := []crawler.Filter{filter_url, filter_unique, filter_samehost}
	////////////////////define no external urls filter - end/////////////////

	c := crawler.New(&crawler.Config{
		Url:     url_str,
		Filters: filters,
	})

	//url handler
	receiver_url := func(url string) {
		no++
		fmt.Println(no, "\t ", url)
	}

	//err handler
	receiver_err := func(err error) {
		fmt.Println("error\t", err)
	}

	//trigger
	c.Start(receiver_url, receiver_err)

	fmt.Println("done thanks god")
}

/*
	OUTPUT:
	1 	  http://facebook.com/legal/terms
	2 	  http://facebook.com/about/privacy
	3 	  http://facebook.com/help/cookies
	4 	  http://facebook.com/pages/create/?ref_type=registration_form
	5 	  http://facebook.com/r.php
	6 	  http://facebook.com/login/
	7 	  http://facebook.com/mobile/?ref=pf
	8 	  http://facebook.com/find-friends?ref=pf
	9 	  http://facebook.com/badges/?ref=pf
	10 	  http://facebook.com/directory/people/
	done thanks god
*/
