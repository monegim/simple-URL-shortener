package urlverifier_test

import (
	"log"
	"simple-url-shortener/urlverifier"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testURLs = []struct {
	rawURL string
	isURL  bool
	isURI  bool
}{
	{
		rawURL: "http://example.com",
		isURL:  true,
		isURI:  true,
	},
	{
		rawURL: "https://example.com",
		isURL:  true,
	},
	{
		rawURL: "example.com/mostafa/test",
		isURL:  true,
	},
	{
		rawURL: "_example.com/mostafa/test",
		isURL:  false,
	},
}

var (
	urlToCheck string
	verifier   *urlverifier.VerifyData
	err        error
	isURI      bool
)

func TestCheckVerify(t *testing.T) {

	for _, test := range testURLs {
		urlToCheck = test.rawURL
		verifier = urlverifier.NewVerifier(urlToCheck)
		err = verifier.Verify()
		if err != nil {
			log.Fatal(err)
		}
		assert.Equal(t, verifier.Result.IsURL, test.isURL)

	}
}

func TestIsRequestURI(t *testing.T) {
	for _, test := range testURLs {
		verifier = urlverifier.NewVerifier(test.rawURL)
		isURI = verifier.IsRequestURI()
		assert.Equal(t, isURI, test.isURI)
	}
}
