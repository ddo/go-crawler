package crawler

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Picker interface {
	Pick(string) ([]string, error)
}

//default picker
type AnchorPicker struct{}

func (p *AnchorPicker) Pick(html string) (urls []string, err error) {
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
