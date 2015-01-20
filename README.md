# go-crawler
crawler in go

> Project is not ready yet

## Quick Glance

```go
package main

import (
    "fmt"
    "github.com/ddo/go-crawler"
)

func main() {
    no := 0

    c := crawler.New("http://facebook.com/", 10) //limit 10

    receiver_url := func(url string) {
        no++
        fmt.Println(no, "\t ", url)
    }

    receiver_err := func(err error) {
        fmt.Println("error\t", err)
    }

    c.Start(receiver_url, receiver_err)

    fmt.Println("done")
}
```

output

```shell
1     http://facebook.com/legal/terms
2     http://facebook.com/about/privacy
3     http://facebook.com/help/cookies
4     http://facebook.com/pages/create/?ref_type=registration_form
5     http://facebook.com/r.php
6     http://facebook.com/login/
7     http://facebook.com/mobile/?ref=pf
8     http://facebook.com/find-friends?ref=pf
9     http://facebook.com/badges/?ref=pf
10    http://facebook.com/directory/people/
done
```

## Todo

* [ ] #Use filter
* [ ] timeout
* [ ] delay
* [ ] crawler testing
* [ ] travis-ci
* [ ] coveralls.io
* [ ] README

