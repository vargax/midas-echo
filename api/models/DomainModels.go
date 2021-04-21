package models

import (
	"gorm.io/gorm"
)

type Catalogo struct {
	ID         int
	EsPublico  bool
	Lotes      []Lote `json:",omitempty"`
	gorm.Model `json:"-"`
}

type Lote struct {
	ID            int
	CatalogoID    int
	Descripcion   string
	Publicaciones []Publicacion `json:",omitempty"`
	gorm.Model    `json:"-"`
}

type Publicacion struct {
	ID         int
	LoteID     int
	Nombre     string
	Archivos   []File `json:",omitempty"`
	gorm.Model `json:"-"`
}
