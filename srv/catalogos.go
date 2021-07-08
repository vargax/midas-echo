package srv

import (
	"github.com/vargax/midas-echo"
)

type CatalogoSrv struct {
	s midas.StorageSrv
}

func NewCatalogoSrv(s midas.StorageSrv) *CatalogoSrv {
	return &CatalogoSrv{s: s}
}

func (cs CatalogoSrv) New(c *midas.Catalogo) error {
	return cs.s.CreateCatalogo(c)
}

func (cs CatalogoSrv) AddLote(lt *midas.Lote) error {
	return cs.s.CreateLote(lt)
}

func (cs CatalogoSrv) AddPublicacion(pub *midas.Publicacion) error {
	return cs.s.CreatePublicacion(pub)
}

func (cs CatalogoSrv) Catalogo(id uint) (*midas.Catalogo, error) {
	c := midas.Catalogo{ID: id}
	err := cs.s.SelectCatalogo(&c)
	return &c, err
}

func (cs CatalogoSrv) All(preload bool) (*[]midas.Catalogo, error) {
	var cc []midas.Catalogo
	err := cs.s.SelectAllCatalogos(&cc, preload)
	return &cc, err
}
