package models

type PostInfo struct {
	Title          string   `json:"title"`
	ImagenFondo    string   `json:"imagenFondo"`
	Paginas        []string `json:"paginas"`
	UrlLectorManga string   `json:"urlLectorManga"`
}
