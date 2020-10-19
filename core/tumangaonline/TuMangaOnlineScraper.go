package tumangaonline

import (
	"fmt"

	"net/http"
	"strings"

	s "strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
	"github.com/julioolivares90/TumangaOnlineApi/models"
	"github.com/julioolivares90/TumangaOnlineApi/utilities"
)

//GetMangasPopulares obtiene los mangas mas populares
func GetMangasPopulares(pageNumber int) []models.MangaTMO {
	var mangasPopulares []models.MangaTMO
	url := utilities.TUMANGAONLINE_BASE_URL
	c := colly.NewCollector()

	c.OnHTML("#app > main > div:nth-child(2) > div.col-12.col-lg-8.col-xl-9 > div:nth-child(1)", func(element *colly.HTMLElement) {
		element.ForEach("div.element", func(i int, element *colly.HTMLElement) {
			dataItentificador := element.Attr("data-identifier")
			//fmt.Println(dataItentificador)
			mangaPopular := models.MangaTMO{
				Title:       element.ChildText("a > div > div > h4"),
				MangaUrl:    element.ChildAttr("a", "href"),
				Type:        element.ChildText("a > div > span.book-type"),
				Demography:  element.ChildText("a > div > span.demography"),
				Score:       element.ChildText("a > div > span.score > span"),
				MangaImagen: getImagenManga(element.ChildText("a > div > style"), dataItentificador),
			}
			mangasPopulares = append(mangasPopulares, mangaPopular)

		})
	})
	c.Visit(fmt.Sprintf("%s/populars?page=%d", url, pageNumber))
	return mangasPopulares
}

//GetMangasPopularesJosei
func GetMangasPopularesJosei() []models.MangaTMO {
	var mangasPopulares []models.MangaTMO
	url := utilities.TUMANGAONLINE_BASE_URL
	c := colly.NewCollector()

	c.OnHTML("#app > main > div:nth-child(2) > div.col-12.col-lg-8.col-xl-9 > div:nth-child(1)", func(element *colly.HTMLElement) {
		element.ForEach("div.element", func(i int, element *colly.HTMLElement) {
			dataItentificador := element.Attr("data-identifier")
			//fmt.Println(dataItentificador)
			mangaPopular := models.MangaTMO{
				Title:       element.ChildText("a > div > div > h4"),
				MangaUrl:    element.ChildAttr("a", "href"),
				Type:        element.ChildText("a > div > span.book-type"),
				Demography:  element.ChildText("a > div > span.demography"),
				Score:       element.ChildText("a > div > span.score > span"),
				MangaImagen: getImagenManga(element.ChildText("a > div > style"), dataItentificador),
			}
			mangasPopulares = append(mangasPopulares, mangaPopular)

		})
	})
	c.Visit(fmt.Sprintf("%s/populars-girls", url))
	return mangasPopulares
}

//GetMangasPopularesSeinen
func GetMangasPopularesSeinen() []models.MangaTMO {
	var mangasPopulares []models.MangaTMO
	url := utilities.TUMANGAONLINE_BASE_URL
	c := colly.NewCollector()

	c.OnHTML("#app > main > div:nth-child(2) > div.col-12.col-lg-8.col-xl-9 > div:nth-child(1)", func(element *colly.HTMLElement) {
		element.ForEach("div.element", func(i int, element *colly.HTMLElement) {
			dataItentificador := element.Attr("data-identifier")
			//fmt.Println(dataItentificador)
			mangaPopular := models.MangaTMO{
				Title:       element.ChildText("a > div > div > h4"),
				MangaUrl:    element.ChildAttr("a", "href"),
				Type:        element.ChildText("a > div > span.book-type"),
				Demography:  element.ChildText("a > div > span.demography"),
				Score:       element.ChildText("a > div > span.score > span"),
				MangaImagen: getImagenManga(element.ChildText("a > div > style"), dataItentificador),
			}
			mangasPopulares = append(mangasPopulares, mangaPopular)

		})
	})
	c.Visit(fmt.Sprintf("%s/populars-boys", url))
	return mangasPopulares
}

//GetInfoManga obtiene la informacion de un manga
func GetInfoManga(url string) models.MangaInfoTMO {
	c := colly.NewCollector()
	mangaInfo := models.MangaInfoTMO{}

	c.OnHTML("#app > section", func(element *colly.HTMLElement) {
		mangaInfo.Title = element.ChildText("header > section.element-header-content > div.container.h-100 > div > div.col-12.col-md-9.element-header-content-text > h1")
		mangaInfo.Image = element.ChildAttr("header > section.element-header-content > div.container.h-100 > div > div.col-12.col-md-3.text-center > div > img", "src")
		mangaInfo.Tipo = element.ChildText("header > section.element-header-content > div.container.h-100 > div > div.col-12.col-md-3.text-center > h1")
		mangaInfo.Score = element.ChildText("header > section.element-header-content > div.container.h-100 > div > div.col-12.col-md-3.text-center > div > div.score > a > span")
		mangaInfo.Demografia = element.ChildText("header > section.element-header-content > div.container > div.row > div.col-12 > div.element-image > div.demography")
		mangaInfo.Descripcion = element.ChildText("header > section.element-header-content > div.container.h-100 > div > div.col-12.col-md-9.element-header-content-text > p.element-description")
		mangaInfo.Estado = element.ChildText("header > section.element-header-content > div.container.h-100 > div > div.col-12.col-md-9.element-header-content-text > span.book-status")

		var generos []string
		element.ForEach("header > section > div.container > div.row > div.col-12 > h6", func(i int, element *colly.HTMLElement) {
			generos = append(generos, element.Text)
		})

		mangaInfo.Generos = generos
		var capitulos []models.Capitulo

		//obtener los li que no estan en colapse
		element.ForEach("#chapters > ul.list-group > li", func(i int, element *colly.HTMLElement) {
			cap := models.Capitulo{
				Title:   element.ChildText("h4 > div.row > div > a.btn-collapse"),
				UrlLeer: element.ChildAttr("div > div > ul > li  > div.row > div.col-2 > a", "href"),
			}
			capitulos = append(capitulos, cap)
		})

		//obtener los li que estan en colapsed
		element.ForEach("#chapters > ul.list-group > div > li", func(i int, element *colly.HTMLElement) {
			cap := models.Capitulo{
				Title:   element.ChildText("h4 > div.row > div > a.btn-collapse"),
				UrlLeer: element.ChildAttr("div > div > ul > li  > div.row > div.col-2 > a", "href"),
			}
			capitulos = append(capitulos, cap)
		})

		mangaInfo.Capitulos = capitulos
	})
	c.Visit(url)
	return mangaInfo
}

//getImagenManga obtiene la imagen de fondo de la lista de mangas de la pagina de inicio
func getImagenManga(imagenUrl string, mangaIdentificador string) string {
	cadenaBorrar := fmt.Sprintf(".book-thumbnail-%s::before{\n                     ", mangaIdentificador)

	cadenaSinEspacios := s.Trim(imagenUrl, "")

	cad1 := s.ReplaceAll(cadenaSinEspacios, "background-image: url('", " ")
	cad2 := s.ReplaceAll(cad1, "');\n                }", "")
	cad3 := s.ReplaceAll(cad2, cadenaBorrar, "")

	return s.TrimLeft(cad3, "")
}

//getImagenListaManga obtiene la imagen de fondo de la lista de mangas de la pagina de inicio
func getImagenListaManga(imagenUrl string, mangaIdentificador string) string {

	cadenaSinBackground := s.Replace(imagenUrl, "background-image: url('", " ", -1)

	cadenaSinClase := s.Replace(cadenaSinBackground, ".list-thumbnail", " ", -1)

	cadenaSinBefore := s.Replace(cadenaSinClase, "::before{\n", " ", -1)

	cadenaSinLLave := s.Replace(cadenaSinBefore, "}", " ", -1)
	cad1 := s.Replace(cadenaSinLLave, "');\n", " ", -1)

	//cad2 := fmt.Sprintf("-%s", mangaIdentificador)
	//cadSinDataSrc := s.Replace(cad1, cad2, " ", -1)

	return cad1

}

//https://lectormanga.com/library?title=&order_field=title&order_item=likes_count&order_dir=desc&type=&demography=seinen&webcomic=&yonkoma=&amateur=&erotic=true
//https://lectortmo.com/library?title=&order_field=title&order_item=likes_count&order_dir=desc&type=&demography=seinen&webcomic=&yonkoma=&amateur=&erotic=true
//https://lectortmo.com/library?order_item=likes_count&order_dir=desc&title=&_page=1&filter_by=title&type=&demography=&status=&translation_status=&webcomic=&yonkoma=&amateur=&erotic=
//title string, orderField string, orderItem string, orderDir string, Type string, demography string, webcomic string, yonkoma string, amateur string, erotic string
func BuscarMangas(order_item string, order_dir string, title string, _page string, filter_by string, Type string, demography string, status string, translation_status string, webcomic string, yonkoma string, amateur string, erotic string) []models.Library {
	var mangas []models.Library
	c := colly.NewCollector()
	url := fmt.Sprintf("https://lectortmo.com/library?order_item=%s&order_dir=%s&title=%s&_page=%s&filter_by=%s&type=%s&demography=%s&status=%s&translation_status=%s&webcomic=%s&yonkoma=%s&amateur=%s&erotic=%s",
		order_item, order_dir, title, _page, filter_by, Type, demography, status, translation_status, webcomic, yonkoma, amateur, erotic)

	c.OnHTML("#app > main > div:nth-child(2) > div.col-12.col-lg-8.col-xl-9 > div:nth-child(3)", func(element *colly.HTMLElement) {

		element.ForEach("div.element", func(i int, element *colly.HTMLElement) {

			dataItentificador := element.Attr("data-identifier")

			library := models.Library{
				Title:       element.ChildText("a > div > div > h4"),
				MangaUrl:    element.ChildAttr("a", "href"),
				Type:        element.ChildText("a > div > span.book-type"),
				Demography:  element.ChildText("a > div > span.demography"),
				Score:       element.ChildText("a > div > span.score > span"),
				MangaImagen: getImagenManga(element.ChildText("a > div > style"), dataItentificador),
			}

			mangas = append(mangas, library)

		})
	})

	c.Visit(url)
	return mangas

}

func GetLibraryMangas() []models.ListaManga {
	c := colly.NewCollector()
	var lista []models.ListaManga
	c.OnHTML("#app > main > div:nth-child(2) > div.col-12.col-lg-8.col-xl-9 > div:nth-child(3)", func(element *colly.HTMLElement) {
		element.ForEach("div.col-12.col-sm-12", func(i int, element *colly.HTMLElement) {

			urlLista := element.ChildAttr("a", "href")
			titleLista := element.ChildText("div.thumbnail > div.thumbnail-title > h4.text-truncate")
			descripcionLista := element.ChildText("div.thumbnail > div.thumbnail-description > p")
			cantidadDeSeguidoresLista := element.ChildText("div.thumbnail > div.thumbnail-container > span.followers_count")

			imagenLista := element.ChildText("div.thumbnail > style")
			dataSRC := utilities.GetDataSRCFromURL2(urlLista, titleLista)

			//fmt.Println(imagenLista)
			imagenFinal := getImagenListaManga(imagenLista, dataSRC)
			datos := models.ListaManga{
				Url:                       urlLista,
				Title:                     titleLista,
				Descripcion:               descripcionLista,
				CantidadDeSeguidoresLista: cantidadDeSeguidoresLista,
				ImagenLista:               imagenFinal,
				DataSRC:                   dataSRC,
			}
			lista = append(lista, datos)
		})
	})
	c.Visit("https://lectortmo.com/lists")
	return lista
}

func GetCookiesFromTMO() (map[string]string, error) {
	cookies := make(map[string]string)
	response, err := http.Get("https://lectortmo.com/")

	if err != nil {
		return cookies, err
	}

	cookiesResp := response.Cookies()

	for _, cookie := range cookiesResp {
		cookies[cookie.Name] = cookie.Value
	}

	return cookies, nil

}

func GetImageChapter(url string) ([]string, error) {
	var images []string
	page, err := GetPageFromTMO(url)
	if err != nil {
		return images, err
	}

	//document, err := goquery.NewDocumentFromReader(page.Body)

	defer page.Body.Close()

	if err != nil {
		return images, err
	}

	if strings.Contains(page.Request.URL.String(), "/paginated") {

		fmt.Println(page.Request.URL.String())

		fmt.Println("Obteniendo nueva pagina con vista de cascade")

		oldUrl := page.Request.URL.String()

		newUrl := strings.Replace(oldUrl, "/paginated", "/cascade", 1)

		newPage, err := GetPageFromTMO(newUrl)

		if err != nil {
			return images, err
		}

		document, err := goquery.NewDocumentFromReader(newPage.Body)

		if err != nil {
			return images, err
		}

		defer newPage.Body.Close()

		doc := document.Find("#app").First()
		images = getImagesFromHTMLParsed(doc)

	} else if strings.Contains(page.Request.URL.String(), "/cascade") {
		fmt.Println("Obteniendo imagenes")
		url := page.Request.URL.String()
		newPage, err := GetPageFromTMO(url)
		if err != nil {
			return images, err
		}

		document, err := goquery.NewDocumentFromReader(newPage.Body)

		if err != nil {
			return images, err
		}

		defer newPage.Body.Close()

		doc := document.Find("#app").First()
		images = getImagesFromHTMLParsed(doc)
	}

	return images, nil
}

func getImagesFromHTMLParsed(document *goquery.Selection) []string {
	var images []string

	viewContainer := document.Find("#app > viewer-container").First()
	if viewContainer != nil {
		viewContainer.Find("viewer-image-container").Each(func(i int, imagen *goquery.Selection) {
			img := imagen.Find("img").First().AttrOr("data-src", "")
			if img != "" {
				images = append(images, img)
			}
		})
	}
	mainContainer := document.Find("#app > #main-container").First()
	if mainContainer != nil {
		mainContainer.Find("div.img-container").Each(func(i int, imagen *goquery.Selection) {
			img := imagen.Find("img").First().AttrOr("data-src", "")
			if img != "" {
				images = append(images, img)
			}
		})
	}
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

	//defer response.Body.Close()

	/*
		bodyResponse, err := ioutil.ReadAll(response.Body)

		if err != nil {
			return data, err
		}
	*/
	data = *response
	//fmt.Println("DATA => GetPageFromTMO", data)
	return data, nil
}
