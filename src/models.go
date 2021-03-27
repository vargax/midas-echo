package main

import (
	"gorm.io/gorm"
)

type Catalogo struct {
	gorm.Model
	EsPublico bool
	Lotes     []Lote
}

type Lote struct {
	gorm.Model
	CatalogoID    uint
	Descripcion   string
	Publicaciones []Publicacion
}

type Publicacion struct {
	gorm.Model
	LoteID   uint
	Nombre   string
	Archivos []Archivo
}

type Archivo struct {
	gorm.Model
	PublicacionID uint
	Nombre        string
	Ruta          string
}
