package solr

import "net/url"

// URLEncoded - This encodes the URL
func URLEncoded(str string) string {
	u, err := url.Parse(str)
	if err != nil {
		return str
	}
	return u.String()
}
