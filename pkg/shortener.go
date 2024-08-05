package pkg

func URLShortener(url string) string {
	id := getIdFromDB(url)
	shortenedString := id2String(id)
	return shortenedString
}

func getIdFromDB(url string) uint64 {
	return 0
}
func id2String(id uint64) string {
	return ""
}
