package models

type MangaInfoTMO struct {
	Title string `json:"title"`

	Image string `json:"image"`

	Tipo string `json:"tipo"`

	Score string `json:"score"`

	Demografia string `json:"demografia"`

	Descripcion string `json:"descripcion"`

	Estado string `json:"estado"`

	Generos []string `json:"generos"`

	Capitulos []Capitulo `json:"capitulo"`
}

type Capitulo struct {
	Title   string
	UrlLeer string
}
