package models

type PaginasManga struct {
	CantidadPaginas int      `json:"cantidadPaginas"`
	Paginas         []string `json:"paginas"`
}
