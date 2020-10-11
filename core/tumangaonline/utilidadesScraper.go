package tumangaonline

import (
	"strings"
)

func AdapterStringUrl(url string) string {
	if strings.Contains(url, "paginated") {
		return url
	} else if strings.Contains(url, "cascade") {
		newUrl := strings.Replace(url, "/cascade", "/paginated", -1)
		return newUrl
	}
	return url
}