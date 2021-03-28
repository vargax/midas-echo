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
	CatalogoID    int
	Descripcion   string
	Publicaciones []Publicacion
}

type Publicacion struct {
	gorm.Model
	LoteID   int
	Nombre   string
	Archivos []Archivo
}

type Archivo struct {
	gorm.Model
	PublicacionID int
	Nombre        string
	Ruta          string
}
