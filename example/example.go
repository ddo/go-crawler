package main

import "github.com/ddo/go-crawler"

func main() {
	no := 0

	c := crawler.New(50)

	ch_url := make(chan string)
	ch_err := make(chan error)
	ch_done := make(chan bool)

	go c.Start("http://facebook.com", ch_url, ch_err, ch_done)

loop:
	for {
		select {
		case url := <-ch_url:
			no++
			println(no, "\t ", url)
		case err := <-ch_err:
			println("error ", err)
		case <-ch_done:
			println("done thanks god")
			break loop
		}
	}

	println("done thanks god again")
}
