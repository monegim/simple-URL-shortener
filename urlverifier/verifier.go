package urlverifier

import (
	"net/url"
)

type VerifyData struct {
	rawURL                    string
	urlComponents          *url.URL
	result                 *Result
	httpCheckEnabled       bool
	allowHttpCheckInterval bool
}

type Result struct {
	IsURL        bool
	IsRFC3986URI bool
	HTTP         *HTTP
}

func NewVerifier(rawURL string)*VerifyData {
	return &VerifyData{
		rawURL: rawURL,
		allowHttpCheckInterval: false,
		httpCheckEnabled: false,
	}
}
