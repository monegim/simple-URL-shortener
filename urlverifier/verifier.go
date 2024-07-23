package urlverifier

import (
	"net/url"

	"github.com/asaskevich/govalidator"
)

type VerifyData struct {
	rawURL                 string
	urlComponents          *url.URL
	Result                 *Result
	httpCheckEnabled       bool
	allowHttpCheckInterval bool
}

type Result struct {
	IsURL        bool
	IsRFC3986URI bool
	IsRFC3986URL bool
	HTTP         *HTTP
}

func NewVerifier(rawURL string) *VerifyData {
	return &VerifyData{
		rawURL:                 rawURL,
		allowHttpCheckInterval: false,
		httpCheckEnabled:       false,
		Result: &Result{},
	}
}
func (v *VerifyData) Verify() error {
	v.Result.IsURL = govalidator.IsURL(v.rawURL)
	if v.Result.IsURL {
		p, err := url.Parse(v.rawURL)
		if err != nil {
			return err
		}
		v.urlComponents = p
	}
	// v.result.IsRFC3986URI = v.is
	return nil
}

func (v *VerifyData) IsRequestURI() bool {
	_, err := url.ParseRequestURI(v.rawURL)
	return err == nil
}
