package models

type MangaUnion struct {
	Title       string   `json:"title"`
	Descripcion string   `json:"descripcion"`
	Capitulos   []string `json:"capitulos"`
	Imagen      string   `json:"imagen"`
}
