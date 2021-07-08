package midas

import (
	"errors"
	"gorm.io/gorm"
)

type StorageSrv interface {
	DbInitRequired() bool

	CreateUser(u *User) error
	SelectUser(u *User) error

	CreateCatalogo(c *Catalogo) error
	SelectCatalogo(c *Catalogo) error
	SelectAllCatalogos(cc *[]Catalogo, preload bool) error
	CreateLote(lt *Lote) error
	CreatePublicacion(pub *Publicacion) error
}

var (
	DuplicateKey   = errors.New("duplicate key")
	RecordNotFound = gorm.ErrRecordNotFound
)
