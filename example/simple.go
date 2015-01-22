package main

import (
	"fmt"
	"github.com/ddo/go-crawler"
)

func main() {
	//counter, just for better log
	no := 0

	/*
		default limit: 	10
		default client: timeout 10s
		default filter: http(s), no duplicated
		default scope: 	http(s), no duplicated, same host only
	*/
	c, err := crawler.New(&crawler.Config{
		Url: "http://facebook.com/",
	})

	//your url is invalid
	if err != nil {
		panic(err)
	}

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
	1 	  https://www.facebook.com/recover/initiate
	2 	  http://facebook.com/legal/terms
	3 	  http://facebook.com/about/privacy
	4 	  http://facebook.com/help/cookies
	5 	  http://facebook.com/pages/create/?ref_type=registration_form
	6 	  https://vi-vn.facebook.com/
	7 	  https://www.facebook.com/
	8 	  https://zh-tw.facebook.com/
	9 	  https://ko-kr.facebook.com/
	10 	  https://ja-jp.facebook.com/
	done
*/
