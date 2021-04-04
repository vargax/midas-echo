package app

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
	Archivos   []Archivo `json:",omitempty"`
	gorm.Model `json:"-"`
}

type Archivo struct {
	ID            int
	PublicacionID int
	Nombre        string
	Ruta          string
	gorm.Model    `json:"-"`
}

type User struct {
	ID    int
	email string
	gorm.Model
}
