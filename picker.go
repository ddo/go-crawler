package crawler

import (
	"io"

	"golang.org/x/net/html"
)

type Picker interface {
	Pick(r io.Reader) ([]string, error)
}

//default picker
type PickerAttr struct {
	TagName string
	Attr    string
}

func (p *PickerAttr) Pick(r io.Reader) (data []string, err error) {
	z := html.NewTokenizer(r)

	for {
		tt := z.Next()

		switch tt {
		case html.ErrorToken:
			if z.Err() == io.EOF {
				return data, nil
			}

		case html.StartTagToken:
			tag_name, attr := z.TagName()

			if string(tag_name) != p.TagName {
				continue
			}

			var key, value []byte

			for attr {
				key, value, attr = z.TagAttr()

				if string(key) == p.Attr {
					data = append(data, string(value))
				}
			}
		}
	}

	return data, z.Err()
}
