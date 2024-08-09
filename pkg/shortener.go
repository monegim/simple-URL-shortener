package pkg

import (
	"strings"
)

const (
	HASH_LENGTH = 6
)

var (
	charSet     = "abcdefghijklmnopqrstuvwxyz"
	BASE_DOMAIN = "mostafa.com"
)

func URLShortener(id uint64, url string) string {
	// id := getIdFromDB(url)
	shortenedString := id2String(id)
	return shortenedString
}

// func getIdFromDB(url string) uint64 {
// 	return 0
// }
func id2String(id uint64) string {
	base := uint64(26*2 + 10)
	ch := buildCharset()
	var rem int
	shortURL := ""
	number := id
	for {
		if number/base == 0 {
			shortURL = string(ch[number]) + shortURL
			break
		}
		rem = int(number % base)
		shortURL = string(ch[rem]) + shortURL
		number = number / base
	}

	return shortURL
}

func string2Id(url string) uint64 {
	n := len(url)
	ch := buildCharset()
	result := uint64(0)
	for i := n - 1; i >= 0; i-- {
		m := IntPow(len(ch), n-i-1)
		index := uint64(getIndex(string(url[i])))
		result += uint64(index * m)
	}
	return result
}

func getIndex(c string) int {
	baseString := buildCharset()
	return strings.Index(baseString, c)
}

func buildCharset() string {
	return charSet + strings.ToUpper(charSet) + "0123456789"
}

func IntPow(n, m int) uint64 {
	if m == 0 {
		return 1
	}

	if m == 1 {
		return uint64(n)
	}

	result := uint64(n)
	for i := 2; i <= m; i++ {
		result *= uint64(n)
	}
	return result
}
