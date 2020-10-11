package utilities

import (
	"fmt"
	s "strings"
)

func CleanString(cad string) string {

	return s.Replace(cad, "https://ouo.io/s/lbH3iXRW?s=", "", 1)
}

func GetDataSRCFromURL(url string, nombreLista string) string {
	var nombreList string
	var cadenaFinal string
	//cad1 := s.Replace("https://lectortmo.com/lists/", url, "", 0)
	cad1 := s.Replace(url, "https://lectortmo.com/lists/", "", -1)
	if s.Contains(nombreLista, " ") {

		nombreList = s.Replace(nombreLista, " ", "-", -1)

		cadenaFinal = s.Replace(cad1, nombreList, "", -1)
		fmt.Println("cadena final si contiene espacios ", cadenaFinal)
	} else {

		//nombreList = fmt.Sprintf("%s", nombreLista)
		println("cadena 1 ", cad1)
		cad2 := fmt.Sprintf("/%s", nombreLista)
		fmt.Println("cad2 ", cad2)
		cadenaFinal = s.Replace(cad1, cad2, "", -1)

	}
	return cadenaFinal
}

func GetDataSRCFromURL2(url string, nombreLista string) string {
	var nombreList string
	var cadenaFinal string
	var cadenaSinLectorTMO string
	if s.Contains(nombreLista, " ") {
		nombreList = s.Replace(nombreLista, " ", "-", 0)
		cadenaSinLectorTMO = s.Replace(url, nombreList, " ", -0)
	} else {
		nombreList = nombreLista
		cadenaSinLectorTMO = s.Replace(url, nombreList, "", -0)
	}
	cadenaFinal = cadenaSinLectorTMO
	fmt.Println(cadenaFinal)
	return cadenaFinal
}
