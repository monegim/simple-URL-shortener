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
}{
	{
		rawURL: "http://example.com",
		isURL:  true,
	},
	{
		rawURL: "example.com",
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

func TestCheckVerify(t *testing.T) {
	var (
		urlToCheck string
		verifier   *urlverifier.VerifyData
		err        error
	)

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
