package gorm

import (
	"fmt"
	"github.com/jackc/pgconn"
	"github.com/vargax/midas-echo"
	"github.com/vargax/midas-echo/env"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
	"os"
	"strconv"
)

const (
	duplicateKey = "23505"
)

type StorageSrv struct {
	db *gorm.DB
}

func New() *StorageSrv {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		os.Getenv(env.DbHost),
		os.Getenv(env.DbPort),
		os.Getenv(env.DbUser),
		os.Getenv(env.DbName),
		os.Getenv(env.DbPass),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	debug, _ := strconv.ParseBool(os.Getenv(env.DebugRepo))
	if debug {
		db.Logger = logger.Default.LogMode(logger.Info)
	}

	err = db.AutoMigrate(
		&Catalogo{}, &Lote{}, &Publicacion{},
		&User{})
	if err != nil {
		panic(err)
	}

	return &StorageSrv{db: db}
}

func (s *StorageSrv) DbInitRequired() bool {
	count := int64(-1)

	result := s.db.Model(&User{}).Count(&count)
	if result.Error != nil {
		panic(result.Error)
	}

	return count == 0
}

type Catalogo struct {
	*midas.Catalogo
	gorm.Model
}

func (s *StorageSrv) CreateCatalogo(c *midas.Catalogo) error {
	cat := Catalogo{Catalogo: c}

	result := s.db.Create(&cat)
	if e, ok := result.Error.(*pgconn.PgError); ok && e.Code == duplicateKey {
		return midas.DuplicateKey
	}
	return result.Error
}
func (s *StorageSrv) SelectCatalogo(c *midas.Catalogo) error {
	result := s.db.Preload(clause.Associations).Where(&c).First(&c)
	return result.Error
}
func (s *StorageSrv) SelectAllCatalogos(cc *[]midas.Catalogo, preload bool) error {
	if preload {
		result := s.db.Preload(clause.Associations).Find(cc)
		return result.Error
	}

	result := s.db.Find(cc)
	return result.Error
}

type Lote struct {
	*midas.Lote
	gorm.Model
}

func (s *StorageSrv) CreateLote(lt *midas.Lote) error {
	l := Lote{Lote: lt}
	result := s.db.Create(&l)
	return result.Error
}

type Publicacion struct {
	*midas.Publicacion
	gorm.Model
}

func (s *StorageSrv) CreatePublicacion(pub *midas.Publicacion) error {
	p := Publicacion{Publicacion: pub}
	result := s.db.Create(&p)
	return result.Error
}

type User struct {
	*midas.User
	gorm.Model
}

func (s *StorageSrv) CreateUser(u *midas.User) error {
	user := User{User: u}

	result := s.db.Create(&user)
	if e, ok := result.Error.(*pgconn.PgError); ok && e.Code == duplicateKey {
		return midas.DuplicateKey
	}
	return result.Error
}
func (s *StorageSrv) SelectUser(u *midas.User) error {
	result := s.db.Where(&u).First(&u)
	return result.Error
}
