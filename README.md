# go-crawler [![Build Status][travis-img]][travis-url] [![Build Status][coveralls-img]][coveralls-url]
just a crawler in go

[travis-img]: https://img.shields.io/travis/ddo/go-crawler.svg?style=flat-square
[travis-url]: https://travis-ci.org/ddo/go-crawler

[coveralls-img]: https://img.shields.io/coveralls/ddo/go-crawler.svg?style=flat-square
[coveralls-url]: https://coveralls.io/r/ddo/go-crawler

> configable - concurrency

## Quick Glance

```go
package main

import (
    "fmt"
    "github.com/ddo/go-crawler"
)

func main() {
    //counter, just for better log
    no := 0

    /*
        default limit:  10
        default client: timeout 10s
        default filter: http(s), no duplicated
        default scope:  http(s), no duplicated, same host only
    */
    c := crawler.New(&crawler.Config{
        Url: "http://facebook.com/",
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

    fmt.Println("done")
}
```

output

```shell
1     https://www.facebook.com/recover/initiate
2     http://facebook.com/legal/terms
3     http://facebook.com/about/privacy
4     http://facebook.com/help/cookies
5     http://facebook.com/pages/create/?ref_type=registration_form
6     https://vi-vn.facebook.com/
7     https://www.facebook.com/
8     https://zh-tw.facebook.com/
9     https://ko-kr.facebook.com/
10    https://ja-jp.facebook.com/
done
```

## Todo

* [x] init with Filter
* [x] init with http.Client
* [x] crawler testing
* [ ] init with Fetcher
* [ ] mutex/chan limit/worker counter
* [ ] delay
* [ ] travis-ci
* [ ] coveralls.io
* [ ] README advanced doc

