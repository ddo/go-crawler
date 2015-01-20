package main

import (
	"fmt"
	"github.com/ddo/go-crawler"
)

func main() {
	no := 0

	c := crawler.New("http://talktv.vn/", 50)

	c.Start(func(url string) {
		no++
		fmt.Println(no, "\t ", url)
	}, func(err error) {
		fmt.Println("error\t", err)
	})

	fmt.Println("done thanks god")
}
