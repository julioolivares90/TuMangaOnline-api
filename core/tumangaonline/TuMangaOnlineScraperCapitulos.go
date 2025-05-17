package tumangaonline

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gocolly/colly"
)

// GetCookiesFromTMO obtiene las cookies de una sesion en la pagina de tmo
func GetCookiesFromTMO() (map[string]string, error) {

	cookies := make(map[string]string)

	//dataResponse := http.Response{}
	client := http.Client{}

	request, err := http.NewRequest("GET", "https://lectortmo.com/", nil)

	if err != nil {
		return cookies, err
	}

	request.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.80 Safari/537.36 Edg/86.0.622.43")

	request.Header.Set("sec-fetch-dest", "document")
	request.Header.Set("referer", "https://lectortmo.com/")
	request.Header.Set("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	request.Header.Set("method", "GET")
	request.Header.Set("authority", "lectortmo.com")
	request.Header.Set("accept-encoding", "gzip, deflate, br")
	request.Header.Set("accept-language", "es-419,es;q=0.9,es-ES;q=0.8,en;q=0.7,en-GB;q=0.6,en-US;q=0.5")

	response, err := client.Do(request)

	if err != nil {
		return cookies, err
	}

	cookiesResp := response.Cookies()

	for _, cookie := range cookiesResp {
		cookies[cookie.Name] = cookie.Value
	}

	return cookies, nil

}

//getImagesFromHTMLParsed obtiene los enlaces de las imagenes de cada manga

//GetPageFromTMO get the page that contains the list of images of a manga
//parameter url, urlRefer
//
//exple GetPageFromTMO("https://lectortmo.com/view_uploads/490456")

func getPathFromUrl(url string) string {
	return strings.Replace(url, "https://zonatmo.com", "", 1)
}

func GetImageChapter2(urlRefer string, url string) ([]string, error) {
	var imagenes []string

	c := colly.NewCollector()

	c.OnRequest(func(request *colly.Request) {
		path := getPathFromUrl(url)

		request.Headers.Set("path", path)
		request.Headers.Set("referer", urlRefer)
		request.Headers.Set("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
		request.Headers.Set("method", "GET")
		request.Headers.Set("authority", "zonatmo.com")

	})

	c.OnResponse(func(response *colly.Response) {
		urlVisitated := response.Request.URL.String()
		fmt.Print(string(response.Body))
		if strings.Contains(urlVisitated, "/paginated") || strings.Contains(urlVisitated, "/paginated/1") {
			newUrl := strings.Replace(urlVisitated, "/paginated", "/cascade", 1)
			if strings.Contains(newUrl, "/cascade/1") {
				newUrl = strings.Replace(newUrl, "/cascade/1", "/cascade", 1)
				//imagenes , _ = GetImageChapter2(urlRefer,newUrl)
				imagenes = ScraperNewPage(newUrl, urlVisitated)
			}
			imagenes = ScraperNewPage(newUrl, urlVisitated)

		} else if strings.Contains(urlVisitated, "/cascade") {
			response.Request.Visit(urlVisitated)
		}
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Errorf(err.Error())
	})
	c.Visit(url)
	return imagenes, nil
}

func ScraperNewPage(newUrlToScrape string, urlRefered string) []string {
	var imagenes []string
	c := colly.NewCollector()

	c.OnHTML("#app > #main-container", func(element *colly.HTMLElement) {
		fmt.Println(element.DOM.Text())
		imagenes = getImagesFromHTMLParsed2(element)
	})

	c.OnRequest(func(request *colly.Request) {
		path := getPathFromUrl(newUrlToScrape)

		request.Headers.Set("path", path)
		//request.Headers.Set("referer", urlRefered)
		request.Headers.Set("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
		request.Headers.Set("method", "GET")
		request.Headers.Set("authority", "zonatmo.com")
		//request.Headers.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/136.0.0.0 Safari/537.36")

	})

	c.Visit(newUrlToScrape)
	return imagenes
}

func getImagesFromHTMLParsed2(document *colly.HTMLElement) []string {
	var images []string

	fmt.Print(document.DOM.Html())

	document.ForEach("div.img-container", func(i int, element *colly.HTMLElement) {
		image := element.ChildAttr("img", "data-src")
		images = append(images, image)
		fmt.Println(image)

	})

	fmt.Println(images)
	return images
}
