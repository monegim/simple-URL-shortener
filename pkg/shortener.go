package pkg

import "strings"

const (
	HASH_LENGTH = 6
)

func URLShortener(url string) string {
	id := getIdFromDB(url)
	shortenedString := id2String(id)
	return shortenedString
}

func getIdFromDB(url string) uint64 {
	return 0
}
func id2String(id uint64) string {
	base := uint64(26*2 + 10)
	charSet := "abcdefghijklmnopqrstuvwxyz"
	charSet = charSet + strings.ToUpper(charSet) + "01234566789"
	var rem int
	shortURL := ""
	number := id
	for {
		rem = int(number % base)
		shortURL =  string(charSet[rem]) + shortURL
		number = number / base
		if number/base == 0 {
			shortURL =  string(charSet[number]) + shortURL
			break
		} 
	}

	return shortURL
}
