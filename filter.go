package crawler

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Filter interface {
	Filter(string) ([]string, error)
}

//default filter
type AnchorFilter struct{}

func (f *AnchorFilter) Filter(html string) (urls []string, err error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))

	if err != nil {
		return nil, err
	}

	doc.Find("a").Each(func(i int, a *goquery.Selection) {
		url, exist := a.Attr("href")

		if !exist {
			return
		}

		urls = append(urls, url)
	})

	return urls, nil
}
