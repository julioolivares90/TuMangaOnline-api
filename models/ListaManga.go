package models

type ListaManga struct {
	Url                       string `json:"url"`
	Title                     string `json:"title"`
	Descripcion               string `json:"descripcion"`
	CantidadDeSeguidoresLista string `json:"cantidad_seguidores_lista"`
	ImagenLista               string `json:"imagen_lista"`
	DataSRC                   string `json:"dataSRC"`
}
