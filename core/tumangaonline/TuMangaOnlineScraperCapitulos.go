package tumangaonline

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strings"
)

//GetCookiesFromTMO obtiene las cookies de una sesion en la pagina de tmo
func GetCookiesFromTMO() (map[string]string, error) {
	cookies := make(map[string]string)
	//dataResponse := http.Response{}
	client := http.Client{}

	request, err := http.NewRequest("GET", "https://lectortmo.com/", nil)

	if err != nil {
		return cookies, err
	}

	request.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.80 Safari/537.36 Edg/86.0.622.43")
	request.Header.Set("sec-fetch-site", "none")
	request.Header.Set("sec-fetch-mode", "navigate")
	request.Header.Set("sec-fetch-user", "?1")
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

//GetImageChapter retorna la lista de las imagenes o un error
func GetImageChapter(url string) ([]string, error) {

	var images []string

	page, err := GetPageFromTMO(url)

	if err != nil {
		return images, err
	}

	defer page.Body.Close()

	if strings.Contains(page.Request.URL.String(), "/paginated/1") || strings.Contains(page.Request.URL.String(), "/paginated") {

		fmt.Println("URL A VISITAR =>", page.Request.URL.String())

		fmt.Println("Obteniendo nueva pagina con vista de cascade")

		oldUrl := page.Request.URL.String()

		newUrl := strings.Replace(oldUrl, "/paginated", "/cascade", 1)

		if strings.Contains(newUrl, "cascade/1") {
			newUrl = strings.Replace(newUrl, "/cascade/1", "/cascade", 1)
		}

		fmt.Println("NUEVA URL =>", newUrl)

		newPage, err := GetPageFromTMO(newUrl)

		if err != nil {
			fmt.Println(err.Error())
			return images, err
		}

		document, err := goquery.NewDocumentFromReader(newPage.Body)

		if err != nil {
			fmt.Println(err.Error())
			return images, err
		}

		defer newPage.Body.Close()

		doc := document.Find("#app").First()
		images = getImagesFromHTMLParsed(doc)

	} else if strings.Contains(page.Request.URL.String(), "/cascade") {
		fmt.Println("Obteniendo imagenes")

		url := page.Request.URL.String()

		fmt.Print("URL => ", url)

		newPage, err := GetPageFromTMO(url)

		if err != nil {
			fmt.Println(err.Error())
			return images, err
		}

		document, err := goquery.NewDocumentFromReader(newPage.Body)

		if err != nil {
			fmt.Println(err.Error())
			return images, err
		}

		defer page.Body.Close()

		doc := document.Find("#app").First()

		fmt.Println(doc.Html())
		images = getImagesFromHTMLParsed(doc)
	}

	return images, nil
}

//getImagesFromHTMLParsed obtiene los enlaces de las imagenes de cada manga
func getImagesFromHTMLParsed(document *goquery.Selection) []string {
	var images []string

	viewContainer := document.Find("#app > #viewer-container").First()
	if viewContainer != nil {
		viewContainer.Find("div.viewer-image-container").Each(func(i int, imagen *goquery.Selection) {
			img := imagen.Find("img").First().AttrOr("data-src", "")
			if img != "" {
				fmt.Println(img)
				images = append(images, img)
			}
		})
	}
	mainContainer := document.Find("#app > #main-container").First()
	if mainContainer != nil {
		mainContainer.Find("div.img-container").Each(func(i int, imagen *goquery.Selection) {
			img := imagen.Find("img").First().AttrOr("data-src", "")
			if img != "" {
				fmt.Println(img)
				images = append(images, img)
			}
		})
	}
	fmt.Println(images)
	return images
}

//GetPageFromTMO get the page that contains the list of images of a manga
//parameter url
//exple GetPageFromTMO("https://lectortmo.com/view_uploads/490456")
func GetPageFromTMO(url string) (http.Response, error) {
	//var data []byte
	data := http.Response{}

	client := &http.Client{}

	request, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return data, err
	}
	cookies, errCookie := GetCookiesFromTMO()

	if errCookie != nil {
		return data, errCookie
	}

	request.Header.Set("referer", "https://lectortmo.com/")
	request.Header.Set("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	request.Header.Set("method", "GET")
	request.Header.Set("authority", "lectortmo.com")
	request.Header.Set("accept-encoding", "gzip, deflate, br")
	request.Header.Set("accept-language", "es-419,es;q=0.9,es-ES;q=0.8,en;q=0.7,en-GB;q=0.6,en-US;q=0.5")

	fmt.Println("COOKIES => ", cookies)
	cookieSession := cookies["tumangaonline_session"]
	cookieXSRF_TOKEN := cookies["XSRF-TOKEN"]
	cookie__cfduid := cookies["__cfduid"]

	cookie1 := http.Cookie{Name: "tumangaonline_session", Value: cookieSession}
	cookie2 := http.Cookie{Name: "XSRF-TOKEN", Value: cookieXSRF_TOKEN}
	cookie3 := http.Cookie{Name: "__cfduid", Value: cookie__cfduid}
	request.AddCookie(&cookie1)
	request.AddCookie(&cookie2)
	request.AddCookie(&cookie3)

	response, err := client.Do(request)

	if err != nil {
		return data, err
	}

	data = *response
	//fmt.Println("DATA => GetPageFromTMO", data)
	return data, nil
}
