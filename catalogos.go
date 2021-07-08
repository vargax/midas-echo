package midas

type Catalogo struct {
	ID        uint
	EsPublico bool
	Lotes     []Lote `json:",omitempty"`
}

type Lote struct {
	ID            uint
	CatalogoID    uint
	Descripcion   string
	Publicaciones []Publicacion `json:",omitempty"`
}

type Publicacion struct {
	ID     uint
	LoteID uint
	Nombre string
}

type CatalogoSrv interface {
	New(c *Catalogo) error
	AddLote(lt *Lote) error
	AddPublicacion(pub *Publicacion) error
	Catalogo(id uint) (*Catalogo, error)
	All(preload bool) (*[]Catalogo, error)
}
