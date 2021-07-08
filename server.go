package midas

type Services struct {
	UserSrv UserSrv
	CatSrv  CatalogoSrv
}

type Server interface {
	Start(s *Services) error
}
