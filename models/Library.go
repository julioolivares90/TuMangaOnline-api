package models

type Library struct {
	Title       string `json:"title"`
	Score       string `json:"score"`
	Type        string `json:"type"`
	Demography  string `json:"demography"`
	MangaUrl    string `json:"mangaUrl"`
	MangaImagen string `json:"mangaImagen"`
}
