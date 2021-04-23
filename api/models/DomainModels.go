package models

import (
	"gorm.io/gorm"
)

type Catalogo struct {
	ID         uint
	EsPublico  bool
	Lotes      []Lote `json:",omitempty"`
	gorm.Model `json:"-"`
}

type Lote struct {
	ID            uint
	CatalogoID    uint
	Descripcion   string
	Publicaciones []Publicacion `json:",omitempty"`
	gorm.Model    `json:"-"`
}

type Publicacion struct {
	ID         uint
	LoteID     uint
	Nombre     string
	Archivos   []File `json:",omitempty"`
	gorm.Model `json:"-"`
}
